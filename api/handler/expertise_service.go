package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type expertiseServiceHandler struct {
	expertiseServiceService service.ExpertiseServiceService
}

type ExpertiseServiceHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByFilter(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewExpertiseServiceHandler(expertiseServiceService service.ExpertiseServiceService) ExpertiseServiceHandler {
	return &expertiseServiceHandler{
		expertiseServiceService: expertiseServiceService,
	}
}

func (t *expertiseServiceHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	expertiseService, err := t.expertiseServiceService.FindByID(idInt)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Expertise Service could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "Expertise Service Candidate found successfully", nil, expertiseService)
	ctx.JSON(http.StatusOK, response)
}

func (t *expertiseServiceHandler) Create(ctx *gin.Context) {
	var expertiseService web.ExpertiseServiceRequest
	if err := ctx.ShouldBindJSON(&expertiseService); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertiseServiceModel db.ExpertiseService
	expertiseServiceModel.ServiceName = expertiseService.ServiceName
	expertiseServiceModel.IdentityNumber = expertiseService.IdentityNumber
	expertiseServiceModel.PhoneNumber = expertiseService.PhoneNumber
	expertiseServiceModel.Email = expertiseService.Email
	expertiseServiceModel.Iban = expertiseService.Iban

	err := t.expertiseServiceService.Create(&expertiseServiceModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "", nil, expertiseServiceModel)
	ctx.JSON(http.StatusOK, response)
}

func (t *expertiseServiceHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	var expertiseService web.ExpertiseServiceUpdateRequest
	if err := ctx.ShouldBindJSON(&expertiseService); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	expertiseServiceModel, err := t.expertiseServiceService.FindByID(idInt)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Expertise Service could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	expertiseServiceModel.ServiceName = expertiseService.ServiceName
	expertiseServiceModel.Name = expertiseService.Name
	expertiseServiceModel.Surname = expertiseService.Surname
	expertiseServiceModel.Email = expertiseService.Email
	expertiseServiceModel.PhoneNumber = expertiseService.PhoneNumber
	expertiseServiceModel.About = expertiseService.About
	expertiseServiceModel.Address = expertiseService.Address

	err = t.expertiseServiceService.Update(&expertiseServiceModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Expertise Service updated successfully", nil, expertiseServiceModel)
	ctx.JSON(http.StatusOK, response)
}

func (t *expertiseServiceHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *expertiseServiceHandler) GetAll(ctx *gin.Context) {

	expertiseServices, err := t.expertiseServiceService.FindAll()

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

	response := utils.HandleResponseModel(true, "", nil, expertiseServices)
	ctx.JSON(http.StatusOK, response)
}

func (t *expertiseServiceHandler) GetAllByFilter(ctx *gin.Context) {

	brandId, err := strconv.Atoi(ctx.Query("brand_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Brand id not found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	deviceTypeId, err := strconv.Atoi(ctx.Query("device_type_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Device Type id not found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	var expertiseServices []web.ExpertiseServiceSearchResponse
	expertiseServices, err = t.expertiseServiceService.FindByBrandIdDeviceTypeId(brandId, deviceTypeId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Expertise Service not found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "Expertise Service found successfully", nil, expertiseServices)
	ctx.JSON(http.StatusOK, response)
}
