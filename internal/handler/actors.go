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
// @Security ApiKeyAuth
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
			return
		}
		byt, _ := json.Marshal(actors)
		w.Header().Set("Content-Type", "application/json")
		w.Write(byt)
		slog.Info("got actors", slog.Any("actors", actors))
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
// @Security ApiKeyAuth
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
		return
	} else {
		var body models.Actor

		byt, _ := io.ReadAll(r.Body)
		json.Unmarshal(byt, &body)

		id, err := h.service.InsertActor(body)
		if err != nil {
			slog.Error("error while inserting actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		mapa := map[string]int{"id": id}
		byt2, _ := json.Marshal(mapa)
		w.Header().Set("Content-Type", "application/json")
		slog.Info("inserted actor", slog.Any("actor", body))
		w.Write(byt2)
	}
}

// @Summary UpdateActor
// @Security ApiKeyAuth
// @Tags actors
// @Description update actor
// @ID update-actor
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure default {string} error
// @Router /api/actors/{id} [put]
func (h *Handl) UpdateActor(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		return
	} else {
		id := r.URL.Query().Get("id")
		byt, _ := io.ReadAll(r.Body)

		var actor models.ActorUpdate

		err := json.Unmarshal(byt, &actor)
		if err != nil {
			slog.Error("error while unmarshalling body", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}

		newid, err := strconv.Atoi(id)
		if err != nil {
			slog.Error("error while converting string to ind", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}

		err = h.service.UpdateActor(newid, actor)
		if err != nil {
			slog.Error("error while updating actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		slog.Info("updated actor", slog.Any("actor", actor))
		byt2, _ := json.Marshal(map[string]int{"id": newid})
		w.Header().Set("Content-Type", "application/json")
		w.Write(byt2)
	}

}

// @Summary DeleteActor
// @Security ApiKeyAuth
// @Tags actors
// @Description delete actor
// @ID delete-actor
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure default {string} error
// @Router /api/actors/{id} [delete]
func (h *Handl) DeleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if !h.GetPrivileage(r) {
			http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
			return
		} else {
			newid := r.URL.Query().Get("id")
			err := h.service.DeleteActor(newid)
			if err != nil {
				slog.Error("error while deleting actor", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
				return
			}
			byt2, _ := json.Marshal(map[string]string{"id": newid})
			w.Header().Set("Content-Type", "application/json")
			w.Write(byt2)
			slog.Info("deleted actor", slog.Any("id", newid))
		}

	}
}
