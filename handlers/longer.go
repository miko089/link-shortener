package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"linkShortener/database"
)

func longerLinkHandler(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		short := c.Params("short")
		link, err := database.GetLink(db, short)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Link not found",
			})
		}
		return c.Redirect(link.Full)
	}
}
