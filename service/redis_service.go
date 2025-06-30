package service

import (
	"my-api/domain/redis"
	"my-api/utils/errors"
)

var (
	MovieService movieServiceInterface = &movieService{}
)

type movieService struct {
}

type movieServiceInterface interface {
	CreateMovie(redis.Movie) (*redis.Movie, *errors.RestErr)
	GetMovie(string) (*redis.Movie, *errors.RestErr)
	GetMovies(redis.Movie) ([]redis.Movie, *errors.RestErr)
	UpdateMovie(string, redis.Movie) (*redis.Movie, *errors.RestErr)
	DeleteMovie(string) *errors.RestErr
}

func (s *movieService) CreateMovie(movie redis.Movie) (*redis.Movie, *errors.RestErr) {
	if err := movie.CreateMovie(); err != nil {
		return nil, err
	}
	return &movie, nil
}

func (s *movieService) GetMovie(id string) (*redis.Movie, *errors.RestErr) {
	movie := &redis.Movie{}
	if err := movie.GetMovie(id); err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *movieService) GetMovies(movie redis.Movie) ([]redis.Movie, *errors.RestErr) {
	movies, err := movie.GetMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *movieService) UpdateMovie(id string, movie redis.Movie) (*redis.Movie, *errors.RestErr) {
	if err := movie.UpdateMovie(id); err != nil {
		return nil, err
	}
	return &movie, nil
}

func (s *movieService) DeleteMovie(id string) *errors.RestErr {
	movie := &redis.Movie{}
	if err := movie.DeleteMovie(id); err != nil {
		return err
	}
	return nil
}
