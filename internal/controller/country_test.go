package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	vmocks "github.com/junkd0g/covid-vaccine/internal/mocks/controller"
	"github.com/junkd0g/covid-vaccine/internal/vaccine"

	"github.com/junkd0g/covid-vaccine/internal/controller"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	t     *testing.T
	rd    *vmocks.MockVaccine
	rjson *vmocks.MockJSON
}

func getMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	return &mocks{
		t:     t,
		rd:    vmocks.NewMockVaccine(ctrl),
		rjson: vmocks.NewMockJSON(ctrl),
	}
}

func Test_NewCountry(t *testing.T) {
	m := getMocks(t)
	t.Run("Create object successfully", func(t *testing.T) {
		client, err := controller.NewCountry(m.rd, m.rjson)
		assert.Nil(t, err)
		assert.NotNil(t, client)
	})

	t.Run("No country data error", func(t *testing.T) {
		expectedError := "controller_NewCountry_empty_countryData"

		_, err := controller.NewCountry(nil, nil)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)

	})

	t.Run("No json error", func(t *testing.T) {
		expectedError := "controller_NewCountry_empty_json"

		_, err := controller.NewCountry(m.rd, nil)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)

	})
}

func Test_Middleware(t *testing.T) {
	m := getMocks(t)
	t.Run("Return a 500 error when country Data returns an erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/data/{country}", nil)
		res := httptest.NewRecorder()
		m.rd.
			EXPECT().
			CountryData(gomock.Any()).
			Return(vaccine.CountryDataResponse{}, fmt.Errorf("some error"))
		client, err := controller.NewCountry(m.rd, m.rjson)
		assert.Nil(t, err)
		assert.NotNil(t, client)
		client.Middleware(res, req)
		assert.Equal(t, 500, res.Code)
		expected := `{"message":"some error","status":500}`
		assert.Equal(t, expected, res.Body.String())

	})

	t.Run("Return a 500 error when country json.Marshal an error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/data/{country}", nil)
		res := httptest.NewRecorder()
		m.rd.
			EXPECT().
			CountryData(gomock.Any()).
			Return(vaccine.CountryDataResponse{}, nil)

		m.rjson.
			EXPECT().
			Marshal(vaccine.CountryDataResponse{}).
			Return(nil, fmt.Errorf("some error"))

		client, err := controller.NewCountry(m.rd, m.rjson)
		assert.Nil(t, err)
		assert.NotNil(t, client)
		client.Middleware(res, req)
		assert.Equal(t, 500, res.Code)
		expected := `{"message":"some error","status":500}`
		assert.Equal(t, expected, res.Body.String())
	})

	t.Run("Runs successfully", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/data/greece", nil)
		res := httptest.NewRecorder()
		expectCountries := vaccine.CountryDataResponse{
			TotalVaccinations:               88,
			PeopleFullyVaccinations:         432,
			TotalBooster:                    122,
			TotalVaccinationsPerHundred:     "232",
			PeopleVaccinatedPerHundred:      "83",
			PeopleFullyVaccinatedPerHundred: "322",
			TotalBoostersPerHundred:         "17",
		}

		m.rd.
			EXPECT().
			CountryData(gomock.Any()).
			Return(expectCountries, nil).AnyTimes()

		jsonData, _ := json.Marshal(expectCountries)

		m.rjson.
			EXPECT().
			Marshal(gomock.Any()).
			Return(jsonData, nil).AnyTimes()

		client, err := controller.NewCountry(m.rd, m.rjson)
		assert.Nil(t, err)
		assert.NotNil(t, client)
		client.Middleware(res, req)
		assert.Equal(t, 200, res.Code)
		expected := `{"total_vaccinations":88,"fully_vaccinations":432,"booster_vaccinations":122,"total_vaccinations_per_hundred":"232","people_vaccinated_per_hundred":"83","people_fully_vaccinated_per_hundred":"322","total_boosters_per_hundred":"17"}`
		assert.Equal(t, expected, res.Body.String())
	})
}
