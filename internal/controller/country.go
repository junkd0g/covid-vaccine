package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/junkd0g/covid-vaccine/internal/vaccine"
	merror "github.com/junkd0g/neji"
)

type EmailSendResponse struct {
	Success bool `json:"success"`
}

type emailMiddlewareClient struct{}

// NewCountry creates new object of mailjet's client
func NewCountry() (*emailMiddlewareClient, error) {
	return &emailMiddlewareClient{}, nil
}

//SendEmailMiddleware middleware for /api/email/send
func (c emailMiddlewareClient) Middleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	jsonBody, status := c.perform(vars["country"])

	w.WriteHeader(status)
	w.Write(jsonBody)
}

func (c emailMiddlewareClient) perform(country string) ([]byte, int) {

	vaccineClient := vaccine.Client{}

	data, err := vaccineClient.CountryData(country)
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
