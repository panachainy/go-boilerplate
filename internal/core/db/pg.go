package db

import (
	"fmt"
	"go-boilerplate/internal/core/config"
	"go-boilerplate/internal/core/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	dbConn *gorm.DB
	config *config.Configuration
	log    *logrus.Logger
}

func (db *database) Initial() {
	var dsn string
	if db.config.DATABASE_DSN != "" {
		db.log.Infoln("[DB-CONFIG] Use config DATABASE_DSN")
		dsn = db.config.DATABASE_DSN
	} else {
		db.log.Infoln("[DB-CONFIG] Use split config DATABASE")

		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			db.config.DATABASE_HOST,
			db.config.DATABASE_USER,
			db.config.DATABASE_PASSWORD,
			db.config.DATABASE_NAME,
			db.config.DATABASE_PORT,
			db.config.DATABASE_SSLMODE,
			db.config.DATABASE_TIMEZONE)
	}
	var err error
	db.dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.migrate()
}

func (db *database) migrate() {
	err := db.dbConn.AutoMigrate(&domain.ProductEntity{})
	if err != nil {
		panic("Failed to auto migrate database")
	}
}
