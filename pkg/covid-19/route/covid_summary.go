package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markpassawat/lmwn-assignment/pkg/covid-19/svc"
)

type CovidHandler struct {
	covidService svc.CovidService
}

func NewHandler(covidService svc.CovidService) *CovidHandler {
	return &CovidHandler{
		covidService: covidService,
	}
}

func (r *CovidHandler) GetSummary() gin.HandlerFunc {
	return func(c *gin.Context) {
		dataSummary, err := r.covidService.GetCovidSummaryData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, dataSummary)
	}
}
