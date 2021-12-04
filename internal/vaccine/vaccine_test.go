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

	t.Run("CountryData runs successfully", func(t *testing.T) {
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
}
