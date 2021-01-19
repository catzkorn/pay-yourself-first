package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/shopspring/decimal"
)

//ToDo build out testing

type StubDatabase struct {
	income      income.Income
	incomes     []income.Income
	deleteCount []uint32
}

func (s *StubDatabase) RetrieveMonthIncome(_ context.Context, date time.Time) (*income.Income, error) {

	if date != s.income.Date {
		return nil, income.ErrNoIncomeForMonth
	}

	return &s.income, nil
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

func (s *StubDatabase) DeleteIncome(_ context.Context, id uint32) error {
	s.deleteCount = append(s.deleteCount, id)
	return nil
}

func TestGetMonthIncome(t *testing.T) {

	t.Run("retrieves an income for a specific month", func(t *testing.T) {

		amount, _ := decimal.NewFromString("3500.00")
		i := income.Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		store := &StubDatabase{income: i}

		server := NewServer(store)

		request := newMonthIncomeRequest(t, i.Date)
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
			t.Errorf("incorrect amount retrieved got %v want %v", retrievedIncome.Amount, i.Amount)
		}

		if retrievedIncome.Source != i.Source {
			t.Errorf("incorrect source of income retrieved got %v want %v", retrievedIncome.Source, i.Source)
		}
	})

	t.Run("returns an income struct with only date if data for that month does not yet exist", func(t *testing.T) {

		date := time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC)

		store := &StubDatabase{}

		server := NewServer(store)

		request := newMonthIncomeRequest(t, date)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var retrievedIncome income.Income

		err := json.NewDecoder(response.Body).Decode(&retrievedIncome)
		if err != nil {
			t.Fatalf("unable to parse response from server into income: %v", err)
		}

		if !retrievedIncome.Date.Equal(date) {
			t.Errorf("incorrect month retrieved got %v want %v", retrievedIncome.Date, date)
		}

		amount, _ := decimal.NewFromString("0.00")

		if !retrievedIncome.Amount.Equal(amount) {
			t.Errorf("incorrect amount retrieved got %v want %v", retrievedIncome.Amount, amount)
		}

		if retrievedIncome.Source != "" {
			t.Errorf("incorrect source of income retrieved got %v want %v", retrievedIncome.Source, "")
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

		store := &StubDatabase{}

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

	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newMonthIncomeRequest(t testing.TB, date time.Time) *http.Request {

	dateString := date.Format("2006-01-02")

	request, err := http.NewRequest(http.MethodGet, "/", nil)
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

	request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(bodyStr))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	return request
}
