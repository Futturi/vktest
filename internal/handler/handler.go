package handler

import (
	"net/http"

	_ "github.com/Futturi/vktest/docs"
	"github.com/Futturi/vktest/internal/service"
	"github.com/swaggo/http-swagger/v2"
)

type Handl struct {
	service *service.Service
}

func NewHandl(service *service.Service) *Handl {
	return &Handl{service: service}
}

func (h *Handl) NewHan() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	mux.HandleFunc("/auth/signup", h.SignUp)
	mux.HandleFunc("/auth/signin", h.SignIn)
	mux.HandleFunc("/auth/admin/signup", h.SignUpAdmin)
	mux.HandleFunc("/auth/admin/signin", h.SignInAdmin)
	mux.Handle("/api/actors", h.CheckIdentity(http.HandlerFunc(h.GetActors)))
	mux.Handle("/api/actors/", h.CheckIdentity(http.HandlerFunc(h.UpdateActor)))
	mux.Handle("/api/cinemas", h.CheckIdentity(http.HandlerFunc(h.InsertFilm)))
	mux.Handle("/api/cinemas/search", h.CheckIdentity(http.HandlerFunc(h.Search)))
	return mux
}
