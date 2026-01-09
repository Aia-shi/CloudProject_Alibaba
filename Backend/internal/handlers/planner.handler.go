package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"smart-expense-planner-backend/internal/models"
	"smart-expense-planner-backend/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

type PlannerHandler struct {
	PlannerService services.PlanerService
}

func (ph *PlannerHandler) GetPlannerDataById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		http.Error(w, "token was empty", http.StatusUnauthorized)
	}

	budgetData, err := ph.PlannerService.GetUserPlannerData(ctx, int32(claims["sub"].(float64)))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bugdet was not found", http.StatusInternalServerError)
	}
	jsonBudgetData, _ := json.Marshal(budgetData)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBudgetData)
}

func (ph *PlannerHandler) SendPlannerData(w http.ResponseWriter, r *http.Request) {
	var req models.BudgetData

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(req)

	ctx := r.Context()
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		http.Error(w, "token was empty", http.StatusUnauthorized)
	}

	if err := ph.PlannerService.AddUserPlannerData(r.Context(), req, int32(claims["sub"].(float64))); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"data created"}`))
}
