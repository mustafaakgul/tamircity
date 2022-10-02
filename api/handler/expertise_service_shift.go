package handler

import (
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type expertiseServiceShiftHandler struct {
	expertiseServiceShiftService service.ExpertiseServiceShiftService
}

type ExpertiseServiceShiftHandler interface {
	GetByExpertiseServiceId(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewExpertiseServiceShiftHandler(expertiseServiceShiftService service.ExpertiseServiceShiftService) ExpertiseServiceShiftHandler {
	return &expertiseServiceShiftHandler{expertiseServiceShiftService: expertiseServiceShiftService}
}
func (t *expertiseServiceShiftHandler) GetByExpertiseServiceId(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	//var expertiseServices []web.ExpertiseServiceSearchResponse
	expertiseServices, err := t.expertiseServiceShiftService.FindByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, expertiseServices)
	ctx.JSON(http.StatusOK, response)
}

func (t *expertiseServiceShiftHandler) Create(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	var expertiseServiceShiftRequest []web.ExpertiseServiceShiftRequest
	if err := ctx.ShouldBindJSON(&expertiseServiceShiftRequest); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	t.expertiseServiceShiftService.Create(expertiseServiceId, expertiseServiceShiftRequest)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, nil)
	ctx.JSON(http.StatusOK, response)

}
