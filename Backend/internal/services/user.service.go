package services

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"smart-expense-planner-backend/internal/models"
	repository "smart-expense-planner-backend/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

var key []byte = []byte(os.Getenv("APP_JWT_KEY"))

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) error
	LoginUser(ctx context.Context, loginData *models.LoginUserRequest) (string, error)
}

type BaseUserService struct {
	Repo *repository.Queries
}

func NewBaseUserService(conn *pgxpool.Pool) BaseUserService {
	return BaseUserService{
		Repo: repository.New(conn),
	}
}

func (s *BaseUserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) error {
	if err := req.Validate(); err != nil {
		return errors.New("wrong properties")
	}

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(req.Passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println("password hashing failed:", err)
		return errors.New("internal failure")
	}

	err = s.Repo.CreateUser(ctx, repository.CreateUserParams{
		Username:   req.Username,
		Passwdhash: string(hashedPasswd),
		Email:      req.Email,
	})
	if err != nil {
		log.Println("create user failed:", err)
		return errors.New("internal failure")
	}

	return nil
}

func (s *BaseUserService) LoginUser(ctx context.Context, loginData *models.LoginUserRequest) (string, error) {
	user, err := s.Repo.LoginUserWithUsername(ctx, loginData.Login)
	if err != nil {
		return "", errors.New("unauthorized")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Passwdhash), []byte(loginData.Password)); err != nil {
		return "", errors.New("unauthorized")
	}

	token, err := s.generateJWT(user.ID)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("internal failure")
	}

	return token, nil
}

func (s *BaseUserService) generateJWT(userId int32) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
	return t.SignedString(key)
}
