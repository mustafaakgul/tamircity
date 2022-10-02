package handler

import (
	"github.com/anthophora/tamircity/pkg/middleware"
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type technicalServiceCandidateHandler struct {
	technicalServiceCandidateService service.TechnicalServiceCandidateService
}

type TechnicalServiceCandidateHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	PrepareAndSendAgreement(ctx *gin.Context)
	ChangeStatusAndCreateTechnicalService(ctx *gin.Context)
}

func NewTechnicalServiceCandidateHandler(technicalServiceCandidateService service.TechnicalServiceCandidateService) TechnicalServiceCandidateHandler {
	return &technicalServiceCandidateHandler{
		technicalServiceCandidateService: technicalServiceCandidateService,
	}
}

func (c *technicalServiceCandidateHandler) Create(ctx *gin.Context) {
	var technicalServiceCandidate web.TechnicalServiceCandidateRequest
	if err := ctx.ShouldBindJSON(&technicalServiceCandidate); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var technicalServiceCandidateModel tech_service.TechnicalServiceCandidate
	technicalServiceCandidateModel.ServiceName = technicalServiceCandidate.ServiceName
	technicalServiceCandidateModel.BusinessType = technicalServiceCandidate.BusinessType
	technicalServiceCandidateModel.Address = technicalServiceCandidate.Address
	technicalServiceCandidateModel.NumberofBranches = technicalServiceCandidate.NumberofBranches
	technicalServiceCandidateModel.NumberofTechnicians = technicalServiceCandidate.NumberofTechnicians
	technicalServiceCandidateModel.Name = technicalServiceCandidate.Name
	technicalServiceCandidateModel.Surname = technicalServiceCandidate.Surname
	technicalServiceCandidateModel.Email = technicalServiceCandidate.Email
	technicalServiceCandidateModel.PhoneNumber = technicalServiceCandidate.PhoneNumber
	technicalServiceCandidateModel.Status = tech_service.TechnicalServiceCandidateStatus_PENDING

	err := c.technicalServiceCandidateService.Create(&technicalServiceCandidateModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be created", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Technical Service Candidate created successfully", nil, technicalServiceCandidateModel)
	ctx.JSON(http.StatusOK, response)
}

func (c *technicalServiceCandidateHandler) PrepareAndSendAgreement(ctx *gin.Context) {
	var emailRequest web.EmailRequest
	if err := ctx.ShouldBindJSON(&emailRequest); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// TODO: Prepare Agreement as PDF
	// TODO: Save PDF

	middleware.SendEmail4Agreement(emailRequest.Email)
}

func (c *technicalServiceCandidateHandler) ChangeStatusAndCreateTechnicalService(ctx *gin.Context) {
	status, err := ctx.GetQuery("status")
	if err != true {
		response := utils.HandleResponseModel(false, "Status could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	id, err := ctx.GetQuery("id")
	if err != true {
		response := utils.HandleResponseModel(false, "Id could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	idInt, _ := strconv.Atoi(id)

	if status != "approved" && status != "cancelled" {
		response := utils.HandleResponseModel(false, "Incorrect Query Parameters", nil, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	} else if status == "approved" {
		CreateApprovedFlow(ctx, c, idInt)
	} else if status == "cancelled" {
		CreateCancelledFlow(ctx, c, idInt)
	}
}

func (c *technicalServiceCandidateHandler) Update(ctx *gin.Context) {
	var technicalServiceCandidate web.TechnicalServiceCandidateRequest
	if err := ctx.ShouldBindJSON(&technicalServiceCandidate); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var technicalServiceCandidateModel tech_service.TechnicalServiceCandidate
	technicalServiceCandidateModel.ServiceName = technicalServiceCandidate.ServiceName
	technicalServiceCandidateModel.BusinessType = technicalServiceCandidate.BusinessType
	technicalServiceCandidateModel.Address = technicalServiceCandidate.Address
	technicalServiceCandidateModel.NumberofBranches = technicalServiceCandidate.NumberofBranches
	technicalServiceCandidateModel.NumberofTechnicians = technicalServiceCandidate.NumberofTechnicians
	technicalServiceCandidateModel.Name = technicalServiceCandidate.Name
	technicalServiceCandidateModel.Surname = technicalServiceCandidate.Surname
	technicalServiceCandidateModel.Email = technicalServiceCandidate.Email
	technicalServiceCandidateModel.PhoneNumber = technicalServiceCandidate.PhoneNumber
	technicalServiceCandidateModel.Status = tech_service.TechnicalServiceCandidateStatus_PENDING

	err := c.technicalServiceCandidateService.Update(&technicalServiceCandidateModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Technical Service Candidate updated successfully", nil, technicalServiceCandidateModel)
	ctx.JSON(http.StatusOK, response)
}

func (c *technicalServiceCandidateHandler) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	technicalServiceCandidate, err := c.technicalServiceCandidateService.FindByID(idInt)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleResponseModel(true, "Technical Service Candidate found successfully", nil, technicalServiceCandidate)
	ctx.JSON(http.StatusOK, response)
}

func CreateApprovedFlow(ctx *gin.Context, c *technicalServiceCandidateHandler, id int) {
	technicalServiceCandidate, err := c.technicalServiceCandidateService.FindByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	technicalServiceCandidate.Status = tech_service.TechnicalServiceCandidateStatus_APPROVED
	err = c.technicalServiceCandidateService.Update(&technicalServiceCandidate)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	var technicalService tech_service.TechnicalService
	technicalService.ServiceName = technicalServiceCandidate.ServiceName
	technicalService.IdentityNumber = "123"
	technicalService.PhoneNumber = technicalServiceCandidate.PhoneNumber
	technicalService.Email = technicalServiceCandidate.Email
	technicalService.Iban = "987"
	technicalService.Address = technicalServiceCandidate.Address

	// TODO: Create Technical Service
	// err = c.technicalServiceService.Create(&technicalService)
}

func CreateCancelledFlow(ctx *gin.Context, c *technicalServiceCandidateHandler, id int) {
	technicalServiceCandidate, err := c.technicalServiceCandidateService.FindByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	technicalServiceCandidate.Status = tech_service.TechnicalServiceCandidateStatus_CANCELLED
	err = c.technicalServiceCandidateService.Update(&technicalServiceCandidate)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
}
