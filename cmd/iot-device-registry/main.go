package main

import (
	"strings"

	"github.com/diwise/iot-device-registry/internal/pkg/application"
	"github.com/diwise/iot-device-registry/internal/pkg/infrastructure/repositories/database"
	"github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/rs/zerolog/log"
)

func main() {

	serviceName := "iot-device-registry"

	logger := log.With().Str("service", strings.ToLower(serviceName)).Logger()
	logger.Info().Msg("starting up ...")

	config := messaging.LoadConfiguration(serviceName, logger)
	messenger, _ := messaging.Initialize(config)

	defer messenger.Close()

	db, _ := database.NewDatabaseConnection(database.NewPostgreSQLConnector(logger))
	application.CreateRouterAndStartServing(logger, messenger, db)
}
