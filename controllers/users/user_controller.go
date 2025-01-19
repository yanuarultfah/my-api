package users

import (
	"my-api/domain/users"
	"my-api/service"
	"my-api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (string, *errors.RestErr) {
	return userIdParam, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := service.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	userid, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := service.UsersService.GetUser(userid)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	users, err := service.UsersService.SearchByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
