package service

import (
	"errors"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
)

type Cinema_Service struct {
	repo repository.CinemaRepo
}

func NewCinemaService(repo repository.CinemaRepo) *Cinema_Service {
	return &Cinema_Service{repo: repo}
}

func (a *Cinema_Service) InsertCinema(cinema models.Cinema) (int, error) {
	if len(cinema.Name) > 150 {
		return 0, errors.New("your name is longer than 150")
	}
	if len(cinema.Description) > 1000 {
		return 0, errors.New("your description is longer than 1000")
	}
	if cinema.Rating > 10 || cinema.Rating < 0 {
		return 0, errors.New("incorrect rating")
	}
	return a.repo.InsertCinema(cinema)
}
