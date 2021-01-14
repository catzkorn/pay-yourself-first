package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/catzkorn/pay-yourself-first/internal/database"
)

// Server is the HTTP interface
type Server struct {
	database *database.Database
	router   *http.ServeMux
}

// NewServer returns an instance of a Server
func NewServer(database *database.Database) *Server {

	s := &Server{database: database, router: http.NewServeMux()}

	s.router.Handle("/", http.HandlerFunc(s.incomeHandler))

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
