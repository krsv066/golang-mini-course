package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	var app = fiber.New()

	app.Post("/account", CreateAccount)
	app.Get("/account/:name", GetAccount)
	app.Patch("/account/:name", UpdateAmount)
	app.Put("/account/:name", UpdateName)
	app.Delete("/account/:name", DeleteAccount)

	log.Fatal(app.Listen(":8000"))
}
