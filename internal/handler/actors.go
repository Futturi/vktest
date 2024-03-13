package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/models"
)

func (h *Handl) GetActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.service.GetActors()
	if err != nil {
		slog.Error("error while giving actors", slog.Any("error", err))
		http.Error(w, "error", http.StatusInternalServerError)
	}
	byt, err := json.Marshal(actors)
	if err != nil {
		slog.Error("error while marshalling actors", slog.Any("error", err))
		http.Error(w, "error", http.StatusInternalServerError)
	}
	w.Write(byt)
}

func (h *Handl) InsertActor(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "you have not priviliage for doing this", http.StatusBadRequest)
	}
	var body models.Actor
	id, err := h.service.InsertActor(body)
	if err != nil {
		slog.Error("error while inserting actor", slog.Any("error", err))
		http.Error(w, "error", http.StatusInternalServerError)
	}
	mapa := map[string]int{"id": id}
	byt, err := json.Marshal(mapa)
	if err != nil {
		slog.Error("error while marshall actor", slog.Any("error", err))
		http.Error(w, "error", http.StatusInternalServerError)
	}
	w.Write(byt)
}
