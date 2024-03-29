package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/models"
)

// @Summary InsertCinema
// @Security ApiKeyAuth
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
			w.Header().Set("Content-Type", "application/json")
			w.Write(byt2)
			slog.Info("cinema inserted with id", slog.Int("id", id))
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
// @Security ApiKeyAuth
// @Tags cinemas
// @Description update cinema
// @ID update-cinemas
// @Accept json
// @Produce json
// @Success 200 {string} id
// @Failure default {string} error
// @Router /api/cinemas/{id} [put]
func (h *Handl) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		return
	} else {
		var cinema models.CinemaUpdate

		byt, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(byt, &cinema)
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
		byt2, _ := json.Marshal(map[string]string{"id": id})
		w.Header().Set("Content-Type", "application/json")
		w.Write(byt2)
		slog.Info("cinema updated with id", slog.String("id", id))
	}
}

// @Summary DeleteCinemas
// @Security ApiKeyAuth
// @Tags cinemas
// @Description delete cinema
// @ID get-cinemas
// @Produce json
// @Success 200 {string} id
// @Failure default {string} error
// @Router /api/cinemas/{id} [delete]
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
		byt, _ := json.Marshal(map[string]string{"id": id})
		w.Header().Set("Content-Type", "application/json")
		w.Write(byt)
		slog.Info("cinema deleted with id", slog.String("id", id))
	}
}

// @Summary GetAllCinemas
// @Security ApiKeyAuth
// @Tags cinemas
// @Description get all cinemas
// @ID get-cinemas
// @Produce json
// @Success 200 {object} []models.Cinema
// @Failure default {string} error
// @Router /api/cinemas/{sort} [get]
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
	byt, _ := json.Marshal(cinemas)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
	slog.Info("cinemas gotten", slog.Any("cinemas", cinemas))
}

// @Summary SearchCinema
// @Security ApiKeyAuth
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
			return
		}
		err = json.Unmarshal(byt, &search)
		if err != nil {
			slog.Error("error while unmarshalling data", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		cinemas, err := h.service.Search(search)
		if err != nil {
			slog.Error("error while searching data", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		byt2, _ := json.Marshal(cinemas)
		w.Header().Set("Content-Type", "application/json")
		w.Write(byt2)
		slog.Info("cinemas gotten", slog.Any("cinemas", cinemas))
	}
}
