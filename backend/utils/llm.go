package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ILLMRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ILLMResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
}

func GenLLMResponse(context string, question string) (string, error) {
	promptTemplate := fmt.Sprintf(`
	You are an LLM specifically designed to process PDF documents and answer questions by retrieving and generating information. The following text was extracted from a PDF using a similarity search:
  %s

	Please analyze this text and answer the following question:
  %s

  Make sure to explain the reasoning behind the answer and refer to the key points in the text when necessary.
	`, context, question)

	reqBody := ILLMRequest{
		Model:  "gemma:2b",
		Prompt: promptTemplate,
		Stream: false,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	url := ollamaUrl + "/api/generate"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var llmResponse ILLMResponse
	err = json.Unmarshal(body, &llmResponse)
	if err != nil {
		return "", err
	}

	return llmResponse.Response, nil
}
