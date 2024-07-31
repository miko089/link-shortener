package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterHandlers(app *fiber.App, db *gorm.DB) {
	app.Post("/shorten", shortenLinkHandler(db))
	app.Get("/:short", longerLinkHandler(db))
}
