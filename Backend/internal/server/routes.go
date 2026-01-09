package server

import (
	"net/http"

	"smart-expense-planner-backend/internal/handlers"
	"smart-expense-planner-backend/internal/middlewares"
	"smart-expense-planner-backend/internal/services"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	conn := NewConnection()

	userService := services.NewBaseUserService(conn)
	userHandler := handlers.UserHandler{
		UserService: &userService,
	}

	plannerService := services.NewBasePlannerService(conn)
	plannerHandler := handlers.PlannerHandler{
		PlannerService: &plannerService,
	}

	mux.HandleFunc("POST /user/register", userHandler.CreateUser)
	mux.HandleFunc("POST /user/login", userHandler.LoginUser)
	mux.HandleFunc("GET /user/logout", userHandler.Logout)

	stack := middlewares.CreateStack(
		middlewares.Logging,
		middlewares.CorsMiddleware,
	)

	authMux := http.NewServeMux()

	authMux.HandleFunc("GET /planner", plannerHandler.GetPlannerDataById)
	authMux.HandleFunc("POST /planner", plannerHandler.SendPlannerData)

	mux.Handle("/", middlewares.Authentication(authMux))

	return stack(mux)
}
