package handlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := config.DB.FindUserByID(c.Context(), userIDUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}

func DeleteProfile(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindUserByID(c.Context(), userIDUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveUser(c.Context(), userIDUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindUserByID(c.Context(), userIDUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var userInfo updateUser
	err = c.BodyParser(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := utils.HashPassword(userInfo.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := config.DB.EditUser(c.Context(), database.EditUserParams{Username: userInfo.Username, Email: userInfo.Email, Password: hashedPassword})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}
