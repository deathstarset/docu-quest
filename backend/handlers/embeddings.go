package handlers

import (
	"strings"

	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
)

type createEmbedding struct {
	ContentID uuid.UUID `json:"content_id"`
}

func CreateEmbedding(c *fiber.Ctx) error {
	var embeddingInfo createEmbedding
	err := c.BodyParser(&embeddingInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	extractedContent, err := config.DB.FindExtractedContentByID(c.Context(), embeddingInfo.ContentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	extractedContentLines := strings.Split(extractedContent.Content, "\n")

	for _, line := range extractedContentLines {
		embeddings, err := utils.GenerateEmbedding(line)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		_, err = config.DB.AddEmbedding(c.Context(), database.AddEmbeddingParams{ContentID: embeddingInfo.ContentID, Embedding: pgvector.NewVector(embeddings), Text: line})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	}

	return c.SendStatus(fiber.StatusCreated)
}

type ISimilarEmbeddingBody struct {
	Question   string    `json:"question"`
	DocumentID uuid.UUID `json:"document_id"`
}

func GetSimilarEmbeddings(c *fiber.Ctx) error {
	var similarEmbeddingBody ISimilarEmbeddingBody
	err := c.BodyParser(&similarEmbeddingBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	embedding, err := utils.GenerateEmbedding(similarEmbeddingBody.Question)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	similarEmbeddings, err := config.DB.FindSimilarVec(c.Context(), database.FindSimilarVecParams{Embedding: embedding, DocumentID: similarEmbeddingBody.DocumentID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"similar_embeddings": similarEmbeddings})
}
