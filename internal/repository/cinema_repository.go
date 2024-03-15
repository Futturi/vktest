package repository

import (
	"fmt"
	"strings"

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
	query := fmt.Sprintf("INSERT INTO %s(name, description, data, rating) VALUES($1,$2,$3,$4) RETURNING ID", cinematable)
	row := r.db.QueryRow(query, cinema.Name, cinema.Description, cinema.Data, cinema.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	for _, acr := range cinema.Actors {
		fmt.Println(acr)
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
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", actortable)
	row := r.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Cinema_Repo) UpdateFilm(id string, cinema models.CinemaUpdate) error {
	b := "0"
	args := make([]interface{}, 0)
	setVal := make([]string, 0)
	argid := 1
	if cinema.Data != &b {
		args = append(args, *cinema.Data)
		setVal = append(setVal, fmt.Sprintf("data = $%d", argid))
		argid++
	}
	if cinema.Name != nil {
		args = append(args, *cinema.Name)
		setVal = append(setVal, fmt.Sprintf("name=$%d", argid))
		argid++
	}
	if cinema.Description != nil {
		args = append(args, *cinema.Description)
		setVal = append(setVal, fmt.Sprintf("description=$%d", argid))
		argid++
	}
	if cinema.Rating != nil {
		args = append(args, *cinema.Rating)
		setVal = append(setVal, fmt.Sprintf("rating=$%d", argid))
		argid++
	}
	setQuery := strings.Join(setVal, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", cinematable, setQuery, argid)
	args = append(args, argid)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	if cinema.Actors != nil {
		for _, act := range *cinema.Actors {
			idact, err := r.GetIdActorByName(act)
			if err != nil {
				return err
			}
			query2 := fmt.Sprintf("INSERT INTO %s(actor_id, cinema_id) VALUES($1, $2)", authorcinema)
			_, err = r.db.Exec(query2, idact, id)
			if err != nil {
				return err
			}
		}
	}
	return nil

}

func (r *Cinema_Repo) DeleteFilm(id string) error {
	query2 := fmt.Sprintf("DELETE FROM %s WHERE cinema_id = $1", authorcinema)
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", cinematable)
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

func (r *Cinema_Repo) GetCinemas(sor string) ([]models.Cinema, error) {
	var result []models.Cinema
	var cinemas []models.Cinema
	query := fmt.Sprintf("SELECT id, name, description, data, rating FROM %s", cinematable)
	if err := r.db.Select(&cinemas, query); err != nil {
		return []models.Cinema{}, err
	}
	for _, val := range cinemas {
		var rs []string
		id := val.Id
		query := fmt.Sprintf("SELECT name FROM %s a INNER JOIN %s ca ON a.id = ca.actor_id WHERE ca.cinema_id = $1", actortable, authorcinema)
		if err := r.db.Select(&rs, query, id); err != nil {
			return []models.Cinema{}, nil
		}
		result = append(result, models.Cinema{
			Id:          val.Id,
			Name:        val.Name,
			Description: val.Description,
			Data:        val.Data,
			Rating:      val.Rating,
			Actors:      rs,
		})
	}
	return result, nil
}

func (r *Cinema_Repo) Search(search models.Search) ([]models.Cinema, []models.Cinema, error) {
	var cinemas1 []models.Cinema
	var cinemas2 []models.Cinema
	str := "%" + search.Search + "%"
	query1 := fmt.Sprintf("SELECT id, name, description, data, rating FROM %s WHERE name LIKE('%s')", cinematable, str)
	query2 := fmt.Sprintf(`select c.id, c.name, c.description, c.data, c.rating from %s c 
	inner join %s ac on c.id = ac.cinema_id 
	inner join %s a on a.id = ac.actor_id where a.name LIKE('%s')`, cinematable, authorcinema, actortable, str)
	if err := r.db.Select(&cinemas1, query1); err != nil {
		return []models.Cinema{}, []models.Cinema{}, err
	}
	if err := r.db.Select(&cinemas2, query2); err != nil {
		return []models.Cinema{}, []models.Cinema{}, err
	}
	return cinemas1, cinemas2, nil

}

func (r *Cinema_Repo) Unification(hash map[string]int) ([]models.Cinema, error) {
	result := make([]models.Cinema, 0)
	var result2 []models.Cinema
	for key := range hash {
		var mod []models.Cinema
		query := fmt.Sprintf("SELECT id, name, description, data, rating FROM %s WHERE name = $1", cinematable)
		if err := r.db.Select(&mod, query, key); err != nil {
			return []models.Cinema{}, err
		}
		result = append(result, mod...)
	}
	for _, val := range result {
		var names []string
		query := fmt.Sprintf("SELECT name FROM %s a INNER JOIN %s ca ON ca.actor_id = a.id WHERE ca.cinema_id = $1", actortable, authorcinema)
		if err := r.db.Select(&names, query, val.Id); err != nil {
			return []models.Cinema{}, err
		}
		result2 = append(result2, models.Cinema{
			Id:          val.Id,
			Name:        val.Name,
			Description: val.Description,
			Data:        val.Data,
			Rating:      val.Rating,
			Actors:      names,
		})
	}

	return result2, nil
}
