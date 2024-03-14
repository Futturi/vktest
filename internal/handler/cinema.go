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
		}
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
