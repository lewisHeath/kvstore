package server

import (
	"fmt"
	"log"
	"net/http"
)

// ServeHTTP registers HTTP routes and starts the HTTP server on the given port.
func (s *Server) ServeHTTP(port string) {
	log.Printf("Initialising HTTP server on port %v\n", port)

	http.HandleFunc("GET /{key}", s.handleGET)
	http.HandleFunc("PUT /{key}/{value}", s.handlePUT)
	http.HandleFunc("DELETE /{key}", s.handleDELETE)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	log.Printf("Performing HTTP GET on key=%v\n", key)
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
	log.Printf("Performing HTTP PUT on key=%v value=%v\n", key, value)
	s.store.Put(key, value)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDELETE(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	log.Printf("Performing DELETE on key=%v\n", key)
	s.store.Delete(key)
	w.WriteHeader(http.StatusOK)
}
