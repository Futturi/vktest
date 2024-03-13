package repository

import (
	"fmt"

	"github.com/Futturi/vktest/internal/models"
	"github.com/jmoiron/sqlx"
)

type Actor_Repo struct {
	db *sqlx.DB
}

func NewActorRepo(db *sqlx.DB) *Actor_Repo {
	return &Actor_Repo{db: db}
}

func (r *Actor_Repo) GetActors() ([]models.Actor, error) {
	var actors []models.Actor
	query := fmt.Sprintf("SELECT * FROM %s", actortable)
	if err := r.db.Select(&actors, query); err != nil {
		return []models.Actor{}, err
	}

	for _, act := range actors {
		var cinemas []string
		query2 := fmt.Sprintf("SELECT name FROM %s c INNER JOIN %s ca ON c.id = ca.cinema_id WHERE ca.actor_id = $1", cinematable, authorcinema)
		if err := r.db.Select(&cinemas, query2); err != nil {
			act.Cinemas = []string{}
		}
		act.Cinemas = cinemas
	}
	return actors, nil
}

func (r *Actor_Repo) InsertActor(actor models.Actor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, genre, data) VALUES($1, $2, $3) RETURNING id", actortable)
	row := r.db.QueryRow(query, &id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
