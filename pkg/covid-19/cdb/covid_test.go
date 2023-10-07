package cdb

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestUseData(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name     string
		filePath string
		want     *CovidData
		wantErr  bool
	}{
		{
			name:     "valid file path",
			filePath: "pkg/covid-19/cdb/data.json",
			want: &CovidData{
				path: "pkg/covid-19/cdb/data.json",
			},
			wantErr: false,
		},
		{
			name:     "empty file path",
			filePath: "",
			want: &CovidData{
				path: "",
			},
			wantErr: false,
		},
		{
			name:     "invalid file path",
			filePath: "pkg/covid-19/cdb/data.json",
			want: &CovidData{
				path: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UseData(tt.filePath)
			if !tt.wantErr {
				r.Equal(tt.want, got)
			}
		})
	}
}

func TestGetCases(t *testing.T) {
	r := require.New(t)

	mockJSONFile, err := os.CreateTemp("", "covid_cases_*.json")
	logrus.Info("mockJSONFile.Name(), ", mockJSONFile.Name())
	require.NoError(t, err)
	defer mockJSONFile.Close()
	defer os.Remove(mockJSONFile.Name())

	var covidData model.CasesJSON
	err = json.NewEncoder(mockJSONFile).Encode(covidData)
	r.NoError(err)

	covidDataInstance := &CovidData{path: mockJSONFile.Name()}
	cases, err := covidDataInstance.GetCases()
	r.NoError(err)

	r.Equal(covidData.Data, cases)
}

func TestGetCasesFromNonExistFile(t *testing.T) {
	r := require.New(t)
	covidDataInstance := &CovidData{path: "non-exist-file.json"}
	cases, err := covidDataInstance.GetCases()
	r.Error(err)
	r.Nil(cases)
}

func TestGetCasesFromInvalidJSONFile(t *testing.T) {
	r := require.New(t)

	mockJSONFile, err := os.CreateTemp("", "covid_cases_*.json")
	require.NoError(t, err)
	defer mockJSONFile.Close()
	defer os.Remove(mockJSONFile.Name())

	_, err = mockJSONFile.WriteString("{invalid json}")
	r.NoError(err)

	covidDataInstance := &CovidData{path: mockJSONFile.Name()}
	cases, err := covidDataInstance.GetCases()
	r.Error(err)
	r.Nil(cases)
}
