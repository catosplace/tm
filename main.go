package main

import (
	"github.com/adlio/trello"
	log "github.com/sirupsen/logrus"
)

type trelloClientConfig struct {
	appKey string
	token  string
	logger *log.Logger
}

func main() {
	// Get a Trello API Logger
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	// Get a new Trello Client
	getClientInstance(&trelloClientConfig{
		appKey: "",
		token:  "",
		logger: logger,
	})
}

func getClientInstance(config *trelloClientConfig) *trello.Client {
	client := trello.NewClient("", "") // (appKey, token)
	if config.logger != nil {
		client.Logger = config.logger
	}
	return client
}
