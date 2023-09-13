package controllers

import "github.com/gofiber/fiber/v2"

func Accounts(c *fiber.Ctx) error {
	// grab accounts from database
	return c.JSON("working!")
}
