package vaccine

type Countries struct {
	Data []struct {
		Name                            string `json:"name"`
		PeopleVaccinations              []int  `json:"people_vaccinations"`
		TotalBoostersPerHundred         string `json:"total_boosters_per_hundred"`
		TotalVaccinations               []int  `json:"total_vaccinations"`
		PeopleVaccinatedPerHundred      string `json:"people_vaccinated_per_hundred"`
		TotalVaccinationsPerHundred     string `json:"total_vaccinations_per_hundred"`
		TotalBooster                    []int  `json:"total_booster"`
		PeopleFullyVaccinations         []int  `json:"people_fully_vaccinations"`
		PeopleFullyVaccinatedPerHundred string `json:"people_fully_vaccinated_per_hundred"`
	} `json:"data"`
}
