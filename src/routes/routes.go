package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarattia3143/paytabs-backend-challenge/src/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("api/v1")

	api.Get("accounts", controllers.Accounts)
	api.Get("accounts/:id", controllers.Account)

	api.Post("transfer", controllers.Transfer)
}
