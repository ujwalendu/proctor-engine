package server

import (
	"time"

	"github.com/gojekfarm/proctor-engine/config"
	"github.com/gojekfarm/proctor-engine/logger"

	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
)

func Start() error {
	appPort := ":" + config.AppPort()

	server := negroni.New(negroni.NewRecovery())
	server.UseHandler(router)

	logger.Info("Starting server on port", appPort)

	graceful.Run(appPort, 2*time.Second, server)

	logger.Info("Stopped server")
	return nil
}
