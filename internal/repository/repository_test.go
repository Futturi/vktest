package repository

import (
	"testing"

	"github.com/magiconair/properties/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestRepository(t *testing.T) {
	db, _, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewRepostitory(db)

	assert.Equal(t, r, &Repository{Authorization: NewAuthRepo(db), ActorRepo: NewActorRepo(db), CinemaRepo: NewCinemaRepo(db)})
}
