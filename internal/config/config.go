package config

/*
   return stuct of a yaml config file
*/

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

/*
	Example of a config file
	server:
		port: ":8888"
*/

//AppConf contains all main structs
type AppConf struct {
	Server ServerConfig `yaml:"server"`
}

//ServerConfig contains the data for the micro servise server
type ServerConfig struct {
	Port string `yaml:"port"`
}

type configOB interface {
	GetAppConfig() AppConf
}

//GetAppConfig reads a spefic file and return the yaml format of it
//return ServerConfig struct yaml format of the config file
func GetAppConfig(path string) (*AppConf, error) {
	var c AppConf

	yamlFile, openfileError := ioutil.ReadFile(path)
	if openfileError != nil {
		return nil, fmt.Errorf("internal_config_GetAppConfig_open_file %w", openfileError)
	}

	err := yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("internal_config_GetAppConfig_yaml_unmarshal %w", err)
	}

	return &c, nil
}
