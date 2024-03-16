package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Futturi/vktest/internal/models"
)

// @Summary GetAllActors
// @Secutiry ApiKeyAuth
// @Tags actors
// @Description get all actors
// @ID get-actors
// @Produce json
// @Success 200 {object} models.Actor
// @Failure default {string} error
// @Router /api/actors [get]
func (h *Handl) GetActors(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
	if r.Method == "POST" {
		h.InsertActor(w, r)
	}
	if r.Method == "PUT" {
		h.UpdateActor(w, r)
	}
	if r.Method == "DELETE" {
		h.DeleteActor(w, r)
	}
}

// @Summary InsertActor
// @Secutiry ApiKeyAuth
// @Tags actors
// @Description insert actor
// @ID insert-actor
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure default {string} error
// @Router /api/actors [post]
func (h *Handl) InsertActor(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "you have not priviliage for doing this", http.StatusBadRequest)
	} else {
		var body models.Actor

		byt, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("error with body", slog.Any("error", err))
			http.Error(w, "error", http.StatusInternalServerError)
		}
		json.Unmarshal(byt, &body)

		id, err := h.service.InsertActor(body)
		if err != nil {
			slog.Error("error while inserting actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusInternalServerError)
		}
		mapa := map[string]int{"id": id}
		byt2, err := json.Marshal(mapa)
		if err != nil {
			slog.Error("error while marshall actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusInternalServerError)
		}
		w.Write(byt2)
	}
}

// @Summary UpdateActor
// @Secutiry ApiKeyAuth
// @Tags actors
// @Description update actor
// @ID update-actor
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure default {string} error
// @Router /api/actors [put]
func (h *Handl) UpdateActor(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
	} else {
		id := r.URL.Query().Get("id")
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

// @Summary DeleteActor
// @Secutiry ApiKeyAuth
// @Tags actors
// @Description delete actor
// @ID delete-actor
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure default {string} error
// @Router /api/actors [delete]
func (h *Handl) DeleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if !h.GetPrivileage(r) {
			http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		} else {
			newid := r.URL.Query().Get("id")
			err := h.service.DeleteActor(newid)
			if err != nil {
				slog.Error("error while deleting actor", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
			}
			byt2, err := json.Marshal(map[string]string{"id": newid})
			if err != nil {
				slog.Error("error while marshalling result", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
			}
			w.Write(byt2)
		}

	}
}
