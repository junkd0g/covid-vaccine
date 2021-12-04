package vaccine

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Client struct{}

type Response struct {
	TotalVaccinations       int `json:"total_vaccinations"`
	PeopleFullyVaccinations int `json:"fully_vaccinations"`
	TotalBooster            int `json:"booster_vaccinations"`

	TotalVaccinationsPerHundred     string `json:"total_vaccinations_per_hundred"`
	PeopleVaccinatedPerHundred      string `json:"people_vaccinated_per_hundred"`
	PeopleFullyVaccinatedPerHundred string `json:"people_fully_vaccinated_per_hundred"`
	TotalBoostersPerHundred         string `json:"total_boosters_per_hundred"`
}

func (c Client) readData() (Countries, error) {
	jsonFile, err := os.Open("./scripts/get_data/data_out.json")
	if err != nil {
		return Countries{}, err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Countries{}, err
	}
	var countries Countries

	json.Unmarshal(byteValue, &countries)
	return countries, nil
}

func (c Client) CountryData(country string) (Response, error) {
	var resp Response

	data, err := c.readData()
	if err != nil {
		return Response{}, err
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
