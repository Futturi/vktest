package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Futturi/vktest/internal/err"
)

const (
	userHeader  = "userid"
	adminHeader = "isAdmin"
)

func (h *Handl) CheckIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			err.NewErr(errors.New("empty auth header"))
			slog.Error("error", errors.New("empty auth header"))
		}
		headerParts := strings.Split(auth, " ")
		if len(headerParts) != 2 {
			err.NewErr(errors.New("invalid header"))
			slog.Error("error", errors.New("invalid header"))
		}

		userId, IsAdmin, er := h.service.ParseToken(headerParts[1])
		if er != nil {
			err.NewErr(er)
			slog.Error("error", er)
		}
		var isadmin string
		if IsAdmin {
			isadmin = "true"
		} else {
			isadmin = "false"
		}
		r.Header.Add(userHeader, userId)
		r.Header.Add(adminHeader, isadmin)
		w.Header().Add(userHeader, userId)
		w.Header().Add(adminHeader, isadmin)
		next.ServeHTTP(w, r)
	})
}

func (h *Handl) GetPrivileage(r *http.Request) bool {
	header := r.Header.Get(adminHeader)
	return header == "true"
}
