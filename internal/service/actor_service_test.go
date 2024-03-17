package service

import (
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
	mock_repo "github.com/Futturi/vktest/internal/repository/mocksr"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestNewActorService(t *testing.T) {
	repo := &repository.Repository{}
	s := NewActorService(repo)
	assert.Equal(t, s.repo, repo)
}

func TestActorService_GetActors(t *testing.T) {
	type mockBehaviour func(s *mock_repo.MockActorRepo)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		expected      []models.ActorSelect
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repo.MockActorRepo) {
				s.EXPECT().GetActors().Return([]models.ActorSelect{}, nil)
			},
			expected: []models.ActorSelect{},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mock := mock_repo.NewMockActorRepo(c)
			testCase.mockBehaviour(mock)

			repository := &repository.Repository{ActorRepo: mock}

			s := NewActorService(repository)

			_, err := s.GetActors()

			assert.Equal(t, nil, err)
		})
	}
}

func TestActorService_InsertActor(t *testing.T) {
	type mockBehaviour func(s *mock_repo.MockActorRepo)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		actor         models.Actor
		expected      int
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repo.MockActorRepo) {
				s.EXPECT().InsertActor(models.Actor{Name: "test",
					Genre: "test",
					Data:  "1136160000"}).Return(0, nil)
			},
			actor: models.Actor{
				Name:  "test",
				Genre: "test",
				Data:  "2006-Jan-02",
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_repo.NewMockActorRepo(c)
			testCase.mockBehaviour(mock)

			repository := &repository.Repository{ActorRepo: mock}

			s := NewActorService(repository)

			got, err := s.InsertActor(testCase.actor)

			assert.Equal(t, testCase.expected, got)
			assert.Equal(t, nil, err)
		})
	}
}

func TestActorService_DeleteActor(t *testing.T) {
	type mockBehaviour func(s *mock_repo.MockActorRepo)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		id            string
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repo.MockActorRepo) {
				s.EXPECT().DeleteActor("1").Return(nil)
			},
			id: "1",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_repo.NewMockActorRepo(c)
			testCase.mockBehaviour(mock)
			repository := &repository.Repository{ActorRepo: mock}

			s := NewActorService(repository)

			err := s.DeleteActor(testCase.id)

			assert.Equal(t, nil, err)
		})
	}
}

func TestActorService_UpdateActor(t *testing.T) {
	type mockBehaviour func(s *mock_repo.MockActorRepo, actor models.ActorUpdate)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		actor         models.Actor
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repo.MockActorRepo, actor models.ActorUpdate) {
				s.EXPECT().UpdateActor(1, actor).Return(nil)
			},
			actor: models.Actor{
				Name:  "test",
				Genre: "test",
				Data:  "2006-Jan-02",
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_repo.NewMockActorRepo(c)
			testCase.mockBehaviour(mock, models.ActorUpdate{
				Name:  &testCase.actor.Name,
				Genre: &testCase.actor.Genre,
			})

			repository := &repository.Repository{ActorRepo: mock}

			s := NewActorService(repository)

			err := s.UpdateActor(1, models.ActorUpdate{
				Name:  &testCase.actor.Name,
				Genre: &testCase.actor.Genre,
			})

			assert.Equal(t, nil, err)
		})
	}
}
