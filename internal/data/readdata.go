package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ReadDataClient struct {
	Path string
}

// NewReadDataClient creates and return new ReadDataClient object
func NewReadDataClient(path string) (ReadDataClient, error) {
	if path == "" {
		return ReadDataClient{}, fmt.Errorf("data_NewReadDataClient_empty_path")
	}

	return ReadDataClient{
		Path: path,
	}, nil
}

//ReadData opens json file and unmarshal the data to Countries struct
func (c ReadDataClient) ReadData() (Countries, error) {
	jsonFile, err := os.Open(c.Path)
	if err != nil {
		return Countries{}, fmt.Errorf("data_ReadData_open_file %w", err)
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Countries{}, fmt.Errorf("data_ReadData_read_all %w", err)
	}

	var countries Countries

	err = json.Unmarshal(byteValue, &countries)
	if err != nil {
		return Countries{}, fmt.Errorf("data_ReadData_unmarshal %w", err)
	}
	return countries, nil
}
