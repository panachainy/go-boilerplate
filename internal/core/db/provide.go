package db

import (
	"go-boilerplate/internal/core/config"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	db     *database
	dbOnce sync.Once
)

func Provide(config *config.Configuration, log *logrus.Logger) *database {
	dbOnce.Do(func() {
		db = &database{config: config, log: log}
		db.Initial()
	})

	return db
}
