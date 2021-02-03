package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/expenses"
	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/catzkorn/pay-yourself-first/internal/saving"
	"github.com/shopspring/decimal"
)

// Server is the HTTP interface
type Server struct {
	dataStore DataStore
	router    *http.ServeMux
}

// Dashboard defines the information needed for the budget dashboard
type Dashboard struct {
	Saving      *saving.Saving
	SavingTotal decimal.Decimal
	Income      *income.Income
	Expense     *expenses.Expense
}

// DataStore defines the interface to persist data
type DataStore interface {
	RecordIncome(ctx context.Context, i income.Income) (*income.Income, error)
	ListIncomes(ctx context.Context) ([]income.Income, error)
	RetrieveMonthIncome(ctx context.Context, date time.Time) (*income.Income, error)
	DeleteIncome(ctx context.Context, id uint32) error
	RecordMonthSavingPercent(ctx context.Context, s saving.Saving) (*saving.Saving, error)
	RetrieveMonthSavingPercent(ctx context.Context, date time.Time) (*saving.Saving, error)
	RecordMonthExpenses(ctx context.Context, e expenses.Expense) (*expenses.Expense, error)
	RetrieveMonthExpenses(ctx context.Context, date time.Time) (*expenses.Expense, error)
}

// NewServer returns an instance of a Server
func NewServer(dataStore DataStore) *Server {

	s := &Server{dataStore: dataStore, router: http.NewServeMux()}

	s.router.Handle("/income", http.HandlerFunc(s.incomeHandler))
	s.router.Handle("/api/v1/budget/income", http.HandlerFunc(s.budgetIncomeHandler))
	s.router.Handle("/api/v1/budget/saving", http.HandlerFunc(s.budgetSavingHandler))
	s.router.Handle("/api/v1/budget/dashboard", http.HandlerFunc(s.budgetDashboardHandler))
	s.router.Handle("/api/v1/budget/expenses", http.HandlerFunc(s.budgetExpensesHandler))
	s.router.Handle("/", http.FileServer(http.Dir("web")))

	return s
}

// ServeHTTP implements the http handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// IncomeHandler handles the routing logic for '/'
func (s *Server) incomeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		s.processListIncome(r.Context(), w)
	}
}

// processListIncome processes the get '/' and returns the income slice in JSON format
func (s *Server) processListIncome(ctx context.Context, w http.ResponseWriter) {

	incomes, err := s.dataStore.ListIncomes(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(incomes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) budgetDashboardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getBudgetDashboardData(w, r)
	}
}

func (s *Server) budgetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postBudgetExpenses(w, r)
	}
}

func (s *Server) postBudgetExpenses(w http.ResponseWriter, r *http.Request) {

	var expense expenses.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnedExpense, err := s.dataStore.RecordMonthExpenses(r.Context(), expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(returnedExpense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getBudgetDashboardData(w http.ResponseWriter, r *http.Request) {

	dateLayout := "2006-01-02"
	date := r.URL.Query().Get("date")
	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthSaving, err := s.dataStore.RetrieveMonthSavingPercent(r.Context(), parsedDate)

	switch {
	case err == saving.ErrNoSavingForMonth:
		monthSaving = &saving.Saving{
			Date: parsedDate,
		}
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthIncome, err := s.dataStore.RetrieveMonthIncome(r.Context(), parsedDate)

	switch {
	case err == income.ErrNoIncomeForMonth:
		monthIncome = &income.Income{
			Date: parsedDate,
		}
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	percent := decimal.NewFromInt(int64(monthSaving.Percent))
	oneHundred := decimal.NewFromInt(100)
	percent = percent.Div(oneHundred)

	monthSavingValue := monthIncome.Amount.Mul(percent)

	monthExpense, err := s.dataStore.RetrieveMonthExpenses(r.Context(), parsedDate)

	switch {
	case err == expenses.ErrNoExpensesForMonth:
		monthExpense = &expenses.Expense{
			Date: parsedDate,
		}
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dashboard := Dashboard{
		Saving:      monthSaving,
		SavingTotal: monthSavingValue,
		Income:      monthIncome,
		Expense:     monthExpense,
	}

	err = json.NewEncoder(w).Encode(dashboard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) budgetIncomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getMonthIncome(w, r)
	case http.MethodPost:
		s.postMonthIncome(w, r)
	}
}

func (s *Server) budgetSavingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getMonthSaving(w, r)
	case http.MethodPost:
		s.postMonthSaving(w, r)
	}
}

func (s *Server) postMonthSaving(w http.ResponseWriter, r *http.Request) {

	var saving saving.Saving

	err := json.NewDecoder(r.Body).Decode(&saving)
	fmt.Println(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnedSaving, err := s.dataStore.RecordMonthSavingPercent(r.Context(), saving)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(returnedSaving)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getMonthSaving(w http.ResponseWriter, r *http.Request) {

	dateLayout := "2006-01-02"
	date := r.URL.Query().Get("date")
	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthSavingPercent, err := s.dataStore.RetrieveMonthSavingPercent(r.Context(), parsedDate)

	switch {
	case err == saving.ErrNoSavingForMonth:
		monthSavingPercent = &saving.Saving{
			Date: parsedDate,
		}
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(monthSavingPercent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getMonthIncome(w http.ResponseWriter, r *http.Request) {
	dateLayout := "2006-01-02"
	date := r.URL.Query().Get("date")
	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthIncome, err := s.dataStore.RetrieveMonthIncome(r.Context(), parsedDate)

	switch {
	case err == income.ErrNoIncomeForMonth:
		monthIncome = &income.Income{
			Date: parsedDate,
		}
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(monthIncome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) postMonthIncome(w http.ResponseWriter, r *http.Request) {
	var income income.Income

	err := json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnedIncome, err := s.dataStore.RecordIncome(r.Context(), income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(returnedIncome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
