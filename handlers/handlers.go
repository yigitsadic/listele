package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigitsadic/listele/database"
)

// HandleList handles listing data taken from repository.
func HandleList(repo database.Repository) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		people, err := repo.FindAll()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.JSON(people)
	}
}
