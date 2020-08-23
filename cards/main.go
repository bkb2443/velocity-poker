package main

import (
	"cards/handler"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("cards"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(new(handler.Cards))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
