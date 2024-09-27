package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var ollamaUrl = "http://localhost:11434"

type IEmbeddingsRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type IEmbeddingResponse struct {
	Model           string      `json:"model"`
	Embeddings      [][]float32 `json:"embeddings"`
	TotalDuration   int64       `json:"total_duration"`
	LoadDuration    int64       `json:"load_duration"`
	PromptEvalCount int         `json:"prompt_eval_count"`
}

func GenerateEmbedding(content string) ([]float32, error) {
	url := ollamaUrl + "/api/embed"
	data := IEmbeddingsRequest{
		Model: "nomic-embed-text:latest",
		Input: content,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var embeddingResponse IEmbeddingResponse
	err = json.Unmarshal(body, &embeddingResponse)
	if err != nil {
		return nil, err
	}

	return embeddingResponse.Embeddings[0], nil
}
