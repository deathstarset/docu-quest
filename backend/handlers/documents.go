package handlers

import (
	"os"

	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ICreateDocument struct {
	FilePath   string    `json:"file_path"`
	UploadedBy uuid.UUID `json:"uploaded_by"`
}

func CreateDocument(c *fiber.Ctx) error {
	uploadedByStr := c.FormValue("uploaded_by")
	uploadedByUUID, err := utils.TextToUUID(uploadedByStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	_, dest, err := utils.ParseAndSaveUploadedFile(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	documentInfo := ICreateDocument{
		FilePath:   dest,
		UploadedBy: uploadedByUUID,
	}

	document, err := config.DB.AddDocument(c.Context(), database.AddDocumentParams(documentInfo))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"document": document,
	})
}

func GetDocument(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	document, err := config.DB.FindDocumentByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"document": document,
	})
}

func GetDocuments(c *fiber.Ctx) error {
	documents, err := config.DB.FindDocuments(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"documents": documents,
	})
}

func DeleteDocument(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := utils.TextToUUID(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	document, err := config.DB.FindDocumentByID(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveDocument(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = os.Remove(document.FilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
