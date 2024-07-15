package handler

import (
	"encoding/json"
	"net/http"

	"github.com/crnvl96/temperature-calculator/internal/address"
	"github.com/crnvl96/temperature-calculator/internal/weather"
	"github.com/paemuri/brdoc"
)

func NewWeatherHandler(r *http.ServeMux) {
	r.HandleFunc("GET /calculate",
		func(w http.ResponseWriter, r *http.Request) {
			zipcode := r.URL.Query().Get("zipcode")

			if brdoc.IsCEP(zipcode) == false {
				http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
				return
			}

			city, err := address.GetAddressCity(zipcode)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if city == "" {
				http.Error(w, "can not find zipcode", http.StatusNotFound)
				return
			}

			temp, err := weather.GetCurrentTemperatureByCity(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(temp)
		},
	)
}
