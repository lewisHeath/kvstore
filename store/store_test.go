package store

import "testing"

func TestGet(t *testing.T) {
	s := Store{
		data: map[string]string{
			"hello": "world",
		},
	}
	r, ok := s.Get("hello")
	if !ok {
		t.Errorf("key not present in store")
	}
	if r != "world" {
		t.Errorf("got %v", r)
	}
}
