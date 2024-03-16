package models

type Actor struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Genre string `json:"genre" db:"genre"`
	Data  string `json:"data" db:"data"`
}

type ActorUpdate struct {
	Id    *int    `json:"id" db:"id"`
	Name  *string `json:"name" db:"name"`
	Genre *string `json:"genre" db:"genre"`
	Data  *string `json:"data" db:"data"`
}

type ActorSelect struct {
	Id      int      `json:"id" db:"id"`
	Name    string   `json:"name" db:"name"`
	Genre   string   `json:"genre" db:"genre"`
	Data    string   `json:"data" db:"data"`
	Cinemas []string `json:"cinemas"`
}

type Cinema struct {
	Id          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Data        string   `json:"data" db:"data"`
	Rating      int      `json:"rating" db:"rating"`
	Actors      []string `json:"actors" db:"actors"`
}

type CinemaUpdate struct {
	Id          *int      `json:"id" db:"id"`
	Name        *string   `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	Data        *string   `json:"data" db:"data"`
	Rating      *int      `json:"rating" db:"rating"`
	Actors      *[]string `json:"actors" db:"actors"`
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

type Search struct {
	Search string `json:"search"`
}
