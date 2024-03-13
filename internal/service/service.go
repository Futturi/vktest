package service

import (
	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
)

type Service struct {
	AuthService
	ActorService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{AuthService: NewAuthService(repo.Authorization),
		ActorService: NewActorService(repo.ActorRepo)}
}

type AuthService interface {
	SignUp(User models.User) (int, error)
	SignIn(User models.User) (string, error)
	ParseToken(token string) (string, string, error)
	SignUpAdmin(Admin models.Admin) (int, error)
	SignInAdmin(Admin models.Admin) (string, error)
}

type ActorService interface {
	GetActors() ([]models.Actor, error)
	InsertActor(actor models.Actor) (int, error)
}
