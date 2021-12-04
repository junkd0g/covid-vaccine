package data_test

import (
	"testing"

	data "github.com/junkd0g/covid-vaccine/internal/data"
	"github.com/stretchr/testify/assert"
)

func Test_NewReadDataClient(t *testing.T) {
	t.Run("Create object successfully", func(t *testing.T) {
		path := "some path"
		client, err := data.NewReadDataClient(path)
		assert.Nil(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, client.Path, path)
	})

	t.Run("No path error", func(t *testing.T) {
		expectedError := "data_NewReadDataClient_empty_path"
		path := ""
		_, err := data.NewReadDataClient(path)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)
	})
}

func Test_ReadData(t *testing.T) {
	t.Run("Error reading file", func(t *testing.T) {
		expectedError := "data_ReadData_open_file"
		path := "some path"
		client, _ := data.NewReadDataClient(path)
		_, err := client.ReadData()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)
	})

	t.Run("Error unmarshal json", func(t *testing.T) {
		expectedError := "data_ReadData_unmarshal"
		path := "test_data_broken.json"
		client, _ := data.NewReadDataClient(path)
		_, err := client.ReadData()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)
	})

	t.Run("Successfully read data", func(t *testing.T) {
		path := "test_data.json"
		client, _ := data.NewReadDataClient(path)
		countries, err := client.ReadData()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(countries.Data))
		assert.Equal(t, "some location", countries.Data[0].Name)
	})
}
