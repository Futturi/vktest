package service

import (
	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
)

type Actor_Service struct {
	repo repository.ActorRepo
}

func NewActorService(repo repository.ActorRepo) *Actor_Service {
	return &Actor_Service{repo: repo}
}

func (a *Actor_Service) GetActors() ([]models.Actor, error) {
	return a.repo.GetActors()
}
func (a *Actor_Service) InsertActor(actor models.Actor) (int, error) {
	return a.repo.InsertActor(actor)
}
