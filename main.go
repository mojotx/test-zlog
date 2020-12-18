package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("Starting up")

	hostname, err := os.Hostname()
	if err != nil {
		log.Error().Msgf("Error calling os.Hostname(): %s", err)
	} else {
		log.Info().Msgf("Hostname is %s", hostname)
	}

	log.Info().Msg("Shutting down")
}
