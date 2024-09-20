package handlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type createUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	var userInfo createUser
	err := c.BodyParser(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := utils.HashPassword(userInfo.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := config.DB.AddUser(c.Context(), database.AddUserParams{Username: userInfo.Username, Email: userInfo.Email, Password: hashedPassword})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"user": user})
}

func GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id uuid.UUID
	err := id.UnmarshalText([]byte(idStr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := config.DB.FindUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}
func GetUsers(c *fiber.Ctx) error {

	user, err := config.DB.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"users": user})
}

func DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id uuid.UUID
	err := id.UnmarshalText([]byte(idStr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = config.DB.RemoveUser(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

type updateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// solve the update problem
func UpdateUser(c *fiber.Ctx) error {
	var userInfo updateUser
	err := c.BodyParser(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	idStr := c.Params("id")
	var id uuid.UUID
	err = id.UnmarshalText([]byte(idStr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := utils.HashPassword(userInfo.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := config.DB.EditUser(c.Context(), database.EditUserParams{Username: userInfo.Username, Email: userInfo.Email, Password: hashedPassword, ID: id})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}
