package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"proj/hw2/dto"
)

func CreateAccount(c *fiber.Ctx) error {
	account := new(dto.Account)
	err := c.BodyParser(account)
	if err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if account.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account name is required"})
	}

	if account.Balance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account amount must be >= 0"})
	}

	dto.AllAccounts.CreateAccount(account)

	return c.Status(fiber.StatusCreated).JSON(account)
}

func GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")
	account, err := dto.AllAccounts.GetAccount(name)

	if err != nil {
		log.Printf("Error fetching account: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	return c.JSON(account)
}

func UpdateAmount(c *fiber.Ctx) error {
	name := c.Params("name")
	balance := new(dto.UpdateBalanceParams)

	err := c.BodyParser(balance)
	if err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if balance.Balance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account amount must be >= 0"})
	}

	account := &dto.Account{Name: name, Balance: balance.Balance}

	err = dto.AllAccounts.UpdateAmount(account)

	if err != nil {
		log.Printf("Error updating account amount: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func UpdateName(c *fiber.Ctx) error {
	oldName := c.Params("name")
	newName := new(dto.ChangeNameParams)

	err := c.BodyParser(newName)
	if err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	fmt.Println(oldName, newName.NewName)
	log.Printf("Received request to change account name from %s to %s\n", oldName, newName.NewName)

	err = dto.AllAccounts.ChangeAccountName(newName.NewName, oldName)

	if err != nil {
		log.Printf("Error changing account name: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("Account name changed successfully from %s to %s\n", oldName, newName.NewName)

	return c.SendStatus(fiber.StatusOK)
}

func DeleteAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	err := dto.AllAccounts.DeleteAccount(name)
	if err != nil {
		log.Printf("Error deleting account: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete account"})
	}

	return c.SendStatus(fiber.StatusOK)
}
