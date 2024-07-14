package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"proj/hw4/dto"
)

func CreateAccount(c *fiber.Ctx) error {
	account := new(dto.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request."})
	}

	if account.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account name is required."})
	}

	if account.Balance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account balance must be >= 0."})
	}

	ctx := context.Background()

	db := dto.GetDB()
	if dto.AccountExists(account.Name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account already exists."})
	}

	_, err := db.ExecContext(ctx, "INSERT INTO accounts (name, balance) VALUES ($1, $2)", account.Name, account.Balance)

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Account created."})
}

func GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	ctx := context.Background()

	db := dto.GetDB()
	if !dto.AccountExists(name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist."})
	}

	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts WHERE name = $1", name)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	account := new(dto.Account)

	for rows.Next() {
		if err := rows.Scan(&account.Name, &account.Balance); err != nil {
			log.Fatal(err)
		}
	}

	return c.JSON(account)
}

func UpdateName(c *fiber.Ctx) error {
	oldName := c.Params("name")
	newName := new(dto.ChangeNameParams)

	if err := c.BodyParser(newName); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request."})
	}

	ctx := context.Background()
	db := dto.GetDB()

	if dto.AccountExists(newName.NewName, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account already exists."})
	}

	if !dto.AccountExists(oldName, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exists."})
	}

	_, err := db.ExecContext(ctx, "UPDATE accounts SET name = $1 WHERE name = $2", newName.NewName, oldName)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{"message": "Account name updated"})
}

func UpdateBalance(c *fiber.Ctx) error {
	name := c.Params("name")
	balance := new(dto.UpdateBalanceParams)

	if err := c.BodyParser(balance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request."})
	}

	if balance.Balance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account balance must be >= 0."})
	}

	ctx := context.Background()

	db := dto.GetDB()
	if !dto.AccountExists(name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist."})
	}

	_, err := db.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE name = $2", balance.Balance, name)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{"message": "Account balance updated."})
}

func DeleteAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	ctx := context.Background()
	db := dto.GetDB()

	if !dto.AccountExists(name, db) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Account does not exist."})
	}

	_, err := db.ExecContext(ctx, "DELETE FROM accounts WHERE name = $1", name)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{"message": "Account deleted."})
}
