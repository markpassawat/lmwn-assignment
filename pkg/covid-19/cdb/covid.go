package cdb

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/model"
)

type CovidData struct {
	path string
}

func UseData(filePath string) *CovidData {
	return &CovidData{path: filePath}
}

func (c *CovidData) GetCases() ([]model.Case, error) {
	var covidCasesJSON model.CasesJSON
	jsonFile, err := os.Open(c.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read covid cases from json: %w", err)
	}
	defer jsonFile.Close()

	if err := json.NewDecoder(jsonFile).Decode(&covidCasesJSON); err != nil {
		return nil, fmt.Errorf("failed to decode covid cases from json: %w", err)
	}
	return covidCasesJSON.Data, nil
}
