package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/Futturi/vktest/internal/service"
	mock_service "github.com/Futturi/vktest/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_SignIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthService, user models.User)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "test",
				"password": "test"
			}`,
			inputUser: models.User{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthService, user models.User) {
				s.EXPECT().SignIn(user).Return("token", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"access_token":"token"}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"username": "test",
				"password": "test"
			}`,
			inputUser: models.User{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthService, user models.User) {
				s.EXPECT().SignIn(user).Return("", errors.New("some error"))
			},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
		{
			name: "Internal server error",
			inputBody: `{
				"username": "test",
			}`,
			inputUser: models.User{
				Username: "test",
			},
			mockBehavior:        func(s *mock_service.MockAuthService, user models.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthService(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{AuthService: auth}
			handler := NewHandl(services)

			r := http.NewServeMux()
			r.HandleFunc("/auth/signin", handler.SignIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/signin", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
		})
	}
}

func TestHandler_SignInAdmin(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthService, user models.Admin)
	testTable := []struct {
		name                string
		inputBody           string
		inputAdmin          models.Admin
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "test",
				"password": "test"
			}`,
			inputAdmin: models.Admin{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthService, user models.Admin) {
				s.EXPECT().SignInAdmin(user).Return("token", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"access_token":"token"}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"username": "test",
				"password": "test"
			}`,
			inputAdmin: models.Admin{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthService, user models.Admin) {
				s.EXPECT().SignInAdmin(user).Return("", errors.New("some error"))
			},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
		{
			name: "Internal server error",
			inputBody: `{
				"username": "test",
			}`,
			inputAdmin: models.Admin{
				Username: "test",
			},
			mockBehavior:        func(s *mock_service.MockAuthService, user models.Admin) {},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthService(c)
			testCase.mockBehavior(auth, testCase.inputAdmin)

			services := &service.Service{AuthService: auth}
			handler := NewHandl(services)

			r := http.NewServeMux()
			r.HandleFunc("/auth/admin/signin", handler.SignInAdmin)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/admin/signin", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
		})
	}
}

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthService, user models.User)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "test",
				"password": "test"
			}`,
			inputUser: models.User{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthService, user models.User) {
				s.EXPECT().SignUp(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"username": "test"
			}`,
			mockBehavior:        func(s *mock_service.MockAuthService, user models.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
		{
			name: "Error",
			inputBody: `{
				"rwgjlwejflewjk": "test"
			}`,
			mockBehavior:        func(s *mock_service.MockAuthService, user models.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthService(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{AuthService: auth}
			handler := NewHandl(services)

			r := http.NewServeMux()
			r.HandleFunc("/auth/signup", handler.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
		})
	}
}

func TestHandler_SignUpAdmin(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthService, user models.Admin)
	testTable := []struct {
		name                string
		inputBody           string
		inputAdmin          models.Admin
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "test",
				"password": "test"
			}`,
			inputAdmin: models.Admin{
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthService, user models.Admin) {
				s.EXPECT().SignUpAdmin(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Bad request",
			inputBody: `{
				"username": "test"
			}`,
			mockBehavior:        func(s *mock_service.MockAuthService, user models.Admin) {},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
		{
			name: "Error",
			inputBody: `{
				"reokgjowef;v": "test"
			}`,
			mockBehavior:        func(s *mock_service.MockAuthService, user models.Admin) {},
			expectedStatusCode:  400,
			expectedRequestBody: "",
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthService(c)
			testCase.mockBehavior(auth, testCase.inputAdmin)

			services := &service.Service{AuthService: auth}
			handler := NewHandl(services)

			r := http.NewServeMux()
			r.HandleFunc("/auth/admin/signup", handler.SignUpAdmin)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/admin/signup", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedRequestBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			}
		})
	}
}
