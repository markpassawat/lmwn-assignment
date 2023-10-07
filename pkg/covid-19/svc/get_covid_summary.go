package svc

import (
	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/cdb"
	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/model"
)

type CovidDataStore interface {
	GetCases() ([]model.Case, error)
}
type CovidService struct {
	covidData CovidDataStore
}

func NewService(covidData *cdb.CovidData) CovidService {
	return CovidService{
		covidData: covidData,
	}
}

var (
	below30       = "0-30"
	between31to60 = "31-60"
	above60       = "61+"
	notAvailable  = "N/A"
)

type CovidSummaryResponse struct {
	AgeGroup map[string]uint `json:"AgeGroup"`
	Province map[string]uint `json:"Province"`
}

func (c *CovidService) GetCovidSummaryData() (*CovidSummaryResponse, error) {
	covidCases, err := c.covidData.GetCases()
	if err != nil {
		return nil, err
	}

	mapProvince := make(map[string]uint)
	mapAgeGroup := map[string]uint{
		below30:       0,
		between31to60: 0,
		above60:       0,
	}

	for _, covidCase := range covidCases {
		province := notAvailable
		if covidCase.Province != nil {
			province = *covidCase.Province
		}
		mapProvince[province]++

		ageRange := notAvailable
		if covidCase.Age != nil {
			if *covidCase.Age >= 0 && *covidCase.Age <= 30 {
				ageRange = below30
			} else if *covidCase.Age >= 31 && *covidCase.Age <= 60 {
				ageRange = between31to60
			} else if *covidCase.Age > 60 {
				ageRange = above60
			}
		}
		mapAgeGroup[ageRange]++
	}

	return &CovidSummaryResponse{
		AgeGroup: mapAgeGroup,
		Province: mapProvince,
	}, nil
}
