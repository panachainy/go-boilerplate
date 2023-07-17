// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"go-boilerplate/cmd/log"
	"go-boilerplate/internal/core/config"
)

// Injectors from wire.go:

func Wire() (*Application, error) {
	configuration := config.Provide()
	logger := log.Provide()
	application := Provide(configuration, logger)
	return application, nil
}
