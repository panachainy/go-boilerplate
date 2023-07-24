package main

import (
	"fmt"
	"go-boilerplate/cmd/app"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	application, err := app.Wire()
	if err != nil {

		log.Fatalf("Failed auto injection to initialize application: %v", err)
	}

	application.Server.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "OK",
		})
	})

	log.Fatal(application.Server.Listen(fmt.Sprintf(":%v", application.Config.APP_PORT)))
}
