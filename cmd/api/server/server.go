package server

import (
	"github.com/gin-gonic/gin"
	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/cdb"
	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/route"
	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/svc"
)

func New() *gin.Engine {
	r := gin.Default()

	covidData := cdb.UseData("pkg/covid-19/cdb/data.json")
	covidService := svc.NewService(covidData)
	covidRoute := route.NewHandler(covidService)
	covidHandler := r.Group("/covid/")
	{
		covidHandler.GET("/summary", covidRoute.GetSummary())
	}

	return r
}
