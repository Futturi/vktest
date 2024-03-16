package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/models"
)

// @Summary SingUpUser
// @Tags auth
// @Description create account 4 user
// @ID create-account-user
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure default {string} error with body
// @Router /auth/signup [post]
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

// @Summary SingInUser
// @Tags auth
// @Description login account 4 user
// @ID login-account-user
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
// @Success 200 {string} token
// @Failure default {string} error with body
// @Router /auth/signin [post]
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

// @Summary SingUpAdmin
// @Tags authAdmin
// @Description create account 4 admin
// @ID create-account-admin
// @Accept json
// @Produce json
// @Param input body models.Admin true "account info"
// @Success 200 {integer} integer 1
// @Failure default {string} error with body
// @Router /auth/admin/signup [post]
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

// @Summary SingInAdmin
// @Tags authAdmin
// @Description login account 4 admin
// @ID login-account-admin
// @Accept json
// @Produce json
// @Param input body models.Admin true "account info"
// @Success 200 {string} token
// @Failure default {string} error with body
// @Router /auth/admin/signin [post]
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
