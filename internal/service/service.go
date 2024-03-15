package service

import (
	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
)

type Service struct {
	AuthService
	ActorService
	CinemaService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{AuthService: NewAuthService(repo.Authorization),
		ActorService:  NewActorService(repo.ActorRepo),
		CinemaService: NewCinemaService(repo.CinemaRepo)}
}

type AuthService interface {
	SignUp(User models.User) (int, error)
	SignIn(User models.User) (string, error)
	ParseToken(token string) (string, bool, error)
	SignUpAdmin(Admin models.Admin) (int, error)
	SignInAdmin(Admin models.Admin) (string, error)
}

type ActorService interface {
	GetActors() ([]models.Actor, error)
	InsertActor(actor models.Actor) (int, error)
	UpdateActor(id int, actor models.ActorUpdate) error
	DeleteActor(id string) error
}

type CinemaService interface {
	InsertCinema(cinema models.Cinema) (int, error)
	UpdateFilm(id string, cinema models.CinemaUpdate) error
	DeleteFilm(id string) error
	GetCinemas(sor string) ([]models.Cinema, error)
	Search(search models.Search) ([]models.Cinema, error)
}
