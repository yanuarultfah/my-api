package redis

import (
	"context"
	"encoding/json"
	rds "my-api/datasource/redis/user_redis"
	"my-api/utils/errors"

	"github.com/google/uuid"
)

func (movie *Movie) CreateMovie() *errors.RestErr {
	c := rds.NewRedisCache()
	var ctx = context.Background()
	movie.ID = uuid.New().String()
	json, err := json.Marshal(movie)
	if err != nil {
		return errors.NewInternalServerError("error marshalling movie data")
	}
	c.HSet(ctx, "movies", movie.ID, json)
	if err != nil {
		return errors.NewInternalServerError("error saving movie data to redis")
	}
	return nil
}

func (movie *Movie) GetMovie(id string) *errors.RestErr {
	c := rds.NewRedisCache()
	ctx := context.Background()
	val, err := c.HGet(ctx, "movies", id).Result()
	if err != nil {
		return errors.NewInternalServerError("error retrieving movie data from redis")
	}
	err = json.Unmarshal([]byte(val), movie)
	if err != nil {
		return errors.NewInternalServerError("error unmarshalling movie data")
	}
	return nil
}

func (movie *Movie) GetMovies() ([]Movie, *errors.RestErr) {
	c := rds.NewRedisCache()
	ctx := context.Background()
	movies := []Movie{}
	val, err := c.HGetAll(ctx, "movies").Result()
	if err != nil {
		return nil, errors.NewInternalServerError("error retrieving movies from redis")
	}
	for _, v := range val {
		movie := &Movie{}
		err = json.Unmarshal([]byte(v), movie)
		if err != nil {
			return nil, errors.NewInternalServerError("error unmarshalling movie data")
		}
		movies = append(movies, *movie)
	}
	return movies, nil
}

func (movie *Movie) UpdateMovie(id string) *errors.RestErr {
	c := rds.NewRedisCache()
	ctx := context.Background()
	movie.ID = id
	json, err := json.Marshal(movie)
	if err != nil {
		return errors.NewInternalServerError("error marshalling movie data")
	}
	_, err = c.HSet(ctx, "movies", id, json).Result()
	if err != nil {
		return errors.NewInternalServerError("error updating movie in redis")
	}
	return nil
}

func (movie *Movie) DeleteMovie(id string) *errors.RestErr {
	c := rds.NewRedisCache()
	ctx := context.Background()
	_, err := c.HDel(ctx, "movies", id).Result()
	if err != nil {
		return errors.NewInternalServerError("error deleting movie from redis")
	}
	return nil
}
