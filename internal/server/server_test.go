package server

import (
	"context"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/shopspring/decimal"
)

//ToDo build out testing

type StubDatabase struct {
	income  income.Income
	incomes []income.Income
}

func (s *StubDatabase) getMonthIncome(_ context.Context, date time.Time) (*income.Income, error) {
	amount, _ := decimal.NewFromString("1550.55")

	income := income.Income{
		Date:   date,
		Source: "Salary",
		Amount: amount,
	}

	return &income, nil
}

func (s *StubDatabase) RecordIncome(_ context.Context, i income.Income) (*income.Income, error) {
	returnedIncome := income.Income{
		ID:     1,
		Date:   i.Date,
		Source: i.Source,
		Amount: i.Amount,
	}

	return &returnedIncome, nil
}

func (s *StubDatabase) ListIncomes(_ context.Context) ([]income.Income, error) {
	amount, _ := decimal.NewFromString("3500.00")

	incomes := []income.Income{{
		Date:   time.Date(2020, time.Now().Month()+1, 12, 0, 0, 0, 0, time.UTC),
		Source: "Salary",
		Amount: amount,
	}}

	return incomes, nil
}
