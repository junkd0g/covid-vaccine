package vaccine

import (
	"fmt"

	"github.com/junkd0g/covid-vaccine/internal/data"
)

type RD interface {
	ReadData() (data.Countries, error)
}

type client struct {
	rd RD
}

type CountryDataResponse struct {
	TotalVaccinations               int    `json:"total_vaccinations"`
	PeopleFullyVaccinations         int    `json:"fully_vaccinations"`
	TotalBooster                    int    `json:"booster_vaccinations"`
	TotalVaccinationsPerHundred     string `json:"total_vaccinations_per_hundred"`
	PeopleVaccinatedPerHundred      string `json:"people_vaccinated_per_hundred"`
	PeopleFullyVaccinatedPerHundred string `json:"people_fully_vaccinated_per_hundred"`
	TotalBoostersPerHundred         string `json:"total_boosters_per_hundred"`
}

// NewClient creates and return new client object
func NewClient(readData RD) (client, error) {
	if readData == nil {
		return client{}, fmt.Errorf("vaccine_NewClient_empty_readData")
	}
	return client{
		rd: readData,
	}, nil
}

// CountryData return one couty's data
func (c client) CountryData(country string) (CountryDataResponse, error) {

	var resp CountryDataResponse
	data, err := c.rd.ReadData()
	if err != nil {
		return CountryDataResponse{}, err
	}

	for _, x := range data.Data {
		if x.Name == country {
			sl := x.TotalVaccinations
			resp.TotalVaccinations = sl[len(sl)-1]
			sl1 := x.TotalBooster
			resp.TotalBooster = sl1[len(sl1)-1]
			sl2 := x.PeopleFullyVaccinations
			resp.PeopleFullyVaccinations = sl2[len(sl2)-1]
			resp.TotalVaccinationsPerHundred = x.TotalVaccinationsPerHundred
			resp.PeopleVaccinatedPerHundred = x.PeopleVaccinatedPerHundred
			resp.PeopleFullyVaccinatedPerHundred = x.PeopleFullyVaccinatedPerHundred
			resp.TotalBoostersPerHundred = x.TotalBoostersPerHundred
		}
	}

	return resp, nil
}
