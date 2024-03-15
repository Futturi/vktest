package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Futturi/vktest/internal/models"
)

func (h *Handl) InsertFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !h.GetPrivileage(r) {
			http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
		} else {
			byt, err := io.ReadAll(r.Body)
			if err != nil {
				slog.Error("error while inserting actor", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
			}
			var cinema models.Cinema
			err = json.Unmarshal(byt, &cinema)
			if err != nil {
				slog.Error("error while unmarshalling body", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
			}
			slog.Info(cinema.Data)
			id, err := h.service.InsertCinema(cinema)
			if err != nil {
				slog.Error("error while inserting body", slog.Any("error", err))
				http.Error(w, "error", http.StatusBadRequest)
			}
			byt2, err := json.Marshal(map[string]int{"id": id})
			if err != nil {
				slog.Error("error while marshalling result", slog.Any("error", err))
				http.Error(w, "error", http.StatusInternalServerError)
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

func (h *Handl) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
	} else {
		var cinema models.CinemaUpdate

		byt, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("error while inserting actor", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		json.Unmarshal(byt, &cinema)

		id := r.URL.Query().Get("id")
		err = h.service.UpdateFilm(id, cinema)
		if err != nil {
			slog.Error("error while inserting body", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		byt2, err := json.Marshal(map[string]string{"id": id})
		if err != nil {
			slog.Error("error while marshalling body", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		w.Write(byt2)
	}
}

func (h *Handl) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	if !h.GetPrivileage(r) {
		http.Error(w, "u havent privileage for doing this", http.StatusBadRequest)
	} else {
		id := r.URL.Query().Get("id")
		err := h.service.DeleteFilm(id)
		if err != nil {
			slog.Error("error while deleting cinema", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		byt, err := json.Marshal(map[string]string{"id": id})
		if err != nil {
			slog.Error("error while marshalling result", slog.Any("error", err))
			http.Error(w, "error", http.StatusBadRequest)
		}
		w.Write(byt)
	}
}

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
	}
	slog.Info("123", slog.Any("123", cinemas))
	byt, err := json.Marshal(cinemas)
	if err != nil {
		slog.Error("error while marshalling cinemas", slog.Any("error", err))
		http.Error(w, "error", http.StatusInternalServerError)
	}
	w.Write(byt)
}

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
