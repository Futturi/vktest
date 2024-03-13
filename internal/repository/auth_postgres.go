package repository

import (
	"fmt"

	"github.com/Futturi/vktest/internal/models"
	"github.com/jmoiron/sqlx"
)

type Auth_Repo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *Auth_Repo {
	return &Auth_Repo{db: db}
}

func (r *Auth_Repo) SignUp(User models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, password) VALUES($1, $2) RETURNING ID", usertable)
	row := r.db.QueryRow(query, User.Username, User.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth_Repo) SignIn(User models.User) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2", usertable)
	row := r.db.QueryRow(query, User.Username, User.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth_Repo) SignUpAdmin(Admin models.Admin) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, password) VALUES ($1, $2) RETURNING id", admintable)
	row := r.db.QueryRow(query, Admin.Username, Admin.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth_Repo) SignInAdmin(Admin models.Admin) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2", admintable)
	row := r.db.QueryRow(query, Admin.Username, Admin.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
