package handlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ICreateConversation struct {
	UserID     uuid.UUID `json:"user_id"`
	DocumentID uuid.UUID `json:"document_id"`
}

func CreateConversation(c *fiber.Ctx) error {
	var conversationInfo ICreateConversation

	err := c.BodyParser(&conversationInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	conversation, err := config.DB.AddConversation(c.Context(), database.AddConversationParams(conversationInfo))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"conversation": conversation})
}

func GetConversation(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id uuid.UUID
	err := id.UnmarshalText([]byte(idStr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	conversation, err := config.DB.FindConversationByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"conversation": conversation})
}

func GetConversations(c *fiber.Ctx) error {
	conversations, err := config.DB.FindConversations(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"conversations": conversations})
}

func DeleteConversation(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindConversationByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveConversation(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
