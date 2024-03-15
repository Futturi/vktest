package repository

import (
	"fmt"
	"strings"

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
	query := fmt.Sprintf("SELECT name, genre, data FROM %s", actortable)
	if err := r.db.Select(&actors, query); err != nil {
		return []models.Actor{}, err
	}
	return actors, nil
}

func (r *Actor_Repo) InsertActor(actor models.Actor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, genre, data) VALUES($1, $2, $3) RETURNING id", actortable)
	row := r.db.QueryRow(query, actor.Name, actor.Genre, actor.Data)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Actor_Repo) UpdateActor(id int, actor models.ActorUpdate) error {
	args := make([]interface{}, 0)
	setVal := make([]string, 0)
	argid := 1
	if actor.Genre != nil {
		args = append(args, *actor.Genre)
		setVal = append(setVal, fmt.Sprintf("genre=$%d", argid))
		argid++
	}
	if actor.Name != nil {
		args = append(args, *actor.Name)
		setVal = append(setVal, fmt.Sprintf("name=$%d", argid))
		argid++
	}
	if actor.Data != nil {
		args = append(args, *actor.Data)
		setVal = append(setVal, fmt.Sprintf("data=$%d", argid))
		argid++
	}
	setQuery := strings.Join(setVal, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", actortable, setQuery, argid)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Actor_Repo) FindIdCinemaByName(cin string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROm %s WHERE name = $1", cinematable)
	row := r.db.QueryRow(query, cin)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Actor_Repo) DeleteActor(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", actortable)
	query2 := fmt.Sprintf("DELETE FROM %s WHERE actor_id = $1", authorcinema)
	_, err := r.db.Exec(query2, id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
