// Package store provides a thread-safe key-value store.
package store

// Store is an in-memory key-value data store.
type Store struct {
	data map[string]string
}

// NewStore returns a ready-to-use Store.
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Get returns the value for key k. The second return value is false
// if the key was not found.
func (s *Store) Get(k string) (string, bool) {
	v, ok := s.data[k]
	return v, ok
}

// Put stores value v under key k.
func (s *Store) Put(k, v string) {
	s.data[k] = v
}

// Delete removes key k and returns true. Returns false if the key
// was not found.
func (s *Store) Delete(k string) bool {
	_, ok := s.data[k]
	delete(s.data, k)
	return ok
}
