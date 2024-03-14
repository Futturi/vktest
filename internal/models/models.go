package models

import "time"

type Actor struct {
	Id      int       `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Genre   string    `json:"genre" db:"genre"`
	Data    time.Time `json:"data" db:"data"`
	Cinemas []string
}

type ActorUpdate struct {
	Id      *int       `json:"id" db:"id"`
	Name    *string    `json:"name" db:"name"`
	Genre   *string    `json:"genre" db:"genre"`
	Data    *time.Time `json:"data" db:"data"`
	Cinemas *[]string
}

type Cinema struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Data        time.Time `json:"data" db:"data"`
	Rating      int       `json:"rating" db:"rating"`
	Actors      []string
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Admin struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Token struct {
	Access string `json:"access_token"`
}
