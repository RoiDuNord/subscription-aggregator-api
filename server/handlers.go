package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := s.manager.DeleteSubscription(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetSum(w http.ResponseWriter, r *http.Request) {
	
}