package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	Driver             = "postgres"
	Host               = "localhost"
	Port               = "5432"
	User               = "postgres"
	DbName             = "ogloszenia"
	Password           = "postgres"
	MaxIdleConnections = 10
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	log.Info().Msg("Trying connect to DB.")
	db, err := gorm.Open(getDbDriver(), getDatabaseConnectionConfiguration())
	handleConnectionDatabaseError(err)

	db.DB().SetMaxIdleConns(MaxIdleConnections)
	log.Info().Msg("Connected to DB.")

	DB = db
	return DB
}

func handleConnectionDatabaseError(err error) {
	if err != nil {
		log.Error().Err(err).Msg("Couldn't connect to DB.")
		os.Exit(10)
	}
}

func getDatabaseConnectionConfiguration() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", Host, Port, User, DbName, Password)
}

func getDbDriver() string {
	return Driver
}

func GetDB() *gorm.DB {
	return DB
}
