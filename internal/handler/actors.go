package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

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

func (h *Handl) UpdateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if !h.GetPrivileage(r) {
			http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		}
		url := r.URL.Path
		splited := strings.Split(url, "/")
		id := splited[len(splited)-1]
		byt, err := io.ReadAll(r.Body)

		var actor models.ActorUpdate

		if err != nil {
			slog.Error("error while inserting actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}

		json.Unmarshal(byt, &actor)

		newid, err := strconv.Atoi(id)
		if err != nil {
			slog.Error("error while converting string to ind", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}

		err = h.service.UpdateActor(newid, actor)
		if err != nil {
			slog.Error("error while updating actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusInternalServerError)
		}
		byt2, err := json.Marshal(map[string]int{"id": newid})
		if err != nil {
			slog.Error("error while marshalling result", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		w.Write(byt2)
	}
}

func (h *Handl) DeleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if !h.GetPrivileage(r) {
			http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		}
		url := r.URL.Path
		splited := strings.Split(url, "/")
		id := splited[len(splited)-1]
		newid, err := strconv.Atoi(id)
		if err != nil {
			slog.Error("error while converting string to ind", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		err = h.service.DeleteActor(newid)
		if err != nil {
			slog.Error("error while deleting actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		byt2, err := json.Marshal(map[string]int{"id": newid})
		if err != nil {
			slog.Error("error while marshalling result", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		w.Write(byt2)
	}
}
