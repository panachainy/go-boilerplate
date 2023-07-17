package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type (
	Product struct {
		// ID       string `json:"id"`
		// Username string `json:"username"`
	}

	ProductEntity struct {
		// ID       string
		// Username string
		// Password string
	}

	ProductRepository interface {
		FetchByUsername(ctx context.Context, username string) (*ProductEntity, error)
	}

	ProductService interface {
		FetchByUsername(ctx context.Context, username string) (*Product, error)
	}

	ProductHandler interface {
		GetProducts() func(*fiber.Ctx) error
	}
)
