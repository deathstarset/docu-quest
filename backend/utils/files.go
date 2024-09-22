package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GenFileName(file *multipart.FileHeader) (string, error) {
	fileName, err := GenRandomChars()
	if err != nil {
		return "", err
	}
	extension := filepath.Ext(file.Filename)
	fullName := fileName + extension
	return fullName, nil
}

func ParseAndSaveUploadedFile(c *fiber.Ctx) (*multipart.FileHeader, string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, "", err
	}
	newFileName, err := GenFileName(file)
	if err != nil {
		return nil, "", err
	}
	destination := filepath.Join(".", "uploads", newFileName)
	err = c.SaveFile(file, destination)
	if err != nil {
		return nil, "", err
	}
	return file, destination, nil
}
