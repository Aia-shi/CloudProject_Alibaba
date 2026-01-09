package models

import (
	repository "smart-expense-planner-backend/internal/repositories"
)

type BudgetData struct {
	Income  []repository.Income  `json:"incomes"`
	Expense []repository.Expense `json:"expenses"`
	Period  []repository.Period  `json:"periods"`
}
