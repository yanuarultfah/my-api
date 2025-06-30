package app

import (
	"my-api/controllers/ping"
	"my-api/controllers/redis"
	"my-api/controllers/thirdparty"
	"my-api/controllers/users"
	"my-api/controllers/workerpool"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/find/:status", users.FindByStatus)
	router.GET("/thirdparty/reqres/get", thirdparty.GetListUser)
	router.POST("/worker/savecsv", workerpool.Simpan)
	router.POST("/movie/create", redis.CreateMovie)
	router.GET("/movie/:id", redis.GetMovie)
	router.GET("/movies", redis.GetMovies)
	router.POST("/movie/update/:id", redis.UpdateMovie)
	router.POST("/movie/del/:id", redis.DeleteMovie)
}
