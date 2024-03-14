package repository

import (
	"fmt"

	"github.com/Futturi/vktest/internal/models"
	"github.com/jmoiron/sqlx"
)

type Cinema_Repo struct {
	db *sqlx.DB
}

func NewCinemaRepo(db *sqlx.DB) *Cinema_Repo {
	return &Cinema_Repo{db: db}
}

func (r *Cinema_Repo) InsertCinema(cinema models.Cinema) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, description, data, rating) VALUES($1,$2,$3,$3) RETURNING ID", cinematable)
	row := r.db.QueryRow(query, cinema.Name, cinema.Description, cinema.Data, cinema.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	for _, acr := range cinema.Actors {
		idacr, err := r.GetIdActorByName(acr)
		if err != nil {
			return 0, err
		}
		query2 := fmt.Sprintf("INSERT INTO %s(actor_id, cinema_id) VALUES($1, $2)", authorcinema)
		_, err = r.db.Exec(query2, idacr, id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (r *Cinema_Repo) GetIdActorByName(name string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROm %s WHERE name = $1", actortable)
	row := r.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
