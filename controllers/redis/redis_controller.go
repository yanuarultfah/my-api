package redis

import (
	"my-api/domain/redis"
	"my-api/service"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var movie redis.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// if err := movie.CreateMovie(); err != nil {
	// 	log.Printf("Error creating movie: %v", err)
	// 	c.JSON(500, gin.H{"error": "Internal server error"})
	// 	return
	// }
	result, saveData := service.MovieService.CreateMovie(movie)
	if saveData != nil {
		c.JSON(saveData.Status, saveData)
	}
	c.JSON(201, gin.H{"message": "Movie created successfully", "Data": result})
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Movie ID is required"})
	}
	movie, err := service.MovieService.GetMovie(id)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(200, gin.H{"message": "Movie retrived successfully", "Data": movie})
}

func GetMovies(c *gin.Context) {
	movie := &redis.Movie{}
	movies, err := service.MovieService.GetMovies(*movie)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(200, gin.H{"message": "Movies retrived successfully", "Data": movies})
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Movie ID is required"})
	}
	var movie redis.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result, err := service.MovieService.UpdateMovie(id, movie)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(200, gin.H{"message": "Movie updated successfully", "Data": result})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": " Movie ID is required"})
		return
	}
	if err := service.MovieService.DeleteMovie(id); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(200, gin.H{"message": "Movie deleted successfully"})

}
