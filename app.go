package main

/*
	Author : Iordanis Paschalidis
	Date   : 03/12/2021

*/

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	config "github.com/junkd0g/covid-vaccine/internal/config"
	"github.com/junkd0g/covid-vaccine/internal/controller"
	"github.com/junkd0g/covid-vaccine/internal/data"
	"github.com/junkd0g/covid-vaccine/internal/vaccine"
)

//Service object that contains the Port and Router of the application
type Service struct {
	Port   string
	Router *mux.Router
}

/*
   Running the service in port 8888 (getting the value from ./assets/config/production.json )

       Endpoints:
		GET:
			api/data/{country}
		POST:
*/
func (s Service) run() {

	configData, err := config.GetAppConfig("./config.yaml")
	if err != nil {
		panic(fmt.Errorf("creating_config %w", err))

	}

	dataClient, err := data.NewReadDataClient("./scripts/get_data/data_out.json")
	if err != nil {
		panic(fmt.Errorf("data_client %w", err))
	}

	vaccineClient, err := vaccine.NewClient(dataClient)
	if err != nil {
		panic(fmt.Errorf("vaccine_client %w", err))
	}

	s.Port = configData.Server.Port
	country, err := controller.NewCountry(vaccineClient)
	if err != nil {
		panic(fmt.Errorf("creating_mail_controller %w", err))
	}
	s.Router.HandleFunc("/api/data/{country}", country.Middleware).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*", "Authorization"},
	})

	handler := c.Handler(s.Router)

	fmt.Println("server running at port " + s.Port)
	err = http.ListenAndServe(s.Port, handler)
	if err != nil {
		panic(fmt.Errorf("listener_and_serve %w", err))
	}
}

func main() {
	service := Service{Router: mux.NewRouter().StrictSlash(true)}
	service.run()
}
