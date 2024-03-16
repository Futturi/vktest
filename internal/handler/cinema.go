package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/models"
)

// @Summary InsertCinema
// @Secutiry ApiKeyAuth
// @Tags cinemas
// @Description insert cinema
// @ID insert-cinemas
// @Accept json
// @Produce json
// @Success 200 {string} id
// @Failure default {string} error
// @Router /api/cinemas [post]
func (h *Handl) InsertFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !h.GetPrivileage(r) {
			http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
			return
		} else {
			byt, err := io.ReadAll(r.Body)
			if err != nil {
				slog.Error("error while inserting actor", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
				return
			}
			var cinema models.Cinema
			err = json.Unmarshal(byt, &cinema)
			if err != nil {
				slog.Error("error while unmarshalling body", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
				return
			}
			slog.Info(cinema.Data)
			id, err := h.service.InsertCinema(cinema)
			if err != nil {
				slog.Error("error while inserting body", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
				return
			}
			byt2, err := json.Marshal(map[string]int{"id": id})
			if err != nil {
				slog.Error("error while marshalling result", slog.Any("error", err))
				http.Error(w, "error", http.StatusInternalServerError)
				return
			}
			w.Write(byt2)
		}
	}
	if r.Method == "PUT" {
		h.UpdateFilm(w, r)
	}
	if r.Method == "DELETE" {
		h.DeleteFilm(w, r)
	}
	if r.Method == "GET" {
		h.GetFilms(w, r)
	}
}

// @Summary UpdateCinema
// @Secutiry ApiKeyAuth
// @Tags cinemas
// @Description update cinema
// @ID update-cinemas
// @Accept json
// @Produce json
// @Success 200 {string} id
// @Failure default {string} error
// @Router /api/cinemas{id} [put]
func (h *Handl) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		return
	} else {
		var cinema models.CinemaUpdate

		byt, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("error while inserting actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byt, &cinema)
		if err != nil {
			slog.Error("error while unmarshalling body", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}

		id := r.URL.Query().Get("id")
		err = h.service.UpdateFilm(id, cinema)
		if err != nil {
			slog.Error("error while inserting body", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		byt2, err := json.Marshal(map[string]string{"id": id})
		if err != nil {
			slog.Error("error while marshalling body", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		w.Write(byt2)
	}
}

// @Summary DeleteCinemas
// @Secutiry ApiKeyAuth
// @Tags cinemas
// @Description delete cinema
// @ID get-cinemas
// @Produce json
// @Success 200 {string} id
// @Failure default {string} error
// @Router /api/cinemas{id} [delete]
func (h *Handl) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		return
	} else {
		id := r.URL.Query().Get("id")
		err := h.service.DeleteFilm(id)
		if err != nil {
			slog.Error("error while deleting cinema", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		byt, err := json.Marshal(map[string]string{"id": id})
		if err != nil {
			slog.Error("error while marshalling result", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		w.Write(byt)
	}
}

// @Summary GetAllCinemas
// @Secutiry ApiKeyAuth
// @Tags cinemas
// @Description get all cinemas
// @ID get-cinemas
// @Produce json
// @Success 200 {object} []models.Cinema
// @Failure default {string} error
// @Router /api/cinemas{sort} [get]
func (h *Handl) GetFilms(w http.ResponseWriter, r *http.Request) {
	var sor string
	switch {
	case !r.URL.Query().Has("sort") || r.URL.Query().Get("sort") == "rating":
		sor = "rating"
	case r.URL.Query().Get("sort") == "name":
		sor = "name"
	case r.URL.Query().Get("sort") == "date":
		sor = "date"
	default:
		sor = "rating"
	}

	cinemas, err := h.service.GetCinemas(sor)
	if err != nil {
		slog.Error("error while getting result", slog.Any("error", err))
		http.Error(w, "error", http.StatusBadRequest)
		return
	}
	slog.Info("123", slog.Any("123", cinemas))
	byt, err := json.Marshal(cinemas)
	if err != nil {
		slog.Error("error while marshalling cinemas", slog.Any("error", err))
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	w.Write(byt)
}

// @Summary SearchCinema
// @Secutiry ApiKeyAuth
// @Tags cinemas
// @Description search cinema
// @ID search-cinemas
// @Accept json
// @Produce json
// @Success 200 {object} []models.Cinema
// @Failure default {string} error
// @Router /api/cinemas/search [post]
func (h *Handl) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var search models.Search
		byt, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("error with data", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		err = json.Unmarshal(byt, &search)
		if err != nil {
			slog.Error("error while unmarshalling data", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		cinemas, err := h.service.Search(search)
		if err != nil {
			slog.Error("error while searching data", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		byt2, err := json.Marshal(cinemas)
		if err != nil {
			slog.Error("error while marshalling data", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		w.Write(byt2)
	}
}
