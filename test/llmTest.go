package main

/* data := Request{
	Model:  "gemma:2b",
	Prompt: "Hello there can you tell me how do i create a loop in python",
}

jsonByte, _ := json.Marshal(data)

res, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonByte))
if err != nil {
	log.Fatal(err)
}

defer res.Body.Close()

decoder := json.NewDecoder(res.Body)
var result []LLMResponse

for {
	var chunk map[string]interface{}

	err := decoder.Decode(&chunk)
	if err != nil {
		if err == io.EOF {
			break
		}
		fmt.Printf("Error decoding chunk %s", err.Error())
		return
	}

	llmRes := LLMResponse{
		Model:     chunk["model"].(string),
		CreatedAt: chunk["created_at"].(string),
		Response:  chunk["response"].(string),
	}

	result = append(result, llmRes)

	fmt.Printf("Processed chunk: %+v\n", llmRes)
}

fmt.Printf("Total processed items: %d\n", len(result)) */

/* type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}
type LLMResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
} */
