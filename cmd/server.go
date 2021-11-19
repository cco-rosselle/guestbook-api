package main

import (
	"home/zellie/Code/guestbook-api/internal/repos"
	"home/zellie/Code/guestbook-api/internal/routes"
	"home/zellie/Code/guestbook-api/internal/services"
	"home/zellie/Code/guestbook-api/internal/settings"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func configureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level := settings.Default().Logging.Level
	switch level {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		log.Warn().
			Str("package", "main").
			Str("logging.level", level).
			Msg("logging.level is not a known value")

		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	log.Debug().
		Str("package", "main").
		Str("logging.level", level).
		Msg("logging level set")
}

func createGinEngine() *gin.Engine {
	r := gin.New()

	// attach logger and recovery
	r.Use(ginzerolog.Logger("gin"))
	r.Use(gin.Recovery())

	return r
}

func main() {
	// load settings
	if err := settings.Load(); err != nil {
		log.Fatal().
			Str("package", "main").
			Stack().Err(err).
			Msg("unable to load settings")
		return
	}

	// configure logger
	configureLogger()

	s := settings.Default()
	log.Info().
		Str("Server.Address", s.Server.Address).
		Msg("initializing server")

	r := createGinEngine()

	cr, err := repos.NewCommentsRepo()
	if err != nil {
		log.Fatal().
			Str("package", "main").
			Err(err).Msg("unable to create comments repo")
		return
	}

	// TODO: create/initialize token handler

	// initialize services
	cs, err := services.NewCommentsService(cr) // doesn't take anything yet
	if err != nil {
		log.Fatal().
			Str("package", "main").
			Err(err).Msg("unable to create comments service")
		return
	}

	// intialize api controllers
	routes.Register(r, cs)

	// start the gin server
	r.Run(s.Server.Address)

	log.Info().
		Str("Server.Address", s.Server.Address).
		Msg("server started")

}
