package main

import (
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)



func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("Starting up")

	sqlString := getDatabaseString()

	var err error
	var db *gorm.DB

	db, err = gorm.Open("mysql", sqlString)
	if err != nil {
		log.Fatal().Msgf("Unable to connect to '%s'", sqlString)
	}
	db = db.BlockGlobalUpdate(true)
	defer func() {
		_ = db.Close()
	}()

	hostname, err := os.Hostname()
	if err != nil {
		log.Error().Msgf("Error calling os.Hostname(): %s", err)
	} else {
		log.Info().Msgf("Hostname is %s", hostname)
	}

	log.Info().Msg("Shutting down")
}
