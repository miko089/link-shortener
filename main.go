package main

import (
	"github.com/gofiber/fiber/v2"
	"linkShortener/database"
	"linkShortener/handlers"
	"linkShortener/utils"
)

func main() {
	app := fiber.New()
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	handlers.RegisterHandlers(app, db)

	app.Static("/", "./frontend")

	err = app.Listen(utils.GetPort())
	if err != nil {
		panic(err)
	}
}
