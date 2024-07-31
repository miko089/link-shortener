package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"linkShortener/database"
	"linkShortener/utils"
)

func shortenLinkHandler(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		payload := new(struct {
			Full string `json:"full"`
		})
		err := c.BodyParser(payload)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}
		if payload.Full == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Full link is required",
			})
		}
		if !utils.IsValidURL(payload.Full) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid URL",
			})
		}
		link, err := database.CreateLink(db, payload.Full)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create link",
			})
		}
		return c.JSON(fiber.Map{
			"short": utils.GetBaseURL() + "/" + link.Short,
		})
	}
}
