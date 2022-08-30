package author

import (
	"deneme/internal/model"
	"sync"
)

type Store struct {
	sync.Mutex
	Author []model.Author
}

func CreateAuthorStore() *Store {
	return new(Store)
}

func (s *Store) CreateAuthor(a model.Author) {
	s.Lock()
	defer s.Unlock()
	s.Author = append(s.Author, a)
}

func (s *Store) ListAuthor() []model.Author {
	s.Lock()
	defer s.Unlock()
	return s.Author
}

func (s *Store) GetAuthor(name string) model.Author {
	s.Lock()
	defer s.Unlock()
	for _, auth := range s.Author {
		if auth.Name == name {
			return auth
		}
	}
	return model.Author{}
}
