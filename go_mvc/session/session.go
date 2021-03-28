package session

import (
	"github.com/aaman007/Golang-Web-Dev/go_mvc/models"
)

func GetSession() map[string]models.User {
	return models.LoadUsers()
}
