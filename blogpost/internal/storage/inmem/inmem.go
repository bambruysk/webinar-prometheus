package inmem

import (
	"context"
	"sync"

	"blogpost/internal/models"
)

type Storage struct {
	data map[string]*models.BlogPost
	mtx  sync.RWMutex
}

func New() *Storage {
	return &Storage{
		data: make(map[string]*models.BlogPost),
	}
}

func (s *Storage) Create(ctx context.Context, post *models.BlogPost) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.data[post.ID]; ok {
		return models.ErrorAlreadyExists
	}

	s.data[post.ID] = post

	return nil
}

func (s *Storage) Get(ctx context.Context, id string) (*models.BlogPost, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	post, ok := s.data[id]
	if !ok {
		return nil, models.ErrorPostNotFound
	}

	return post, nil
}
