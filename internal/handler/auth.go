package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/models"
)

func (h *Handl) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "incorrect methor", http.StatusBadRequest)
	}
	var User models.User
	byt := r.Body
	bytes, err := io.ReadAll(byt)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
	}
	json.Unmarshal(bytes, &User)
	id, err := h.service.SignUp(User)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with sign up", http.StatusBadRequest)
	}
	ret := models.User{Id: id}
	str, err := json.Marshal(ret)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
	}
	w.Write(str)
	slog.Info("created user with ", slog.Any("values", ret))
}

func (h *Handl) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "incorrect methor", http.StatusBadRequest)
	}
	var User models.User
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
	}
	json.Unmarshal(bytes, &User)

	token, err := h.service.SignIn(User)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with sign in", http.StatusBadRequest)
	}
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
	}
	Token := models.Token{Access: token}
	byt, err := json.Marshal(Token)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
	}
	w.Write(byt)
	slog.Info("user signed in with", slog.String("token", Token.Access))
}

func (h *Handl) SignUpAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "incorrect methor", http.StatusBadRequest)
	}
	var Admin models.Admin
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
	}
	json.Unmarshal(bytes, &Admin)
	id, err := h.service.SignUpAdmin(Admin)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with sign up", http.StatusBadRequest)
	}
	ret := models.User{Id: id}
	str, err := json.Marshal(ret)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
	}
	w.Write(str)
	slog.Info("created admin with id: ", slog.Any("id", id))
}

func (h *Handl) SignInAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "incorrect methor", http.StatusBadRequest)
	}
	var Admin models.Admin
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
	}
	json.Unmarshal(bytes, &Admin)
	token, err := h.service.SignInAdmin(Admin)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with sign in", http.StatusBadRequest)
	}
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		str, err := json.Marshal(err)
		if err != nil {
			slog.Error("error while marshalling body", slog.Any("error", err))
		}
		w.Write(str)
	}
	Token := models.Token{Access: token}
	byt, err := json.Marshal(Token)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
	}
	w.Write(byt)
	slog.Info("admin signed in with", slog.String("token", Token.Access))
}
