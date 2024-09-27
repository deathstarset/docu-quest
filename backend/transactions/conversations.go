package transactions

import (
	"context"
	"strings"

	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/database"
	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/pgvector/pgvector-go"
)

func SendMessageTx(ctx context.Context, messageInfo handlers.ICreateMessage) (string, error) {
	// start the transaction
	tx, err := config.Client.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	qtx := config.DB.WithTx(tx)

	// get the conversation
	conversation, err := qtx.FindConversationByID(ctx, messageInfo.ConversationID)
	if err != nil {
		return "", err
	}

	// creating the message
	userMessage, err := qtx.AddMessage(ctx, database.AddMessageParams{ConversationID: messageInfo.ConversationID, Content: messageInfo.Content, Sender: messageInfo.Sender})
	if err != nil {
		return "", err
	}
	// embedding the message
	messageEmbedding, err := utils.GenerateEmbedding(userMessage.Content)
	if err != nil {
		return "", err
	}
	// find similar embeddings to the message
	similarEmbeddings, err := qtx.FindSimilarVec(ctx, database.FindSimilarVecParams{Embedding: pgvector.NewVector(messageEmbedding), DocumentID: conversation.DocumentID})
	if err != nil {
		return "", err
	}

	var similarText []string
	for _, embedding := range similarEmbeddings {
		similarText = append(similarText, embedding.Text)
	}

	// get llm response
	response, err := utils.GenLLMResponse(strings.Join(similarText, "\n"), messageInfo.Content)
	if err != nil {
		return "", err
	}

	// create llm message
	LlmMessage, err := qtx.AddMessage(ctx, database.AddMessageParams{ConversationID: messageInfo.ConversationID, Content: response, Sender: database.SenderTypeBot})
	if err != nil {
		return "", err
	}

	return LlmMessage.Content, tx.Commit()
}
