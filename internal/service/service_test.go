package service

import (
	"testing"

	"github.com/Futturi/vktest/internal/repository"
	mock_repository "github.com/Futturi/vktest/internal/repository/mocksr"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestService(t *testing.T) {

	c := gomock.NewController(t)
	defer c.Finish()
	mock1 := mock_repository.NewMockCinemaRepo(c)
	mock2 := mock_repository.NewMockActorRepo(c)
	mock3 := mock_repository.NewMockAuthorization(c)

	s := NewService(&repository.Repository{CinemaRepo: mock1, ActorRepo: mock2, Authorization: mock3})
	assert.Equal(t, s, &Service{AuthService: NewAuthService(mock3), CinemaService: NewCinemaService(mock1), ActorService: NewActorService(mock2)})
}
