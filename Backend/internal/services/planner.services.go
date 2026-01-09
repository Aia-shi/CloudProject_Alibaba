package services

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"smart-expense-planner-backend/internal/models"
	repository "smart-expense-planner-backend/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type PlanerService interface {
	GetUserPlannerData(ctx context.Context, user_id int32) (models.BudgetData, error)
	AddUserPlannerData(ctx context.Context, budgetData models.BudgetData, user_id int32) error
}

type BasePlannerService struct {
	Repo *repository.Queries
}

func NewBasePlannerService(conn *pgxpool.Pool) BasePlannerService {
	return BasePlannerService{
		Repo: repository.New(conn),
	}
}

func (s *BasePlannerService) GetUserPlannerData(ctx context.Context, user_id int32) (models.BudgetData, error) {
	budgetData := models.BudgetData{}
	var err error
	budgetData.Income, err = s.Repo.GetUserIncome(ctx, user_id)
	if err != nil {
		log.Println(err.Error())
		return budgetData, errors.New("internal error")
	}

	budgetData.Expense, err = s.Repo.GetUserExpenses(ctx, user_id)
	if err != nil {
		log.Println(err.Error())
		return budgetData, errors.New("interal error")
	}

	budgetData.Period, err = s.Repo.GetUserPeriods(ctx, user_id)
	if err != nil {
		log.Println(err.Error())
		return budgetData, errors.New("internal error")
	}

	return budgetData, nil
}

func (s *BasePlannerService) AddUserPlannerData(ctx context.Context, budgetData models.BudgetData, user_id int32) error {
	for _, per := range budgetData.Period {
		_, err := s.Repo.GetUserPeriodsById(ctx, per.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				s.Repo.InsertUserPeriods(ctx, repository.InsertUserPeriodsParams{
					Name:   per.Name,
					UserID: user_id,
				})
			}
		} else {
			s.Repo.UpdateUserPeriods(ctx, repository.UpdateUserPeriodsParams{
				NewName: per.Name,
			})
		}
	}

	periods, _ := s.Repo.GetUserPeriods(ctx, user_id)
	var databasePeriodsIds []int32
	for _, p := range periods {
		databasePeriodsIds = append(databasePeriodsIds, p.ID)
	}

	var recentPeriodsIds []int32
	for _, p := range budgetData.Period {
		recentPeriodsIds = append(recentPeriodsIds, p.ID)
	}

	if len(databasePeriodsIds) != len(recentPeriodsIds) {
		log.Println("Coś usunięto")
		result := s.findExtraXOR(databasePeriodsIds, recentPeriodsIds)
		s.Repo.DeleteUserPeriods(ctx, result)
	}

	for _, per := range budgetData.Expense {
		_, err := s.Repo.GetUserExpensesById(ctx, per.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				s.Repo.InsertUserExpenses(ctx, repository.InsertUserExpensesParams{
					PeriodID:    per.PeriodID,
					UserID:      user_id,
					Title:       per.Title,
					Amount:      per.Amount,
					Description: per.Description,
					Date:        per.Date,
					Status:      per.Status,
					Category:    per.Category,
				})
			}
		} else {
			s.Repo.UpdateUserExpenses(ctx, repository.UpdateUserExpensesParams{
				NewTitle:       per.Title,
				NewAmount:      per.Amount,
				NewDescription: per.Description,
				NewStatus:      per.Status,
				NewDate:        per.Date,
				NewCategory:    per.Category,
			})
		}
	}

	expenses, _ := s.Repo.GetUserExpenses(ctx, user_id)
	var databaseExpensesIds []int32
	for _, p := range expenses {
		databaseExpensesIds = append(databaseExpensesIds, p.ID)
	}

	var recentExpensesIds []int32
	for _, p := range budgetData.Expense {
		recentExpensesIds = append(recentExpensesIds, p.ID)
	}

	if len(databaseExpensesIds) != len(recentExpensesIds) {
		result := s.findExtraXOR(databaseExpensesIds, recentExpensesIds)
		s.Repo.DeleteUserExpenses(ctx, result)
	}

	for _, per := range budgetData.Income {
		_, err := s.Repo.GetUserIncomeById(ctx, per.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				s.Repo.InsertUserIncomes(ctx, repository.InsertUserIncomesParams{
					PeriodID:    per.PeriodID,
					UserID:      user_id,
					Title:       per.Title,
					Amount:      per.Amount,
					Description: per.Description,
					Date:        per.Date,
					Category:    per.Category,
				})
			}
		} else {
			s.Repo.UpdateUserIncomes(ctx, repository.UpdateUserIncomesParams{
				NewTitle:       per.Title,
				NewAmount:      per.Amount,
				NewDescription: per.Description,
				NewDate:        per.Date,
				NewCategory:    per.Category,
			})
		}
	}

	incomes, _ := s.Repo.GetUserIncome(ctx, user_id)
	var databaseIncomesIds []int32
	for _, p := range incomes {
		databaseIncomesIds = append(databaseIncomesIds, p.ID)
	}

	var recentIncomesIds []int32
	for _, p := range budgetData.Income {
		recentIncomesIds = append(recentIncomesIds, p.ID)
	}

	if len(databaseIncomesIds) != len(recentIncomesIds) {
		result := s.findExtraXOR(databaseIncomesIds, recentIncomesIds)
		s.Repo.DeleteUserIncomes(ctx, result)
	}

	return nil
}

func (s *BasePlannerService) findExtraXOR(a, b []int32) int32 {
	var result int32 = 0

	for _, v := range a {
		result ^= v
	}
	for _, v := range b {
		result ^= v
	}

	return result
}
