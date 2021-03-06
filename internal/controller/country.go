package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/junkd0g/covid-vaccine/internal/vaccine"
	merror "github.com/junkd0g/neji"
)

type Vaccine interface {
	CountryData(country string) (vaccine.CountryDataResponse, error)
}

type EmailSendResponse struct {
	Success bool `json:"success"`
}

type countryMiddlewareClient struct {
	vaccineClient Vaccine
}

// NewCountry creates new object of mailjet's client
func NewCountry(countryData Vaccine) (countryMiddlewareClient, error) {
	if countryData == nil {
		return countryMiddlewareClient{}, fmt.Errorf("controller_NewCountry_empty_countryData")
	}

	return countryMiddlewareClient{
		vaccineClient: countryData,
	}, nil
}

//Middleware middleware for /api/data/{country}
func (c countryMiddlewareClient) Middleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	jsonBody, status := c.perform(vars["country"])

	w.WriteHeader(status)
	w.Write(jsonBody)
}

func (c countryMiddlewareClient) perform(country string) ([]byte, int) {

	data, err := c.vaccineClient.CountryData(country)
	if err != nil {
		errorJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, err)
		return errorJSONBody, 500
	}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		errorJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, err)
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
