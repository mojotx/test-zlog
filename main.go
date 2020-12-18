package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

func getDatabaseString() string {
	sqlUsername := os.Getenv("MYSQL_USER")
	sqlPassword := os.Getenv("MYSQL_PASS")
	sqlHost := "127.0.0.1"
	sqlPort := 33060

	log.Debug().Msgf("The value of sqlUsername is '%s'", sqlUsername)
	log.Debug().Msgf("The value of sqlPassword is '%s'", sqlPassword)
	log.Debug().Msgf("The value of sqlHost is '%s'", sqlHost)
	log.Debug().Msgf("The value of sqlPort is (%T) '%d'", sqlPort, sqlPort)
	stmt := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Local", sqlUsername, sqlPassword, sqlHost, sqlPort)
	log.Debug().Msgf("Statement: '%s'", stmt)

	return stmt
}
