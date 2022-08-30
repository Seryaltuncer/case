package blog

import (
	"deneme/internal/model"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Store struct {
	sync.Mutex
	BlogPost []model.BlogPost
}

func CreateBlogStore() *Store {
	return new(Store)
}

func (s *Store) CreateBlogPost(a model.BlogPost) {
	s.Lock()
	defer s.Unlock()

	// Getting random id
	st := make([]rune, 5)
	for i := range st {
		st[i] = letters[rand.Intn(len(letters))]
	}
	a.Id = string(st)
	t := time.Now()
	a.CreatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	s.BlogPost = append(s.BlogPost, a)
}

func (s *Store) ListBlogPost() []model.BlogPost {
	s.Lock()
	defer s.Unlock()
	return s.BlogPost
}

func (s *Store) GetBlogPost(id string) model.BlogPost {
	s.Lock()
	defer s.Unlock()
	for _, post := range s.BlogPost {
		if post.Id == id {
			return post
		}
	}
	return model.BlogPost{}
}
