package handler

import (
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type paymentInfoHandler struct {
	paymentInfoService service.PaymentInfoService
}

type PaymentInfoHandler interface {
	Create(ctx *gin.Context)
	GetAllByExpertiseServiceId(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

func NewPaymentInfoHandler(paymentInfoService service.PaymentInfoService) PaymentInfoHandler {
	return &paymentInfoHandler{
		paymentInfoService: paymentInfoService,
	}
}

func (p paymentInfoHandler) Create(ctx *gin.Context) {
	var paymentInfo web.PaymentInfoRequest
	if err := ctx.ShouldBindJSON(&paymentInfo); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := p.paymentInfoService.Create(&paymentInfo)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

func (p paymentInfoHandler) GetAllByExpertiseServiceId(ctx *gin.Context) {

	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	models, err := p.paymentInfoService.FindAllByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, models)
	ctx.JSON(http.StatusOK, response)
}

func (p *paymentInfoHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	paymentInfo, _ := p.paymentInfoService.FindByID(id)
	ctx.JSON(http.StatusOK, paymentInfo)
}
