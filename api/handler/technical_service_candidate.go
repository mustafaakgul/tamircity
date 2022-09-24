package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type technicalServiceCandidateHandler struct {
	technicalServiceCandidateService service.TechnicalServiceCandidateService
}

type TechnicalServiceCandidateHandler interface {
	Create(ctx *gin.Context)
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

	var technicalServiceCandidateModel db.TechnicalServiceCandidate
	technicalServiceCandidateModel.ServiceName = technicalServiceCandidate.ServiceName
	technicalServiceCandidateModel.BusinessType = technicalServiceCandidate.BusinessType
	technicalServiceCandidateModel.Address = technicalServiceCandidate.Address
	technicalServiceCandidateModel.NumberofBranches = technicalServiceCandidate.NumberofBranches
	technicalServiceCandidateModel.NumberofTechnicians = technicalServiceCandidate.NumberofTechnicians
	technicalServiceCandidateModel.Name = technicalServiceCandidate.Name
	technicalServiceCandidateModel.Surname = technicalServiceCandidate.Surname
	technicalServiceCandidateModel.Email = technicalServiceCandidate.Email
	technicalServiceCandidateModel.PhoneNumber = technicalServiceCandidate.PhoneNumber
	technicalServiceCandidateModel.Status = db.TechnicalServiceCandidateStatus_PENDING

	err := c.technicalServiceCandidateService.Create(&technicalServiceCandidateModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Technical Service Candidate could not be created", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

	response := utils.HandleResponseModel(true, "Technical Service Candidate created successfully", nil, technicalServiceCandidateModel)
	ctx.JSON(http.StatusOK, response)
}
