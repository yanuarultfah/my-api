package thirdparty

import (
	"my-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListUser(c *gin.Context) {
	listuser, getErr := service.ReqresService.GetListUser()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, listuser)
}
