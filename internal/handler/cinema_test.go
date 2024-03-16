package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/service"
	mock_service "github.com/Futturi/vktest/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_InsertFilm(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCinemaService, cinema models.Cinema)
	testTable := []struct {
		name                string
		inputBody           string
		adminHeader         string
		adminHeaderValue    string
		inputCinema         models.Cinema
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "test",
				"description": "test",
				"data": "test",
				"rating": 10,
				"actors": ["test"]
			}`,
			adminHeader:      "isAdmin",
			adminHeaderValue: "true",
			inputCinema: models.Cinema{
				Name:        "test",
				Description: "test",
				Rating:      10,
				Actors:      []string{"test"},
				Data:        "test",
			},
			mockBehavior: func(s *mock_service.MockCinemaService, cinema models.Cinema) {
				s.EXPECT().InsertCinema(cinema).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"name": "test",
				"description": "test",
				"data": "test",
				"rating": 10,
				"actors": ["test"]
			}`,
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockCinemaService, cinema models.Cinema) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
		{
			name:                "Bad request",
			inputBody:           ``,
			adminHeader:         "isAdmin",
			adminHeaderValue:    "true",
			mockBehavior:        func(s *mock_service.MockCinemaService, cinema models.Cinema) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockCinemaService(c)
			testCase.mockBehavior(mock, testCase.inputCinema)
			services := &service.Service{CinemaService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.HandleFunc("/api/cinemas", handler.InsertFilm)
			req := httptest.NewRequest("POST", "/api/cinemas", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("isAdmin", testCase.adminHeaderValue)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
			assert.Equal(t, testCase.adminHeaderValue, req.Header.Get("isAdmin"))
		})
	}
}

func TestHandler_GetFilms(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCinemaService)
	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockCinemaService) {
				s.EXPECT().GetCinemas("rating").Return([]models.Cinema{}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockCinemaService(c)
			testCase.mockBehavior(mock)
			services := &service.Service{CinemaService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.HandleFunc("/api/cinemas", handler.GetFilms)
			req := httptest.NewRequest("GET", "/api/cinemas?sort=rating", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteFilm(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCinemaService)
	testTable := []struct {
		name                string
		adminHeader         string
		adminHeaderValue    string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:             "OK",
			adminHeader:      "isAdmin",
			adminHeaderValue: "true",
			mockBehavior: func(s *mock_service.MockCinemaService) {
				s.EXPECT().DeleteFilm("1").Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":"1"}`,
		},
		{
			name:                "Bad request",
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockCinemaService) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockCinemaService(c)
			testCase.mockBehavior(mock)
			services := &service.Service{CinemaService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.HandleFunc("/api/cinemas", handler.DeleteFilm)
			req := httptest.NewRequest("DELETE", "/api/cinemas?id=1", nil)
			req.Header.Set("isAdmin", testCase.adminHeaderValue)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
			assert.Equal(t, testCase.adminHeaderValue, req.Header.Get("isAdmin"))
		})
	}
}

func TestHandler_UpdateFilm(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCinemaService, film models.CinemaUpdate)
	testTable := []struct {
		name                string
		inputBody           string
		adminHeader         string
		adminHeaderValue    string
		inputFilm           models.Cinema
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "test",
				"description": "test",
				"data": "test"
			}`,
			adminHeader:      "isAdmin",
			adminHeaderValue: "true",
			inputFilm: models.Cinema{
				Name:        "test",
				Description: "test",
				Data:        "test",
			},
			mockBehavior: func(s *mock_service.MockCinemaService, film models.CinemaUpdate) {
				s.EXPECT().UpdateFilm("1", film).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":"1"}`,
		},
		{
			name:                "Bad request",
			inputBody:           "",
			adminHeader:         "isAdmin",
			adminHeaderValue:    "true",
			inputFilm:           models.Cinema{},
			mockBehavior:        func(s *mock_service.MockCinemaService, film models.CinemaUpdate) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
		{
			name:                "Bad request",
			inputBody:           "",
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			inputFilm:           models.Cinema{},
			mockBehavior:        func(s *mock_service.MockCinemaService, film models.CinemaUpdate) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockCinemaService(c)
			testCase.mockBehavior(mock, models.CinemaUpdate{
				Name:        &testCase.inputFilm.Name,
				Description: &testCase.inputFilm.Description,
				Data:        &testCase.inputFilm.Data,
			})
			services := &service.Service{CinemaService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.HandleFunc("/api/cinemas", handler.UpdateFilm)
			req := httptest.NewRequest("PUT", "/api/cinemas?id=1", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("isAdmin", testCase.adminHeaderValue)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
			assert.Equal(t, testCase.adminHeaderValue, req.Header.Get("isAdmin"))
		})
	}
}

func TestHandler_Search(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCinemaService, search models.Search)
	testTable := []struct {
		name                string
		inputBody           string
		inputSearcb         models.Search
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"search": "test"
			}`,
			inputSearcb: models.Search{
				Search: "test",
			},
			mockBehavior: func(s *mock_service.MockCinemaService, search models.Search) {
				s.EXPECT().Search(search).Return([]models.Cinema{}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockCinemaService(c)
			testCase.mockBehavior(mock, testCase.inputSearcb)
			services := &service.Service{CinemaService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.HandleFunc("/api/cinemas/search", handler.Search)
			req := httptest.NewRequest("POST", "/api/cinemas/search", bytes.NewBufferString(testCase.inputBody))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
