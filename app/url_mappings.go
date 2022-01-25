package app

import (
	"github.com/yasir2000/my-go-api-1/controllers/users"
)

func mapUrls() {
	router.POST("/users", users.Create)
}
