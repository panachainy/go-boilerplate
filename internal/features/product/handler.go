package product

import (
	"go-boilerplate/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc domain.ProductService
}

func (h *handler) GetProducts() func(*fiber.Ctx) error {
	panic("implement me")
}
