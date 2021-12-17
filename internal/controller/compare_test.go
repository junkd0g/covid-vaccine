package controller_test

import (
	"testing"

	"github.com/junkd0g/covid-vaccine/internal/controller"
	"github.com/stretchr/testify/assert"
)

func Test_NewCompare(t *testing.T) {
	m := getMocks(t)
	t.Run("Create object successfully", func(t *testing.T) {
		client, err := controller.NewCompare(m.rd)
		assert.Nil(t, err)
		assert.NotNil(t, client)
	})

	t.Run("No country data error", func(t *testing.T) {
		expectedError := "controller_NewCompare_empty_countryData"

		_, err := controller.NewCompare(nil)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), expectedError)

	})

}
