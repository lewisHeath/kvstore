package store

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(k string) (string, bool) {
	v, ok := s.data[k]
	return v, ok
}

func (s *Store) Put(k, v string) {
	s.data[k] = v
}

func (s *Store) Delete(k string) bool {
	_, ok := s.data[k]
	delete(s.data, k)
	return ok
}
