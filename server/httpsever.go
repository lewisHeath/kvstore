package server

import (
	"fmt"
	"net/http"
)

func (s *Server) ServeHTTP(port string) {
	fmt.Printf("Initialising HTTP server on port %v\n", port)

	http.HandleFunc("GET /{key}", s.handleGET)
	http.HandleFunc("PUT /{key}/{value}", s.handlePUT)
	http.HandleFunc("DELETE /{key}", s.handleDELETE)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	fmt.Printf("Performing HTTP GET on key=%v\n", key)
	value, ok := s.store.Get(key)
	if !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, value)
}

func (s *Server) handlePUT(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	value := r.PathValue("value")
	fmt.Printf("Performing HTTP PUT on key=%v value=%v\n", key, value)
	s.store.Put(key, value)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDELETE(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	fmt.Printf("Performing DELETE on key=%v\n", key)
	s.store.Delete(key)
	w.WriteHeader(http.StatusOK)
}
