package userhandlers

import (
	"os"

	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUserDocument(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	userIDUUID, err := utils.TextToUUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, dest, err := utils.ParseAndSaveUploadedFile(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	documentInfo := handlers.ICreateDocument{
		FilePath:   dest,
		UploadedBy: userIDUUID,
	}

	document, err := config.DB.AddDocument(c.Context(), database.AddDocumentParams(documentInfo))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"document": document,
	})
}

func DeleteUserDocument(c *fiber.Ctx) error {
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

	document, err := config.DB.FindUserDocumentByID(c.Context(), database.FindUserDocumentByIDParams{ID: id, UploadedBy: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = os.Remove(document.FilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = config.DB.RemoveUserDocument(c.Context(), database.RemoveUserDocumentParams{ID: id, UploadedBy: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func GetUserDocument(c *fiber.Ctx) error {
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

	document, err := config.DB.FindUserDocumentByID(c.Context(), database.FindUserDocumentByIDParams{ID: id, UploadedBy: userIDUUID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"document": document})
}
