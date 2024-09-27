package middlewares

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
)

func OnlyAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDStr := utils.GetUserID(c)
		userId, err := utils.TextToUUID(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Invalid user"})
		}
		user, err := config.DB.FindUserByID(c.Context(), userId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
		}
		if user.Role != database.UserTypeAdmin {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Access forbidden"})
		}
		return c.Next()
	}
}
