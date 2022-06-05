package utils

import (
	"encoding/json"
	"go-web/models"
	"os"
)

func ReadDatabase() (*[]models.User, error) {
	if database != nil {
		return database, nil
	}

	fileBytes, err := os.ReadFile("users.json")
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return nil, err
	}

	database = &users

	return database, nil
}
