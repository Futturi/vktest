package repository

import (
	"github.com/Futturi/vktest/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	usertable    = "users"
	cinematable  = "cinema"
	authorcinema = "author_cinema"
	actortable   = "actor"
	admintable   = "admins"
)

type Repository struct {
	Authorization
	ActorRepo
	CinemaRepo
}

func NewRepostitory(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthRepo(db),
		ActorRepo:  NewActorRepo(db),
		CinemaRepo: NewCinemaRepo(db)}
}

type Authorization interface {
	SignUp(User models.User) (int, error)
	SignIn(User models.User) (int, error)
	SignUpAdmin(Admin models.Admin) (int, error)
	SignInAdmin(Admin models.Admin) (int, error)
}

type ActorRepo interface {
	GetActors() ([]models.Actor, error)
	InsertActor(actor models.Actor) (int, error)
	UpdateActor(id int, actor models.ActorUpdate) error
	DeleteActor(id int) error
}

type CinemaRepo interface {
	InsertCinema(cinema models.Cinema) (int, error)
}
