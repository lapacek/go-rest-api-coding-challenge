package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lapacek/simple-api-example/internal/model"
)

func AllBookings(model *model.Model, w http.ResponseWriter) {

	bookings, err := model.AllBookings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(*bookings)
}
