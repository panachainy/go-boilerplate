package product

import (
	"context"
	"database/sql"
	"go-boilerplate/internal/core/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (*domain.ProductEntity, error) {
	panic("implement me")
}
