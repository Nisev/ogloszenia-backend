package main

import (
	_ "net/http"
	"os"

	_ "github.com/nifrez/ogloszenia/users"

	"github.com/rs/zerolog/log"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/nifrez/ogloszenia/common"
	"github.com/nifrez/ogloszenia/users"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	db := common.InitDB()
	users.AutoMigrate()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("/api/v1")
	users.RouterRegister(v1.Group("/users"))

	ping := router.Group("/api/ping")
	ping.GET("/", pingRoute)

	log.Info().Msg("Starting http server on port 63636")
	err := router.Run(":63636")
	handleError(err, 666, "Problem during starting http server.")
}

func pingRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleError(err error, exitCode int, msg string) {
	if err != nil {
		log.Error().Err(err).Msg(msg)
		os.Exit(exitCode)
	}
}
