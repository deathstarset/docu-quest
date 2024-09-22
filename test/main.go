package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	pdfPath := "test.pdf"
	content, err := parsePdf(pdfPath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = genEmbeddings(content)
	if err != nil {
		log.Fatalf(err.Error())
	}

	/* f, err := os.Create("result4.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for _, prLine := range processedLines {
		_, err := f.Write([]byte(prLine + "\n"))
		if err != nil {
			log.Fatalln(err)
		}
	} */
}

func parsePdf(path string) (string, error) {
	args := []string{"-layout", "-nopgbrk", path, "-"}
	cmd := exec.CommandContext(context.Background(), "pdftotext", args...)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	outputStr := output.String()
	lines := strings.Split(outputStr, "\n")
	var processedLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		spaceRe := regexp.MustCompile(`\s+`)
		line = spaceRe.ReplaceAllString(line, " ")

		punctuation := "!\"#$%&'()*+,./:;<=>?@[\\]^_`{|}~"
		nonRe := regexp.MustCompile(`[^a-zA-Z0-9\s` + punctuation + `]`)
		line = nonRe.ReplaceAllString(line, "")

		if line != "" {
			processedLines = append(processedLines, line)
		}
	}

	return strings.Join(processedLines, "\n"), nil
}

type IEmbeddingsRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

func genEmbeddings(content string) error {
	contentSplit := strings.Split(content, "\n")
	url := "http://localhost:11434/api/embed"
	data := IEmbeddingsRequest{
		Model: "nomic-embed-text:latest",
		Input: contentSplit,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
