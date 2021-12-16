package main

import (
	"domain-driven-design-layout/infrastructure/builder"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/http"
	"log"
)

func main() {
	// Load application config (env vars)
	appConfig := config.LoadAppConfig()

	repositories, err := builder.CreateRepositories(appConfig.RepositoriesConfig)
	if err != nil {
		log.Fatal(err)
	}

	actions, err := builder.CreateActions(repositories)
	if err != nil {
		log.Fatal(err)
	}

	handlers, err := http.CreateHttpHandlers(actions)
	if err != nil {
		log.Fatal(err)
	}

	webServer, err := http.NewWebServer(appConfig.WebConfig, handlers)
	if err != nil {
		log.Fatal(err)
	}

	_ = webServer.Start()
}
