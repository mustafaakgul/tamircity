package handler

import (
	tech_service2 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	"github.com/anthophora/tamircity/pkg/service/tech_service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type technicalServiceShiftHandler struct {
	technicalServiceShiftService tech_service.TechnicalServiceShiftService
}

type TechnicalServiceShiftHandler interface {
	GetByTechnicalServiceId(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewTechnicalServiceShiftHandler(technicalServiceShiftService tech_service.TechnicalServiceShiftService) TechnicalServiceShiftHandler {
	return &technicalServiceShiftHandler{technicalServiceShiftService: technicalServiceShiftService}
}
func (t *technicalServiceShiftHandler) GetByTechnicalServiceId(ctx *gin.Context) {
	technicalServiceId, err := strconv.Atoi(ctx.Query("technical_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	//var technicalServices []web.TechnicalServiceSearchResponse
	technicalServices, err := t.technicalServiceShiftService.FindByTechnicalServiceId(technicalServiceId)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, technicalServices)
	ctx.JSON(http.StatusOK, response)
}

func (t *technicalServiceShiftHandler) Create(ctx *gin.Context) {
	technicalServiceId, err := strconv.Atoi(ctx.Query("technical_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	var technicalServiceShiftRequest []tech_service2.TechnicalServiceShiftRequest
	if err := ctx.ShouldBindJSON(&technicalServiceShiftRequest); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	t.technicalServiceShiftService.Create(technicalServiceId, technicalServiceShiftRequest)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, nil)
	ctx.JSON(http.StatusOK, response)

}
