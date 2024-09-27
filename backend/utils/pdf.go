package utils

import (
	"bytes"
	"context"
	"os/exec"
	"regexp"
	"strings"
)

func ParsePdf(path string) (string, error) {
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
