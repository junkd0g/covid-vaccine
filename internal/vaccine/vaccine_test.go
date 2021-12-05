package vaccine_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/junkd0g/covid-vaccine/internal/data"
	vmocks "github.com/junkd0g/covid-vaccine/internal/mocks/vaccine"
	"github.com/junkd0g/covid-vaccine/internal/vaccine"

	"github.com/stretchr/testify/assert"
)

type mocks struct {
	t  *testing.T
	rd *vmocks.MockRD
}

func getMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	return &mocks{
		t:  t,
		rd: vmocks.NewMockRD(ctrl),
	}
}

func Test_NewClient(t *testing.T) {
	mocks := getMocks(t)
	t.Run("Create object successfully", func(t *testing.T) {
		client, err := vaccine.NewClient(mocks.rd)
		assert.Nil(t, err)
		assert.NotNil(t, client)
	})

	t.Run("No ReadData error", func(t *testing.T) {
		expectedError := "vaccine_NewClient_empty_readData"
		_, err := vaccine.NewClient(nil)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)
	})
}

func Test_CountryData(t *testing.T) {
	m := getMocks(t)
	t.Run("ReadData returns error", func(t *testing.T) {
		expectedError := "data_ReadData_open_file"
		m.rd.
			EXPECT().
			ReadData().
			Return(data.Countries{}, fmt.Errorf("%s", expectedError)).
			AnyTimes()

		client, err := vaccine.NewClient(m.rd)
		assert.Nil(t, err)

		_, err = client.CountryData("some country")
		assert.NotNil(t, err)
		assert.Contains(t, expectedError, err.Error())
	})

	t.Run("CountryData runs successfully but empty response", func(t *testing.T) {
		m := getMocks(t)
		var dd []data.Country

		el := data.Country{
			Name:                            "Greece",
			PeopleVaccinations:              []int{3, 4, 5, 6, 7, 88},
			TotalBoostersPerHundred:         "17",
			TotalVaccinations:               []int{3, 4, 5, 6, 7, 88},
			PeopleVaccinatedPerHundred:      "83",
			TotalVaccinationsPerHundred:     "232",
			TotalBooster:                    []int{3, 4, 5, 6, 7, 88},
			PeopleFullyVaccinations:         []int{3, 4, 5, 6, 7, 88},
			PeopleFullyVaccinatedPerHundred: "322",
		}

		dd = append(dd, el)

		m.rd.
			EXPECT().
			ReadData().
			Return(data.Countries{
				Data: dd,
			}, nil)

		client, err := vaccine.NewClient(m.rd)
		assert.Nil(t, err)

		resp, err := client.CountryData("some country")
		assert.Nil(t, err)
		assert.Equal(t, vaccine.CountryDataResponse{}, resp)
	})

	t.Run("CountryData runs successfully with a response", func(t *testing.T) {
		m := getMocks(t)
		var dd []data.Country

		el := data.Country{
			Name:                            "Greece",
			PeopleVaccinations:              []int{3, 4, 5, 6, 7, 324},
			TotalBoostersPerHundred:         "17",
			TotalVaccinations:               []int{3, 4, 5, 6, 7, 88},
			PeopleVaccinatedPerHundred:      "83",
			TotalVaccinationsPerHundred:     "232",
			TotalBooster:                    []int{3, 4, 5, 6, 7, 122},
			PeopleFullyVaccinations:         []int{3, 4, 5, 6, 7, 432},
			PeopleFullyVaccinatedPerHundred: "322",
		}

		el2 := data.Country{
			Name:                            "Portugal",
			PeopleVaccinations:              []int{3, 4, 5, 6, 7, 1324},
			TotalBoostersPerHundred:         "117",
			TotalVaccinations:               []int{3, 4, 5, 6, 7, 188},
			PeopleVaccinatedPerHundred:      "183",
			TotalVaccinationsPerHundred:     "232",
			TotalBooster:                    []int{3, 4, 5, 6, 7, 1122},
			PeopleFullyVaccinations:         []int{3, 4, 5, 6, 7, 1432},
			PeopleFullyVaccinatedPerHundred: "1322",
		}

		dd = append(dd, el)
		dd = append(dd, el2)

		m.rd.
			EXPECT().
			ReadData().
			Return(data.Countries{
				Data: dd,
			}, nil)

		client, err := vaccine.NewClient(m.rd)
		assert.Nil(t, err)

		resp, err := client.CountryData("Greece")
		assert.Nil(t, err)
		extectedResponse := vaccine.CountryDataResponse{
			TotalVaccinations:               88,
			PeopleFullyVaccinations:         432,
			TotalBooster:                    122,
			TotalVaccinationsPerHundred:     "232",
			PeopleVaccinatedPerHundred:      "83",
			PeopleFullyVaccinatedPerHundred: "322",
			TotalBoostersPerHundred:         "17",
		}
		assert.Equal(t, extectedResponse, resp)
	})
}
