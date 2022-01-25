package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yasir2000/my-go-api-1/domain/users"
	"github.com/yasir2000/my-go-api-1/services"
	"github.com/yasir2000/my-go-api-1/utils/errors"
)

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	result, SaveErr := services.CreateUser(user)
	if SaveErr != nil {
		c.JSON(SaveErr.Status, SaveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
