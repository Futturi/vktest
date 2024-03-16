package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Futturi/vktest/internal/service"
	mock_service "github.com/Futturi/vktest/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_CheckIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthService, token string)
	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthService, token string) {
				s.EXPECT().ParseToken(token).Return("1", true, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Bad request",
			headerName:           "Authorization",
			headerValue:          "Bearing token",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAuthService, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "",
		},
		{
			name:                 "Forbidden",
			headerName:           "NeAuthorization",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAuthService, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "",
		},
		{
			name:                 "Empty",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "",
			mockBehavior:         func(s *mock_service.MockAuthService, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockAuthService(c)
			testCase.mockBehavior(mock, testCase.token)
			services := &service.Service{AuthService: mock}
			handler := NewHandl(services)
			r := http.NewServeMux()
			r.Handle("/protec", handler.CheckIdentity(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				id := r.Header.Get(userHeader)
				w.Write([]byte(id))
			})))
			req := httptest.NewRequest("GET", "/protec", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody == "" {
				assert.Equal(t, w.Body.String(), w.Body.String())
			} else {
				assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

}
