package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
)

type Actor_Service struct {
	repo repository.ActorRepo
}

func NewActorService(repo repository.ActorRepo) *Actor_Service {
	return &Actor_Service{repo: repo}
}

func (a *Actor_Service) GetActors() ([]models.ActorSelect, error) {
	var result []models.ActorSelect
	actors, err := a.repo.GetActors()
	if err != nil {
		return []models.ActorSelect{}, err
	}
	for _, acr := range actors {
		var b string
		i, err := strconv.ParseInt(acr.Data, 10, 64)
		if err != nil {
			return []models.ActorSelect{}, err
		}
		tm := time.Unix(i, 0)
		b = tm.Format("2006-Jan-02")
		result = append(result, models.ActorSelect{Id: acr.Id, Name: acr.Name, Genre: acr.Genre, Data: b, Cinemas: acr.Cinemas})
	}
	return result, nil
}
func (a *Actor_Service) InsertActor(actor models.Actor) (int, error) {
	data, err := time.Parse("2006-Jan-02", actor.Data)
	if err != nil {
		return 0, err
	}
	actor.Data = fmt.Sprintf("%d", data.Unix())
	return a.repo.InsertActor(actor)
}

func (a *Actor_Service) UpdateActor(id int, actor models.ActorUpdate) error {
	if actor.Data != nil {
		data, err := time.Parse("2006-Jan-02", *actor.Data)
		if err != nil {
			return err
		}
		t := fmt.Sprintf("%d", data.Unix())
		actor.Data = &t
	}
	return a.repo.UpdateActor(id, actor)
}

func (a *Actor_Service) DeleteActor(id string) error {
	return a.repo.DeleteActor(id)
}
