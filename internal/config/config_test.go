package config_test

import (
	"testing"

	"github.com/junkd0g/covid-vaccine/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_GetAppConfig(t *testing.T) {

	t.Run("create config successfully ", func(t *testing.T) {
		configData, configDataError := config.GetAppConfig("./test_config.yaml")

		assert.Nil(t, configDataError)
		assert.NotNil(t, configData)
		assert.Equal(t, ":8888", configData.Server.Port, "wrong port value")
	})

	t.Run("create config error when trying to open file that doesnt exist ", func(t *testing.T) {
		expectedErrorPartOfMessage := "internal_config_GetAppConfig_open_file"
		configData, configDataError := config.GetAppConfig("some random file")

		assert.Nil(t, configData)
		assert.NotNil(t, configDataError)
		assert.Contains(t, configDataError.Error(), expectedErrorPartOfMessage)
	})

	t.Run("create config error when trying unmarshal yaml file", func(t *testing.T) {
		expectedErrorPartOfMessage := "internal_config_GetAppConfig_yaml_unmarshal"
		configData, configDataError := config.GetAppConfig("./test_config2.yaml")

		assert.Nil(t, configData)
		assert.NotNil(t, configDataError)
		assert.Contains(t, configDataError.Error(), expectedErrorPartOfMessage)
	})
}
