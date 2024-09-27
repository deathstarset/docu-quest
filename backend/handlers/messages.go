package handlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ICreateMessage struct {
	ConversationID uuid.UUID           `json:"conversation_id"`
	Content        string              `json:"content"`
	Sender         database.SenderType `json:"sender"`
}

func CreateMessage(c *fiber.Ctx) error {
	var messageInfo ICreateMessage

	err := c.BodyParser(&messageInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	message, err := config.DB.AddMessage(c.Context(), database.AddMessageParams(messageInfo))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": message})
}

func GetMessage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id uuid.UUID
	err := id.UnmarshalText([]byte(idStr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	message, err := config.DB.FindMessageByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func GetMessages(c *fiber.Ctx) error {
	messages, err := config.DB.FindMessages(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"messages": messages})
}

func DeleteMessage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id uuid.UUID
	err := id.UnmarshalText([]byte(idStr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindMessageByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveMessage(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
