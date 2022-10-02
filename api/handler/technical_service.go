package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service3 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/service/tech_service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type technicalServiceHandler struct {
	technicalServiceService tech_service2.TechnicalServiceService
}

type TechnicalServiceHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByFilter(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewTechnicalServiceHandler(technicalServiceService tech_service2.TechnicalServiceService) TechnicalServiceHandler {
	return &technicalServiceHandler{
		technicalServiceService: technicalServiceService,
	}
}

func (t *technicalServiceHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	technicalService, err := t.technicalServiceService.FindByID(idInt)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Technical Service could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "Technical Service Candidate found successfully", nil, technicalService)
	ctx.JSON(http.StatusOK, response)
}

func (t *technicalServiceHandler) Create(ctx *gin.Context) {
	var technicalService tech_service3.TechnicalServiceRequest
	if err := ctx.ShouldBindJSON(&technicalService); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var technicalServiceModel tech_service.TechnicalService
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

func (t *technicalServiceHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	var technicalService tech_service3.TechnicalServiceUpdateRequest
	if err := ctx.ShouldBindJSON(&technicalService); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	technicalServiceModel, err := t.technicalServiceService.FindByID(idInt)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Technical Service could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	technicalServiceModel.ServiceName = technicalService.ServiceName
	technicalServiceModel.Name = technicalService.Name
	technicalServiceModel.Surname = technicalService.Surname
	technicalServiceModel.Email = technicalService.Email
	technicalServiceModel.PhoneNumber = technicalService.PhoneNumber
	technicalServiceModel.About = technicalService.About
	technicalServiceModel.Address = technicalService.Address

	err = t.technicalServiceService.Update(&technicalServiceModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Technical Service updated successfully", nil, technicalServiceModel)
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

	// Data Format
	//TODO: Must be Decided
	/*data := map[string]interface{}{
		"remainingCredit": remainingCredit,
		"meterTime":       meterTime,
	}*/

	response := utils.HandleResponseModel(true, "", nil, technicalServices)
	ctx.JSON(http.StatusOK, response)
}

func (t *technicalServiceHandler) GetAllByFilter(ctx *gin.Context) {

	modelId, err := strconv.Atoi(ctx.Query("model_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	var technicalServices []tech_service3.TechnicalServiceSearchResponse
	technicalServices, err = t.technicalServiceService.FindByModelId(modelId)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, technicalServices)
	ctx.JSON(http.StatusOK, response)
}
