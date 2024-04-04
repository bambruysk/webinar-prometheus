package blog

import (
	"context"

	"blogpost/internal/models"
)

type Storage interface {
	Create(ctx context.Context, post *models.BlogPost) error
	Get(ctx context.Context, id string) (*models.BlogPost, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Create(ctx context.Context, post *models.BlogPost) (*models.BlogPost, error) {
	post.ID = models.NewID()

	if err := s.storage.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Service) Get(ctx context.Context, id string) (*models.BlogPost, error) {
	return s.storage.Get(ctx, id)
}
