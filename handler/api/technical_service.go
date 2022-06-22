package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"net/http"
)

type TechnicalServiceHandler interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type technicalServiceHandler struct {
	technicalServiceService service.TechnicalServiceService
}

func NewTechnicalServiceHandler(technicalServiceService service.TechnicalServiceService) TechnicalServiceHandler {
	return &technicalServiceHandler{
		technicalServiceService: technicalServiceService,
	}
}

func (t *technicalServiceHandler) Get(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *technicalServiceHandler) Create(ctx *gin.Context) {
	var technicalService db.TechnicalServiceRequest
	if err := ctx.ShouldBindJSON(&technicalService); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	var technicalServiceModel db.TechnicalService
	technicalServiceModel.ServiceName = technicalService.ServiceName
	technicalServiceModel.IdentityNumber = technicalService.IdentityNumber
	technicalServiceModel.PhoneNumber = technicalService.PhoneNumber
	technicalServiceModel.Email = technicalService.Email
	technicalServiceModel.Iban = technicalService.Iban

	err := t.technicalServiceService.Create(&technicalServiceModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, technicalServiceModel)
	ctx.JSON(http.StatusOK, response)
}

func (t *technicalServiceHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *technicalServiceHandler) GetAll(ctx *gin.Context) {

	technicalServices, err := t.technicalServiceService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, technicalServices)
	ctx.JSON(http.StatusOK, response)
}
