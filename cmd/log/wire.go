//go:build wireinject
// +build wireinject

//go:generate wire
package log

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var ProviderSet = wire.NewSet(
	Provide,
)

func Wire() (*logrus.Logger, error) {
	panic(wire.Build(ProviderSet))
}
