package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"proj/hw4/dto"
)

func main() {
	if err := dto.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/account", CreateAccount)
	app.Get("/account/:name", GetAccount)
	app.Patch("/account/:name", UpdateBalance)
	app.Put("/account/:name", UpdateName)
	app.Delete("/account/:name", DeleteAccount)

	log.Fatal(app.Listen(":8000"))
}
