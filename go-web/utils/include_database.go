package utils

import "go-web/models"

func IncludeDatabase(user *models.User) {
	*database = append(*database, *user)
}
