package product

import (
	"go-boilerplate/internal/core/domain"

	"context"
)

type service struct {
	repo domain.ProductRepository
}

func (s *service) FetchByUsername(ctx context.Context, username string) (*domain.Product, error) {
	panic("implement me")
}
