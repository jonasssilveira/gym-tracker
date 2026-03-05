package set

import (
	"gym-tracker/app/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SetsHandler struct {
	service Service
}

func NewSetHandler(service Service) *SetsHandler {
	return &SetsHandler{
		service: service,
	}
}

func (handler *SetsHandler) GetAllSet(c *gin.Context) {
	serieID, err := strconv.ParseUint(c.Param("serieID"), 10, 64)
	if err != nil {
		api.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
	}
	set, err := handler.service.GetALlSetsFromSerie(serieID)
	if err != nil {
		api.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error())
	}
	setDTO := make([]SetDTO, len(set))
	for i := range set {
		setDTO[i] = FromEntity(set[i])
	}

	api.OKResponse(c.Writer, setDTO)
}

func (handler *SetsHandler) CreateSet(c *gin.Context) {
	var setDTO SetDTO

	if err := c.ShouldBindJSON(&setDTO); err != nil {
		api.ErrorResponse(c.Writer, http.StatusBadRequest, err.Error())
		return
	}
	if invalid := handler.service.AddSet(setDTO); invalid != nil {
		api.ErrorResponse(c.Writer, http.StatusBadRequest, invalid.Error())
		return
	}
	api.OKResponse(c.Writer, setDTO)
}
