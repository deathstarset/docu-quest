package userhandlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUserExtractedContent(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var extractedContentInfo handlers.ICreateExtractedContent
	err = c.BodyParser(&extractedContentInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	document, err := config.DB.FindUserDocumentByID(c.Context(), database.FindUserDocumentByIDParams{ID: extractedContentInfo.DocumentID, UploadedBy: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	content, err := utils.ParsePdf(document.FilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	extractedContent, err := config.DB.AddExtractedContent(c.Context(), database.AddExtractedContentParams{DocumentID: extractedContentInfo.DocumentID, Content: content})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"extracted_content": extractedContent,
	})
}

func DeleteUserExtractedContent(c *fiber.Ctx) error {
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

	_, err = config.DB.FindUserExtractedContentByID(c.Context(), database.FindUserExtractedContentByIDParams{UploadedBy: userIDUUID, ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveUserExtractedContent(c.Context(), database.RemoveUserExtractedContentParams{UploadedBy: userIDUUID, ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
