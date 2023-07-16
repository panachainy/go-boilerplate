//go:build wireinject
// +build wireinject

//go:generate wire
package app

import (
	"go-boilerplate/cmd/config"

	"github.com/google/wire"
)

func Wire() (*Application, error) {
	wire.Build(
		Provide,
		config.ProviderSet,
	)

	return &Application{}, nil
}
