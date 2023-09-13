package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
	"github.com/omarattia3143/paytabs-backend-challenge/src/services"
)

func Transfer(c *fiber.Ctx) error {
	var transferRequest models.TransferRequest

	if err := c.BodyParser(&transferRequest); err != nil {
		return err
	}

	err := services.StartTransferProcess(transferRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("Success")
}
