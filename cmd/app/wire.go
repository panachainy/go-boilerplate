//go:build wireinject
// +build wireinject

//go:generate wire
package app

import (
	"go-boilerplate/cmd/log"
	"go-boilerplate/internal/core/config"

	"github.com/google/wire"
)

func Wire() (*Application, error) {
	wire.Build(
		Provide,
		config.ProviderSet,
		log.ProviderSet,
	)

	return &Application{}, nil
}
