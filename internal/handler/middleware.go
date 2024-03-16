package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Futturi/vktest/internal/errs"
)

const (
	userHeader  = "userid"
	adminHeader = "isAdmin"
)

func (h *Handl) CheckIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			errs.NewErr(errors.New("empty auth header"))
			slog.Error("error", errors.New("empty auth header"))
			return
		}
		headerParts := strings.Split(auth, " ")
		if len(headerParts) != 2 {
			http.Error(w, "invalid header", http.StatusUnauthorized)
			errs.NewErr(errors.New("invalid header"))
			slog.Error("error", errors.New("invalid header"))
			return
		}

		if headerParts[0] != "Bearer" {
			http.Error(w, "invalid header", http.StatusUnauthorized)
			errs.NewErr(errors.New("invalid header"))
			slog.Error("error", errors.New("invalid header"))
			return
		}
		if headerParts[1] == "" {
			http.Error(w, "empty token", http.StatusUnauthorized)
			errs.NewErr(errors.New("empty token"))
			slog.Error("error", errors.New("empty token"))
			return
		}

		userId, IsAdmin, er := h.service.ParseToken(headerParts[1])
		if er != nil {
			errs.NewErr(er)
			slog.Error("error", er)
			return
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
