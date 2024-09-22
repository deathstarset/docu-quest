package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenRandomChars() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	idStr := id.String()
	idStr = strings.ReplaceAll(idStr, "-", "")
	idStr = strings.ToLower(idStr)
	return idStr, nil
}

func TextToUUID(text string) (uuid.UUID, error) {
	var uuidText uuid.UUID
	err := uuidText.UnmarshalText([]byte(text))
	if err != nil {
		return uuid.Nil, err
	}
	return uuidText, nil
}
