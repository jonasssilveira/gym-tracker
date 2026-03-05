package series

import (
	"gym-tracker/app/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeriesHandler struct {
	service Service
}

func NewSeriesHandler(service Service) *SeriesHandler {
	return &SeriesHandler{service: service}
}

func (handler *SeriesHandler) GetAllSeries(c *gin.Context) {
	series := handler.service.GetALlSeries()

	seriesDTO := make([]SeriesDTO, len(series))
	for i := range series {
		seriesDTO[i] = FromEntity(series[i])
	}

	api.OKResponse(c.Writer, seriesDTO)
}

func (handler *SeriesHandler) CreateSeries(c *gin.Context) {
	var seriesDTO SeriesDTO

	if err := c.ShouldBindJSON(&seriesDTO); err != nil {
		api.ErrorResponse(c.Writer, http.StatusBadRequest, err.Error())
		return
	}
	if invalid := seriesDTO.Validate(); invalid != nil {
		api.ErrorResponse(c.Writer, http.StatusBadRequest, invalid.Error())
		return
	}
	_, err := handler.service.CreateSeries(seriesDTO.ToEntity())
	if err != nil {
		api.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
	}
	api.OKResponse(c.Writer, seriesDTO)
}

func (handler *SeriesHandler) FinalizeSerie(c *gin.Context) {
	var seriesDTO SeriesDTO
	if err := c.ShouldBindJSON(&seriesDTO); err != nil {
		api.ErrorResponse(c.Writer, http.StatusBadRequest, err.Error())
	}
	//err := handler.service.FinalizeSerie(strconv.Itoa(int(seriesDTO.ID)))
	//if err != nil {
	//	api.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
	//	return
	//}
	api.OKResponse(c.Writer, seriesDTO)
}
