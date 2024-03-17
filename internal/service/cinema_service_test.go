package service

import (
	"errors"
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/repository"
	mock_repository "github.com/Futturi/vktest/internal/repository/mocksr"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestNewCinemaService(t *testing.T) {
	repo := &repository.Repository{}
	s := NewCinemaService(repo)
	assert.Equal(t, s.repo, repo)
}

func TestCinemaService_InsertCinema(t *testing.T) {
	type mockBehaviour func(s *mock_repository.MockCinemaRepo, cinema models.Cinema)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		cinema        models.Cinema
		expected      int
		err           error
	}{
		{
			cinema: models.Cinema{Name: "name", Description: "desc", Data: "2006-Jan-02", Rating: 1},
			name:   "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo, cinema models.Cinema) {
				s.EXPECT().InsertCinema(models.Cinema{
					Name:        "name",
					Description: "desc",
					Data:        "1136160000",
					Rating:      1,
				}).Return(1, nil)
			},
			expected: 1,
			err:      nil,
		},
		{
			cinema: models.Cinema{Name: "oiejoqjfpijvpijqpij1pijqfpijeqfiweklmf;wklemfkl;ewmfkewklfmwepofpoakmzkvmlkwmle;k;laskdclmvknvkenrblkkalsm.,cm lklerlkmlksdmlkfmlkvmlkbnerlkmsdklm;flxmvkwmkmsedlkmfvnbnrghtaosfmkrwkektmfsamkfmvnbkrejwepjrksafmvnberkjkwel;sakdmfvngmjrekdsmcnvgfjpjqwpijqwpdjqwfjpeqgmoeqngpeqnopqenvko", Description: "desc", Data: "2006-Jan-02", Rating: 1},
			name:   "Error",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo, cinema models.Cinema) {
			},
			expected: 0,
			err:      errors.New("your name is longer than 150"),
		},
		{
			cinema: models.Cinema{Name: "name", Description: "lkfdm", Data: "2006-Jan-02", Rating: -1},
			name:   "Error2",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo, cinema models.Cinema) {
			},
			expected: 0,
			err:      errors.New("incorrect rating"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_repository.NewMockCinemaRepo(c)
			testCase.mockBehaviour(mock, testCase.cinema)

			repository := &repository.Repository{CinemaRepo: mock}

			s := NewCinemaService(repository)

			_, err := s.InsertCinema(testCase.cinema)

			assert.Equal(t, testCase.err, err)
		})
	}
}

func TestCinemaService_GetCinemas(t *testing.T) {
	type mockBehaviour func(s *mock_repository.MockCinemaRepo)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		expected      []models.Cinema
		sort          string
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo) {
				s.EXPECT().GetCinemas("").Return([]models.Cinema{}, nil)
			},
			expected: []models.Cinema{},
			sort:     "",
		},
		{
			name: "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo) {
				s.EXPECT().GetCinemas("rating").Return([]models.Cinema{}, nil)
			},
			expected: []models.Cinema{},
			sort:     "rating",
		},
		{
			name: "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo) {
				s.EXPECT().GetCinemas("name").Return([]models.Cinema{}, nil)
			},
			expected: []models.Cinema{},
			sort:     "name",
		},
		{
			name: "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo) {
				s.EXPECT().GetCinemas("date").Return([]models.Cinema{}, nil)
			},
			expected: []models.Cinema{},
			sort:     "date",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_repository.NewMockCinemaRepo(c)
			testCase.mockBehaviour(mock)

			repository := &repository.Repository{CinemaRepo: mock}

			s := NewCinemaService(repository)

			_, err := s.GetCinemas(testCase.sort)

			assert.Equal(t, nil, err)
		})
	}
}

func TestCinemaService_DeleteCinema(t *testing.T) {
	type mockBehaviour func(s *mock_repository.MockCinemaRepo)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		id            string
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo) {
				s.EXPECT().DeleteFilm("1").Return(nil)
			},
			id: "1",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_repository.NewMockCinemaRepo(c)
			testCase.mockBehaviour(mock)
			repository := &repository.Repository{CinemaRepo: mock}

			s := NewCinemaService(repository)

			err := s.DeleteFilm(testCase.id)

			assert.Equal(t, nil, err)
		})
	}
}

func TestCinemaService_UpdateCinema(t *testing.T) {
	type mockBehaviour func(s *mock_repository.MockCinemaRepo, cinema models.CinemaUpdate)
	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		cinema        models.Cinema
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_repository.MockCinemaRepo, cinema models.CinemaUpdate) {
				s.EXPECT().UpdateFilm("1", cinema).Return(nil)
			},
			cinema: models.Cinema{
				Name:        "name",
				Description: "desc",
				Rating:      1,
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			a := "0"
			mock := mock_repository.NewMockCinemaRepo(c)
			testCase.mockBehaviour(mock, models.CinemaUpdate{
				Name:        &testCase.cinema.Name,
				Description: &testCase.cinema.Description,
				Rating:      &testCase.cinema.Rating,
				Data:        &a,
			})

			repository := &repository.Repository{CinemaRepo: mock}

			s := NewCinemaService(repository)

			err := s.UpdateFilm("1", models.CinemaUpdate{
				Name:        &testCase.cinema.Name,
				Description: &testCase.cinema.Description,
				Rating:      &testCase.cinema.Rating,
			})

			assert.Equal(t, nil, err)
		})
	}
}
