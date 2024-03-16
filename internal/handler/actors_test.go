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

func TestHandler_GetActors(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActorService)
	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockActorService) {
				s.EXPECT().GetActors().Return([]models.ActorSelect{}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mock := mock_service.NewMockActorService(c)
			testCase.mockBehavior(mock)

			services := &service.Service{ActorService: mock}

			handler := NewHandl(services)

			r := httptest.NewRequest("GET", "/api/actors", nil)
			w := httptest.NewRecorder()
			handler.GetActors(w, r)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_InsertActor(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActorService, actor models.Actor)
	testTable := []struct {
		name                string
		inputBody           string
		adminHeader         string
		adminHeaderValue    string
		inputActor          models.Actor
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "test",
				"genre": "test",
				"data": "test"
			}`,
			adminHeader:      "isAdmin",
			adminHeaderValue: "true",
			inputActor: models.Actor{
				Name:  "test",
				Genre: "test",
				Data:  "test",
			},
			mockBehavior: func(s *mock_service.MockActorService, actor models.Actor) {
				s.EXPECT().InsertActor(actor).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"name": "test",
				"genre": "test",
				"data": "test"
			}`,
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockActorService, actor models.Actor) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
		{
			name: "Unauthorized",
			inputBody: `{
				"jwdgfkjwegjl": "test",
				"qwoje": "test",
				"datfpgokoeka": "test"
			}`,
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockActorService, actor models.Actor) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockActorService(c)
			testCase.mockBehavior(mock, testCase.inputActor)
			services := &service.Service{ActorService: mock}
			handler := NewHandl(services)

			r := http.NewServeMux()
			r.HandleFunc("/api/actors", handler.InsertActor)

			req := httptest.NewRequest("POST", "/api/actors", bytes.NewBufferString(testCase.inputBody))
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

func TestHandler_UpdateActor(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActorService, actor models.ActorUpdate)
	testTable := []struct {
		name                string
		inputBody           string
		adminHeader         string
		adminHeaderValue    string
		inputActor          models.Actor
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "test",
				"data": "test"
			}`,
			adminHeader:      "isAdmin",
			adminHeaderValue: "true",
			inputActor: models.Actor{
				Name: "test",
				Data: "test",
			},
			mockBehavior: func(s *mock_service.MockActorService, actor models.ActorUpdate) {
				s.EXPECT().UpdateActor(1, actor).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"name": "test",
				"data": "test"
			}`,
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockActorService, actor models.ActorUpdate) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
		{
			name: "Unauthorized",
			inputBody: `{
				"jwdgfkjwegjl": "test",
				"qwoje": "test",
				"datfpgokoeka": "test"
			}`,
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockActorService, actor models.ActorUpdate) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockActorService(c)
			testCase.mockBehavior(mock, models.ActorUpdate{
				Name: &testCase.inputActor.Name,
				Data: &testCase.inputActor.Data,
			})
			services := &service.Service{ActorService: mock}
			handler := NewHandl(services)

			r := http.NewServeMux()
			r.HandleFunc("/api/actors", handler.UpdateActor)

			req := httptest.NewRequest("PUT", "/api/actors?id=1", bytes.NewBufferString(testCase.inputBody))
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

func TestHandler_DeleteActor(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActorService)
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
			mockBehavior: func(s *mock_service.MockActorService) {
				s.EXPECT().DeleteActor("1").Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":"1"}`,
		},
		{
			name:                "Bad request",
			adminHeader:         "isAdmin",
			adminHeaderValue:    "",
			mockBehavior:        func(s *mock_service.MockActorService) {},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockActorService(c)
			testCase.mockBehavior(mock)
			services := &service.Service{ActorService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.HandleFunc("/api/actors", handler.DeleteActor)

			req := httptest.NewRequest("DELETE", "/api/actors?id=1", nil)
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
