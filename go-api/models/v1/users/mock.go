package users

import (
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models"
)

var mockUsers models.Users

func init() {
	usr, _ := models.NewUser("Nick", "Kotenberg", "nick@mail.com", "1234")

	mockUsers = models.Users{
		*usr,
	}
}
