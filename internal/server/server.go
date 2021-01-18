package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/database"
	"github.com/catzkorn/pay-yourself-first/internal/income"
)

// Server is the HTTP interface
type Server struct {
	database *database.Database
	router   *http.ServeMux
}

// Database defines the interface to persist data
type Database interface {
	RecordIncome(ctx context.Context, i income.Income) (*income.Income, error)
	ListIncomes(ctx context.Context) ([]income.Income, error)
	GetMonthIncome(ctx context.Context, date time.Time) (*income.Income, error)
	DeleteIncome(ctx context.Context, id uint32) error
}

// NewServer returns an instance of a Server
func NewServer(database *database.Database) *Server {

	s := &Server{database: database, router: http.NewServeMux()}

	s.router.Handle("/income", http.HandlerFunc(s.incomeHandler))
	s.router.Handle("/", http.HandlerFunc(s.budgetHandler))

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

	incomes, err := s.database.ListIncomes(ctx)

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

func (s *Server) budgetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getMonthIncome(r.Context(), w, r)
	case http.MethodPost:
		s.postMonthIncome(w, r)
	}
}

func (s *Server) getMonthIncome(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	dateLayout := "2006-01-02"
	date := r.URL.Query().Get("date")
	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monthIncome, err := s.database.GetMonthIncome(ctx, parsedDate)

	switch {
	case err == database.ErrNoIncomeForMonth:
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

	returnedIncome, err := s.database.RecordIncome(r.Context(), income)
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
