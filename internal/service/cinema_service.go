package service

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"time"

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
	if cinema.Data != "" {
		data, err := time.Parse("2006-Jan-02", cinema.Data)
		if err != nil {
			return 0, err
		}
		cinema.Data = fmt.Sprintf("%d", data.Unix())
	}
	if cinema.Data == "" {
		cinema.Data = "0"
	}
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

func (a *Cinema_Service) UpdateFilm(id string, cinema models.CinemaUpdate) error {
	if cinema.Data != nil {
		data, err := time.Parse("2006-Jan-02", *cinema.Data)
		if err != nil {
			return err
		}
		t := fmt.Sprintf("%d", data.Unix())
		cinema.Data = &t
	}
	if cinema.Data == nil {
		a := "0"
		cinema.Data = &a
	}
	if len(*cinema.Name) > 150 {
		return errors.New("your name is longer than 150")
	}
	if len(*cinema.Description) > 1000 {
		return errors.New("your description is longer than 1000")
	}
	if *cinema.Rating > 10 || *cinema.Rating < 0 {
		return errors.New("incorrect rating")
	}
	return a.repo.UpdateFilm(id, cinema)
}

func (a *Cinema_Service) DeleteFilm(id string) error {
	return a.repo.DeleteFilm(id)
}

func (a *Cinema_Service) GetCinemas(sor string) ([]models.Cinema, error) {
	var result []models.Cinema
	cinemas, err := a.repo.GetCinemas(sor)
	if err != nil {
		return []models.Cinema{}, err
	}
	if sor == "rating" {
		slices.SortFunc(cinemas, func(a, b models.Cinema) int {
			switch a.Rating > b.Rating {
			case true:
				return -1
			case false:
				return 1
			default:
				return 0
			}
		})
	}
	if sor == "name" {
		slices.SortFunc(cinemas, func(a, b models.Cinema) int {
			switch a.Name > b.Name {
			case true:
				return -1
			case false:
				return 1
			default:
				return 0
			}
		})
	}
	if sor == "date" {
		slices.SortFunc(cinemas, func(a, b models.Cinema) int {
			data1, err := strconv.Atoi(a.Data)
			if err != nil {
				return 0
			}
			data2, err := strconv.Atoi(b.Data)
			if err != nil {
				return 0
			}
			switch data1 > data2 {
			case true:
				return -1
			case false:
				return 1
			default:
				return 0
			}
		})
	}
	for _, cin := range cinemas {
		var b string
		i, err := strconv.ParseInt(cin.Data, 10, 64)
		if err != nil {
			return []models.Cinema{}, err
		}
		tm := time.Unix(i, 0)
		b = tm.Format("2006-Jan-02")
		result = append(result, models.Cinema{
			Id:          cin.Id,
			Name:        cin.Name,
			Description: cin.Description,
			Data:        b,
			Rating:      cin.Rating,
			Actors:      cin.Actors,
		})
	}
	return result, nil
}

func (a *Cinema_Service) Search(search models.Search) ([]models.Cinema, error) {
	hash := make(map[string]int, 0)
	cinemas1, cinemas2, err := a.repo.Search(search)
	if err != nil {
		return []models.Cinema{}, err
	}
	for _, val := range cinemas1 {
		hash[val.Name]++
	}
	for _, val := range cinemas2 {
		hash[val.Name]++
	}
	return a.repo.Unification(hash)
}
