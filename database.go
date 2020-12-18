package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func getDatabaseString() string {
	sqlUsername := os.Getenv("MYSQL_USER")
	sqlPassword := os.Getenv("MYSQL_PASS")
	sqlHost := "127.0.0.1"
	sqlPort := 33060

	log.Debug().Msgf("The value of sqlUsername is '%s'", sqlUsername)
	log.Debug().Msgf("The value of sqlPassword is '%s'", sqlPassword)
	log.Debug().Msgf("The value of sqlHost is '%s'", sqlHost)
	log.Debug().Msgf("The value of sqlPort is (%T) '%d'", sqlPort, sqlPort)
	stmt :=fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Local", sqlUsername, sqlPassword, sqlHost, sqlPort)
	log.Debug().Msgf("Statement: '%s'", stmt)

	return stmt
}