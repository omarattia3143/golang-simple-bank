package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/omarattia3143/paytabs-backend-challenge/src/database"
	"github.com/omarattia3143/paytabs-backend-challenge/src/routes"
	"log"
)

func main() {
	// setup in memory db and seed database using the provided mocked accounts json file
	database.SetupAndConnectMemDb()
	database.SeedDb()

	// I am using Fiber v2 ----> api base url: http://localhost:8000/api/v1
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
