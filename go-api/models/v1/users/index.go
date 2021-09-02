package users

import "github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models"

func Index() (users *models.Users, err error) {
	users = &mockUsers

	return
}
