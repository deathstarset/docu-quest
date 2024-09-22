package handlers

import (
	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type createExtractedContent struct {
	DocumentID uuid.UUID `json:"document_id"`
	Content    string    `json:"content"`
}

func CreateExtractedContent(c *fiber.Ctx) error {
	var extractedContentInfo createExtractedContent
	err := c.BodyParser(&extractedContentInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	extractedContent, err := config.DB.AddExtractedContent(c.Context(), database.AddExtractedContentParams(extractedContentInfo))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"extracted_content": extractedContent,
	})
}

func GetExtractedContent(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	extractedContent, err := config.DB.FindExtractedContentByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"extracted_content": extractedContent})
}

func GetExtractedContents(c *fiber.Ctx) error {
	extractedContents, err := config.DB.FindExtractedContents(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"extracted_contents": extractedContents})

}
func DeleteExtractedContent(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = config.DB.FindExtractedContentByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveExtractedContent(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
