package handler

import (
	"github.com/anthophora/tamircity/pkg/middleware"
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type expertiseServiceCandidateHandler struct {
	expertiseServiceCandidateService service.ExpertiseServiceCandidateService
}

type ExpertiseServiceCandidateHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	PrepareAndSendAgreement(ctx *gin.Context)
	ChangeStatusAndCreateExpertiseService(ctx *gin.Context)
}

func NewExpertiseServiceCandidateHandler(expertiseServiceCandidateService service.ExpertiseServiceCandidateService) ExpertiseServiceCandidateHandler {
	return &expertiseServiceCandidateHandler{
		expertiseServiceCandidateService: expertiseServiceCandidateService,
	}
}

func (c *expertiseServiceCandidateHandler) Create(ctx *gin.Context) {
	var expertiseServiceCandidate web.ExpertiseServiceCandidateRequest
	if err := ctx.ShouldBindJSON(&expertiseServiceCandidate); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertiseServiceCandidateModel db.ExpertiseServiceCandidate
	expertiseServiceCandidateModel.ServiceName = expertiseServiceCandidate.ServiceName
	expertiseServiceCandidateModel.BusinessType = expertiseServiceCandidate.BusinessType
	expertiseServiceCandidateModel.Address = expertiseServiceCandidate.Address
	expertiseServiceCandidateModel.NumberofBranches = expertiseServiceCandidate.NumberofBranches
	expertiseServiceCandidateModel.NumberofTechnicians = expertiseServiceCandidate.NumberofTechnicians
	expertiseServiceCandidateModel.Name = expertiseServiceCandidate.Name
	expertiseServiceCandidateModel.Surname = expertiseServiceCandidate.Surname
	expertiseServiceCandidateModel.Email = expertiseServiceCandidate.Email
	expertiseServiceCandidateModel.PhoneNumber = expertiseServiceCandidate.PhoneNumber
	expertiseServiceCandidateModel.Status = db.ExpertiseServiceCandidateStatus_PENDING

	err := c.expertiseServiceCandidateService.Create(&expertiseServiceCandidateModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be created", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Expertise Service Candidate created successfully", nil, expertiseServiceCandidateModel)
	ctx.JSON(http.StatusOK, response)
}

func (c *expertiseServiceCandidateHandler) PrepareAndSendAgreement(ctx *gin.Context) {
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

func (c *expertiseServiceCandidateHandler) ChangeStatusAndCreateExpertiseService(ctx *gin.Context) {
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

func (c *expertiseServiceCandidateHandler) Update(ctx *gin.Context) {
	var expertiseServiceCandidate web.ExpertiseServiceCandidateRequest
	if err := ctx.ShouldBindJSON(&expertiseServiceCandidate); err != nil {
		response := utils.HandleResponseModel(false, "Incorrect JSON Format", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertiseServiceCandidateModel db.ExpertiseServiceCandidate
	expertiseServiceCandidateModel.ServiceName = expertiseServiceCandidate.ServiceName
	expertiseServiceCandidateModel.BusinessType = expertiseServiceCandidate.BusinessType
	expertiseServiceCandidateModel.Address = expertiseServiceCandidate.Address
	expertiseServiceCandidateModel.NumberofBranches = expertiseServiceCandidate.NumberofBranches
	expertiseServiceCandidateModel.NumberofTechnicians = expertiseServiceCandidate.NumberofTechnicians
	expertiseServiceCandidateModel.Name = expertiseServiceCandidate.Name
	expertiseServiceCandidateModel.Surname = expertiseServiceCandidate.Surname
	expertiseServiceCandidateModel.Email = expertiseServiceCandidate.Email
	expertiseServiceCandidateModel.PhoneNumber = expertiseServiceCandidate.PhoneNumber
	expertiseServiceCandidateModel.Status = db.ExpertiseServiceCandidateStatus_PENDING

	err := c.expertiseServiceCandidateService.Update(&expertiseServiceCandidateModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Expertise Service Candidate updated successfully", nil, expertiseServiceCandidateModel)
	ctx.JSON(http.StatusOK, response)
}

func (c *expertiseServiceCandidateHandler) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	expertiseServiceCandidate, err := c.expertiseServiceCandidateService.FindByID(idInt)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleResponseModel(true, "Expertise Service Candidate found successfully", nil, expertiseServiceCandidate)
	ctx.JSON(http.StatusOK, response)
}

func CreateApprovedFlow(ctx *gin.Context, c *expertiseServiceCandidateHandler, id int) {
	expertiseServiceCandidate, err := c.expertiseServiceCandidateService.FindByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	expertiseServiceCandidate.Status = db.ExpertiseServiceCandidateStatus_APPROVED
	err = c.expertiseServiceCandidateService.Update(&expertiseServiceCandidate)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	var expertiseService db.ExpertiseService
	expertiseService.ServiceName = expertiseServiceCandidate.ServiceName
	expertiseService.IdentityNumber = "123"
	expertiseService.PhoneNumber = expertiseServiceCandidate.PhoneNumber
	expertiseService.Email = expertiseServiceCandidate.Email
	expertiseService.Iban = "987"
	expertiseService.Address = expertiseServiceCandidate.Address

	// TODO: Create Expertise Service
	// err = c.expertiseServiceService.Create(&expertiseService)
}

func CreateCancelledFlow(ctx *gin.Context, c *expertiseServiceCandidateHandler, id int) {
	expertiseServiceCandidate, err := c.expertiseServiceCandidateService.FindByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	expertiseServiceCandidate.Status = db.ExpertiseServiceCandidateStatus_CANCELLED
	err = c.expertiseServiceCandidateService.Update(&expertiseServiceCandidate)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Service Candidate could not be updated", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
}
