package handlers

import (
	"os"
	"time"

	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Register(c *fiber.Ctx) error {
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

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var loginInfo login
	err := c.BodyParser(&loginInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := config.DB.FindUserByEmail(c.Context(), loginInfo.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	isMatch := utils.ComparePasswordWithHash(loginInfo.Password, user.Password)
	if !isMatch {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Credentials incorrect"})
	}

	day := time.Hour * 24
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenStr})
}

func Logout(c *fiber.Ctx) error {
	return nil
}
