package controller

import "fmt"

type compareMiddlewareClient struct {
	vaccineClient Vaccine
}

// NewCompare creates new object of compareMiddlewareClient
func NewCompare(countryData Vaccine) (compareMiddlewareClient, error) {
	if countryData == nil {
		return compareMiddlewareClient{}, fmt.Errorf("controller_NewCompare_empty_countryData")
	}

	return compareMiddlewareClient{
		vaccineClient: countryData,
	}, nil
}
