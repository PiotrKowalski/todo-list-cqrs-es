package main

import (
	"todo-list-cqrs-es/internal/adapters/rest"
	"todo-list-cqrs-es/internal/app"
	"todo-list-cqrs-es/pkg/config"
)

var (
	mongoURI, _ = config.ReadEnvString("MONGODB_URI")
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	application := app.NewApplication(app.Config{MongoDBURI: mongoURI})
	service := rest.NewRESTService(application)

	err := service.Run()
	if err != nil {
		return err
	}
	return nil
}
