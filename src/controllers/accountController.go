package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
	"github.com/omarattia3143/paytabs-backend-challenge/src/services"
)

func Accounts(c *fiber.Ctx) error {
	// create a channel to receive the result
	resultChan := make(chan []models.Account)

	// use a goroutine to fetch accounts from the database
	go func() {
		// grab accounts from database
		accounts := services.GetAllAccounts()
		// send the result to the channel
		resultChan <- *accounts
	}()

	accounts := <-resultChan

	// serialize accounts into json and return it to client
	return c.JSON(accounts)
}

func Account(c *fiber.Ctx) error {
	id := c.Params("id")
	//validate uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad params")
	}

	accountChan := make(chan *models.Account)
	errChan := make(chan error)

	go func() {
		account := services.GetAccount(id)
		accountChan <- account
	}()

	select {
	case account := <-accountChan:
		if account == nil {
			// if the account is not found, return a 404 status
			return c.Status(fiber.StatusNotFound).SendString("Resource not found")
		}
		return c.JSON(account)
	case err := <-errChan:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
