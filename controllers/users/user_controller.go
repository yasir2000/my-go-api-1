package users

import (
	"net/http"
	"strconv"

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

func Get(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
