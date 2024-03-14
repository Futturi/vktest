package handler

import (
	"net/http"

	"github.com/Futturi/vktest/internal/service"
)

type Handl struct {
	service *service.Service
}

func NewHandl(service *service.Service) *Handl {
	return &Handl{service: service}
}

func (h *Handl) NewHan() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/signup", h.SignUp)
	mux.HandleFunc("/auth/signin", h.SignIn)
	mux.HandleFunc("/auth/admin/signup", h.SignUpAdmin)
	mux.HandleFunc("/auth/admin/signin", h.SignInAdmin)

	mux.Handle("/api/actors", h.CheckIdentity(http.HandlerFunc(h.GetActors)))
	mux.Handle("/api/actors", h.CheckIdentity(http.HandlerFunc(h.InsertActor)))
	mux.Handle("/api/actors/:id", h.CheckIdentity(http.HandlerFunc(h.UpdateActor)))
	mux.Handle("/api/actors/:id", h.CheckIdentity(http.HandlerFunc(h.DeleteActor)))
	mux.Handle("/api/cinemas", h.CheckIdentity(http.HandlerFunc(h.InsertFilm)))
	return mux
}
