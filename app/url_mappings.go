package app

import (
	"my-api/controllers/ping"
	"my-api/controllers/thirdparty"
	"my-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/find/:status", users.FindByStatus)
	router.GET("/thirdparty/reqres/get", thirdparty.GetListUser)
}
