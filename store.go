package main

import (
	"fmt"
	"strings"
)

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(k string) (string, error) {
	v, ok := s.data[k]
	if !ok {
		return "NOTFOUND", fmt.Errorf("key %v not found", k)
	}
	return v, nil
}

func (s *Store) Put(k, v string) string {
	s.data[k] = v
	return "OK"
}

func (s *Store) Delete(k string) string {
	delete(s.data, k)
	return "OK"
}

func handle(c string, kv *Store) (string, error) {
	s := strings.Split(c, " ")
	switch s[0] {
	case "GET":
		fmt.Printf("Performing GET on key=%v\n", s[1])
		return kv.Get(s[1])
	case "PUT":
		fmt.Printf("Performing PUT with key=%v value=%v\n", s[1], s[2])
		return kv.Put(s[1], s[2]), nil
	case "DELETE":
		fmt.Printf("Performing DELETE on key=%v\n", s[1])
		return kv.Delete(s[1]), nil
	default:
		return "UNRECOGNIZED", fmt.Errorf("unrecognized command %v", s[0])
	}
}
