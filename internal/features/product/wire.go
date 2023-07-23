//go:build wireinject
// +build wireinject

//go:generate wire
package product

import (
	"database/sql"
	"go-boilerplate/internal/core/domain"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideHandler,
	ProvideService,
	ProvideRepository,

	wire.Bind(new(domain.ProductHandler), new(*handler)),
	wire.Bind(new(domain.ProductService), new(*service)),
	wire.Bind(new(domain.ProductRepository), new(*repository)),
)

func Wire(db *sql.DB) *handler {
	panic(wire.Build(ProviderSet))
}
