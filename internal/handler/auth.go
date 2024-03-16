package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/errs"
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
		http.Error(w, "incorrect method", http.StatusBadRequest)
		return
	}
	var User models.User
	byt := r.Body
	bytes, err := io.ReadAll(byt)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, string(errs.NewErr(errors.New("invalid body"))), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bytes, &User)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, string(errs.NewErr(errors.New("invalid body"))), http.StatusBadRequest)
		return
	}
	if User.Password == "" || User.Username == "" {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, string(errs.NewErr(errors.New("invalid body"))), http.StatusBadRequest)
		return
	}
	id, err := h.service.SignUp(User)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, string(errs.NewErr(errors.New("invalid body"))), http.StatusBadRequest)
		return
	}
	ret := map[string]interface{}{"id": id}
	str, err := json.Marshal(ret)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
		return
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
		return
	}
	var User models.User
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bytes, &User)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}

	if User.Password == "" || User.Username == "" {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}

	token, err := h.service.SignIn(User)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with sign in", http.StatusBadRequest)
		return
	}
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}
	Token := models.Token{Access: token}
	byt, err := json.Marshal(Token)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
		return
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
		return
	}
	var Admin models.Admin
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bytes, &Admin)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}
	if Admin.Password == "" || Admin.Username == "" {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}
	id, err := h.service.SignUpAdmin(Admin)
	if err != nil {
		slog.Error("error with sign up", slog.Any("error", err))
		http.Error(w, "error with sign up", http.StatusBadRequest)
		return
	}
	ret := map[string]interface{}{"id": id}
	str, err := json.Marshal(ret)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
		return
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
		return
	}
	var Admin models.Admin
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bytes, &Admin)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}

	if Admin.Password == "" || Admin.Username == "" {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with body", http.StatusBadRequest)
		return
	}

	token, err := h.service.SignInAdmin(Admin)
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		http.Error(w, "error with sign in", http.StatusBadRequest)
		return
	}
	if err != nil {
		slog.Error("error with signing in", slog.Any("error", err))
		str, err := json.Marshal(err)
		if err != nil {
			slog.Error("error while marshalling body", slog.Any("error", err))
			return
		}
		w.Write(str)
	}
	Token := models.Token{Access: token}
	byt, err := json.Marshal(Token)
	if err != nil {
		slog.Error("error while marshalling body", slog.Any("error", err))
		return
	}
	w.Write(byt)
	slog.Info("admin signed in with", slog.String("token", Token.Access))
}
