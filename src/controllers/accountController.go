package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/omarattia3143/paytabs-backend-challenge/src/services"
)

func Accounts(c *fiber.Ctx) error {
	// grab accounts from database
	accounts := services.GetAllAccounts()
	// serialize accounts into json and return it to client
	return c.JSON(accounts)
}

func Account(c *fiber.Ctx) error {
	id := c.Params("id")
	//validate uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	accounts := services.GetAccount(id)
	// serialize accounts into json and return it to client
	return c.JSON(accounts)
}
