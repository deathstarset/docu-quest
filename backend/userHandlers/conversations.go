package userhandlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/transactions"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ICreateUserConversation struct {
	DocuemntID uuid.UUID `json:"document_id"`
}

func CreateUserConversation(c *fiber.Ctx) error {
	var conversationInfo ICreateUserConversation
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = c.BodyParser(&conversationInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	conversation, err := config.DB.AddConversation(c.Context(), database.AddConversationParams{UserID: userIDUUID, DocumentID: conversationInfo.DocuemntID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"conversation": conversation})
}

type ISendMessageReq struct {
	Message string `json:"message"`
}

func SendMessage(c *fiber.Ctx) error {
	var sendMessageBody ISendMessageReq
	conversationID := c.Params("id")
	conversationIDUUID, err := utils.TextToUUID(conversationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = c.BodyParser(&sendMessageBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	response, err := transactions.SendMessageTx(c.Context(), handlers.ICreateMessage{ConversationID: conversationIDUUID, Content: sendMessageBody.Message, Sender: database.SenderTypeUser})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": response})
}

func DeleteUserConversation(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindUserConversation(c.Context(), database.FindUserConversationParams{ID: id, UserID: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveUserConversation(c.Context(), database.RemoveUserConversationParams{ID: id, UserID: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserConversation(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	conversation, err := config.DB.FindUserConversation(c.Context(), database.FindUserConversationParams{ID: id, UserID: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"conversation": conversation})
}
