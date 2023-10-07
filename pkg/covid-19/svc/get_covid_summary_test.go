package svc

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/cdb"
	"github.com/stretchr/testify/require"
)

func TestCovidService_GetCovidSummaryData(t *testing.T) {
	r := require.New(t)
	mockJson := []byte(`{
		"Data":[
			{
				"ConfirmDate": "2021-05-01",
				"No": null,
				"Age": null,
				"Gender": "หญิง",
				"GenderEn": "Female",
				"Nation": null,
				"NationEn": null,
				"Province": "Samut Songkhram",
				"ProvinceId": 58,
				"District": null,
				"ProvinceEn": "Samut Songkhram",
				"StatQuarantine": 11
			  },
		   {
			  "ConfirmDate":"2021-05-01",
			  "No":null,
			  "Age":25,
			  "Gender":null,
			  "GenderEn":null,
			  "Nation":null,
			  "NationEn":"India",
			  "Province":"Phrae",
			  "ProvinceId":46,
			  "District":null,
			  "ProvinceEn":"Phrae",
			  "StatQuarantine":15
		   },
		   {
				"ConfirmDate":"2021-05-02",
				"No":null,
				"Age":39,
				"Gender":null,
				"GenderEn":null,
				"Nation":null,
				"NationEn":"USA",
				"Province":null,
				"ProvinceId":null,
				"District":null,
				"ProvinceEn":null,
				"StatQuarantine":10
		 	},
		   {
			  "ConfirmDate":"2021-05-01",
			  "No":null,
			  "Age":91,
			  "Gender":"ชาย",
			  "GenderEn":"Male",
			  "Nation":null,
			  "NationEn":"India",
			  "Province":"Kamphaeng Phet",
			  "ProvinceId":14,
			  "District":null,
			  "ProvinceEn":"Kamphaeng Phet",
			  "StatQuarantine":12
		   },
		   {
			  "ConfirmDate":null,
			  "No":null,
			  "Age":92,
			  "Gender":"ชาย",
			  "GenderEn":"Male",
			  "Nation":null,
			  "NationEn":"China",
			  "Province":"Nonthaburi",
			  "ProvinceId":35,
			  "District":null,
			  "ProvinceEn":"Nonthaburi",
			  "StatQuarantine":1
		   }
		]
	 }`)

	err := ioutil.WriteFile("mock.json", mockJson, os.ModePerm)
	r.NoError(err)
	defer func() {
		r.NoError(os.Remove("mock.json"))
	}()

	cdb := cdb.UseData("mock.json")

	c := &CovidService{
		covidData: cdb,
	}
	got, err := c.GetCovidSummaryData()
	r.NoError(err)
	r.Equal(&CovidSummaryResponse{
		AgeGroup: map[string]uint{
			"0-30":  1,
			"31-60": 1,
			"61+":   2,
			"N/A":   1,
		},
		Province: map[string]uint{
			"Kamphaeng Phet":  1,
			"N/A":             1,
			"Nonthaburi":      1,
			"Phrae":           1,
			"Samut Songkhram": 1,
		},
	}, got)
}

func TestCovidService_GetCovidSummaryDataNonExistFile(t *testing.T) {
	r := require.New(t)
	mockJson := []byte(`{
		"Data":[
		   {
			  "ConfirmDate":"2021-05-01",
			  "No":null,
			  "Age":91,
			  "Gender":"ชาย",
			  "GenderEn":"Male",
			  "Nation":null,
			  "NationEn":"India",
			  "Province":"Kamphaeng Phet",
			  "ProvinceId":14,
			  "District":null,
			  "ProvinceEn":"Kamphaeng Phet",
			  "StatQuarantine":12
		   },
		   {
			  "ConfirmDate":null,
			  "No":null,
			  "Age":92,
			  "Gender":"ชาย",
			  "GenderEn":"Male",
			  "Nation":null,
			  "NationEn":"China",
			  "Province":"Nonthaburi",
			  "ProvinceId":35,
			  "District":null,
			  "ProvinceEn":"Nonthaburi",
			  "StatQuarantine":1
		   }
		]
	 }`)

	err := ioutil.WriteFile("mock.json", mockJson, os.ModePerm)
	r.NoError(err)
	defer func() {
		r.NoError(os.Remove("mock.json"))
	}()

	cdb := cdb.UseData("mock2.json")

	c := &CovidService{
		covidData: cdb,
	}
	got, err := c.GetCovidSummaryData()
	r.Error(err)
	r.Nil(got)
}
