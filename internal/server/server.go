package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/catzkorn/pay-yourself-first/internal/saving"
)

// Server is the HTTP interface
type Server struct {
	dataStore DataStore
	router    *http.ServeMux
}

// DataStore defines the interface to persist data
type DataStore interface {
	RecordIncome(ctx context.Context, i income.Income) (*income.Income, error)
	ListIncomes(ctx context.Context) ([]income.Income, error)
	RetrieveMonthIncome(ctx context.Context, date time.Time) (*income.Income, error)
	DeleteIncome(ctx context.Context, id uint32) error
	RecordMonthSavingPercent(ctx context.Context, s saving.Saving) (*saving.Saving, error)
	RetrieveMonthSavingPercent(ctx context.Context, date time.Time) (*saving.Saving, error)
}

// NewServer returns an instance of a Server
func NewServer(dataStore DataStore) *Server {

	s := &Server{dataStore: dataStore, router: http.NewServeMux()}

	s.router.Handle("/income", http.HandlerFunc(s.incomeHandler))
	s.router.Handle("/api/v1/budget/income", http.HandlerFunc(s.budgetIncomeHandler))
	s.router.Handle("/api/v1/budget/saving", http.HandlerFunc(s.budgetSavingHandler))
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
		s.getMonthSavingIncome(w, r)
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

func (s *Server) getMonthSavingIncome(w http.ResponseWriter, r *http.Request) {

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
