package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/expenses"
	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/catzkorn/pay-yourself-first/internal/saving"
	"github.com/shopspring/decimal"
)

//ToDo build out testing

type StubDataStore struct {
	income      *income.Income
	incomes     []income.Income
	saving      *saving.Saving
	savings     []saving.Saving
	deleteCount []uint32
	expense     *expenses.Expense
	expenses    []expenses.Expense
}

func (s *StubDataStore) RetrieveIncome(_ context.Context, id uint32) (*income.Income, error) {
	income := income.Income{
		ID: id,
	}

	return &income, nil
}

func (s *StubDataStore) RetrieveMonthIncome(_ context.Context, date time.Time) (*income.Income, error) {

	if s.income == nil || !date.Equal(s.income.Date) {
		return nil, income.ErrNoIncomeForMonth
	}

	return s.income, nil
}

func (s *StubDataStore) RecordIncome(_ context.Context, i income.Income) (*income.Income, error) {

	if s.income == nil {
		s.income = &income.Income{
			ID:     1,
			Date:   i.Date,
			Source: i.Source,
			Amount: i.Amount,
		}
		return s.income, nil
	}

	if !s.income.Date.Equal(i.Date) {
		return nil, fmt.Errorf("specified month can not be updated")
	}
	s.income.Source = i.Source
	s.income.Amount = i.Amount

	return s.income, nil
}

func (s *StubDataStore) ListIncomes(_ context.Context) ([]income.Income, error) {
	amount, _ := decimal.NewFromString("3500.00")

	incomes := []income.Income{{
		Date:   time.Date(2020, time.Now().Month()+1, 12, 0, 0, 0, 0, time.UTC),
		Source: "Salary",
		Amount: amount,
	}}

	return incomes, nil
}

func (s *StubDataStore) DeleteIncome(_ context.Context, id uint32) error {
	s.deleteCount = append(s.deleteCount, id)
	return nil
}

func (s *StubDataStore) RecordMonthSavingPercent(ctx context.Context, sv saving.Saving) (*saving.Saving, error) {

	if s.saving == nil {
		s.saving = &saving.Saving{
			ID:      1,
			Percent: sv.Percent,
			Date:    sv.Date,
		}
		return s.saving, nil
	}

	if !s.saving.Date.Equal(sv.Date) {
		return nil, fmt.Errorf("specified month can not be updated")
	}

	s.saving.Percent = sv.Percent

	return s.saving, nil

}

func (s *StubDataStore) RetrieveMonthSavingPercent(ctx context.Context, date time.Time) (*saving.Saving, error) {
	if s.saving == nil || !date.Equal(s.saving.Date) {
		return nil, saving.ErrNoSavingForMonth
	}

	return s.saving, nil
}

func (s *StubDataStore) RecordMonthExpenses(ctx context.Context, e expenses.Expense) (*expenses.Expense, error) {
	if s.expense == nil {
		s.expense = &expenses.Expense{
			ID:         1,
			Date:       e.Date,
			Source:     e.Source,
			Amount:     e.Amount,
			Occurrence: e.Occurrence,
		}
		return s.expense, nil
	}

	if !s.expense.Date.Equal(e.Date) {
		return nil, fmt.Errorf("specified month can not be updated")
	}

	s.expense.Amount = e.Amount
	s.expense.Occurrence = e.Occurrence
	s.expense.Source = e.Source

	return s.expense, nil
}

func (s *StubDataStore) RetrieveMonthExpenses(ctx context.Context, date time.Time) (*expenses.Expense, error) {
	if s.expense == nil || !date.Equal(s.expense.Date) {
		return nil, expenses.ErrNoExpensesForMonth
	}

	return s.expense, nil
}

func TestGetDashboardData(t *testing.T) {

	t.Run("retrieves saving, saving total and income information for a specific month", func(t *testing.T) {

		amount, _ := decimal.NewFromString("3500.00")
		i := income.Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		s := saving.Saving{
			Percent: 45,
			Date:    time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
		}

		store := &StubDataStore{income: &i, saving: &s}

		server := NewServer(store)

		request := newGetDashboardData(t, i.Date)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var dashboard Dashboard

		err := json.NewDecoder(response.Body).Decode(&dashboard)
		if err != nil {
			t.Fatalf("unable to parse response from server into income: %v", err)
		}

		if !dashboard.Income.Date.Equal(dashboard.Saving.Date) {
			t.Errorf("income and saving dates do not match got %v want %v", dashboard.Income.Date, dashboard.Saving.Date)
		}

		if !dashboard.Income.Date.Equal(i.Date) {
			t.Errorf("income date is incorrect %v want %v", dashboard.Income.Date, i.Date)
		}
		if !dashboard.Income.Amount.Equal(i.Amount) {
			t.Errorf("incorrect amount retrieved got %v want %v", dashboard.Income.Amount, i.Amount)
		}

		if dashboard.Income.Source != i.Source {
			t.Errorf("incorrect source of income retrieved got %v want %v", dashboard.Income.Source, i.Source)
		}

		if !dashboard.Saving.Date.Equal(s.Date) {
			t.Errorf("incorrect month retrieved got %v want %v", dashboard.Saving.Date, s.Date)
		}

		if dashboard.Saving.Percent != s.Percent {
			t.Errorf("incorrect percent retrieved got %v want %v", dashboard.Saving.Percent, s.Percent)
		}

	})

}

func TestPostMonthIncome(t *testing.T) {

	t.Run("checks a user can submit an income", func(t *testing.T) {

		amount, _ := decimal.NewFromString("4550.00")
		i := income.Income{
			Date:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		store := &StubDataStore{}

		server := NewServer(store)

		request := newPostRecordIncomeRequest(t, i)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var retrievedIncome income.Income

		err := json.NewDecoder(response.Body).Decode(&retrievedIncome)
		if err != nil {
			t.Fatalf("unable to parse response from server into income: %v", err)
		}

		if !retrievedIncome.Date.Equal(i.Date) {
			t.Errorf("incorrect month retrieved got %v want %v", retrievedIncome.Date, i.Date)
		}

		if !retrievedIncome.Amount.Equal(i.Amount) {
			t.Errorf("incorrect amount retrieved got %v want %v", retrievedIncome.Amount, amount)
		}

		if retrievedIncome.Source != i.Source {
			t.Errorf("incorrect source of income retrieved got %v want %v", retrievedIncome.Source, i.Source)
		}
	})

	t.Run("checks a user can update a previously submitted income", func(t *testing.T) {

		amount, _ := decimal.NewFromString("3500.00")
		i := income.Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		store := &StubDataStore{income: &i}
		server := NewServer(store)

		updatedAmount, _ := decimal.NewFromString("2600.00")
		updatedI := income.Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: updatedAmount,
		}

		request := newPostRecordIncomeRequest(t, updatedI)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var retrievedIncome income.Income

		err := json.NewDecoder(response.Body).Decode(&retrievedIncome)

		if err != nil {
			t.Fatalf("unable to parse response from server into income: %v", err)
		}

		request = newPostRecordIncomeRequest(t, i)
		response = httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		if !retrievedIncome.Date.Equal(i.Date) {
			t.Errorf("incorrect month retrieved got %v want %v", retrievedIncome.Date, i.Date)
		}

		if retrievedIncome.Amount.Equal(amount) {
			t.Errorf("amount was not updated got %v want %v", retrievedIncome.Amount, updatedAmount)
		}

		if !retrievedIncome.Amount.Equal(updatedI.Amount) {
			t.Errorf("updated amount was not returned got %v want %v", retrievedIncome.Amount, amount)
		}

		if retrievedIncome.Source != i.Source {
			t.Errorf("incorrect source of income retrieved got %v want %v", retrievedIncome.Source, i.Source)
		}

	})
}

func TestPostMonthSavingPercent(t *testing.T) {

	t.Run("checks a user can submit a saving for a month", func(t *testing.T) {

		s := saving.Saving{
			Percent: 65,
			Date:    time.Date(2022, time.May, 1, 0, 0, 0, 0, time.UTC),
		}

		store := &StubDataStore{}
		server := NewServer(store)

		request := newPostRecordSavingRequest(t, s)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var retrievedSaving saving.Saving

		err := json.NewDecoder(response.Body).Decode(&retrievedSaving)
		if err != nil {
			t.Fatalf("unable to parse response from server into saving: %v", err)
		}

		if !retrievedSaving.Date.Equal(s.Date) {
			t.Errorf("incorrect date retrieved got %v want %v", retrievedSaving.Date, s.Date)
		}

		if retrievedSaving.Percent != s.Percent {
			t.Errorf("incorrect percent retrieved got %v want %v", retrievedSaving.Percent, s.Percent)
		}
	})
}

func TestPostRecordExpense(t *testing.T) {

	t.Run("", func(t *testing.T) {

		amount, _ := decimal.NewFromString("80.00")

		e := expenses.Expense{
			Date:       time.Date(2022, time.May, 1, 0, 0, 0, 0, time.UTC),
			Source:     "gas & electric",
			Amount:     amount,
			Occurrence: "monthly",
		}

		store := &StubDataStore{expense: &e}
		server := NewServer(store)

		request := newPostRecordExpensesRequest(t, e)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var retrievedExpense expenses.Expense

		err := json.NewDecoder(response.Body).Decode(&retrievedExpense)
		if err != nil {
			t.Fatalf("unable to parse response from server into expense: %v", err)
		}

		if !retrievedExpense.Date.Equal(e.Date) {
			t.Errorf("incorrect date retrieved got %v want %v", retrievedExpense.Date, e.Date)
		}

		if retrievedExpense.Source != e.Source {
			t.Errorf("source does not match expected source got %v want %v", retrievedExpense.Source, e.Source)
		}

		if !retrievedExpense.Amount.Equal(e.Amount) {
			t.Errorf("amount does not match expected amount got %v want %v", retrievedExpense.Amount, e.Amount)
		}

		if retrievedExpense.Occurrence != e.Occurrence {
			t.Errorf("occurrence does not match expected occurrence got %v want %v", retrievedExpense.Occurrence, e.Occurrence)
		}

	})
}

func TestDeleteIncome(t *testing.T) {
	t.Run("deletes a specific income from the datestore and returns 200", func(t *testing.T) {
		incomes := []income.Income{{ID: 1}}
		store := &StubDataStore{incomes: incomes}
		server := NewServer(store)

		request := newDeleteIncomeRequest(t, 1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if len(store.deleteCount) != 1 {
			t.Errorf("got %d calls to DeleteIncome want %d", len(store.deleteCount), 1)
		}
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetDashboardData(t testing.TB, date time.Time) *http.Request {

	dateString := date.Format("2006-01-02")

	request, err := http.NewRequest(http.MethodGet, "/api/v1/budget/dashboard", nil)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	requestQuery := request.URL.Query()
	requestQuery.Set("date", dateString)
	request.URL.RawQuery = requestQuery.Encode()

	return request
}

func newPostRecordIncomeRequest(t testing.TB, income income.Income) *http.Request {

	bodyStr, err := json.Marshal(&income)
	if err != nil {
		t.Fatalf("fail to marshal user information: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, "/api/v1/budget/income", bytes.NewBuffer(bodyStr))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	return request
}

func newPostRecordSavingRequest(t testing.TB, saving saving.Saving) *http.Request {

	bodyStr, err := json.Marshal(&saving)
	if err != nil {
		t.Fatalf("fail to marshal user information: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, "/api/v1/budget/saving", bytes.NewBuffer(bodyStr))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	return request
}

func newPostRecordExpensesRequest(t testing.TB, expense expenses.Expense) *http.Request {

	bodyStr, err := json.Marshal(&expense)
	if err != nil {
		t.Fatalf("fail to marshal user information: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, "/api/v1/budget/expenses", bytes.NewBuffer(bodyStr))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	return request
}

func newDeleteIncomeRequest(t testing.TB, ID int) *http.Request {
	url := fmt.Sprintf("/api/v1/income/%d", ID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	return req
}
