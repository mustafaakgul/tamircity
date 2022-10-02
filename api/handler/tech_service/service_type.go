package tech_service

import (
	"github.com/anthophora/tamircity/pkg/service/tech_service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceTypeHandler interface {
	GetAll(ctx *gin.Context)
}

type serviceTypeHandler struct {
	serviceTypeService tech_service.ServiceTypeService
}

func NewServiceTypeHandler(serviceTypeService tech_service.ServiceTypeService) ServiceTypeHandler {
	return &serviceTypeHandler{
		serviceTypeService: serviceTypeService,
	}
}

func (s *serviceTypeHandler) GetAll(ctx *gin.Context) {

	serviceTypes, err := s.serviceTypeService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, serviceTypes)
	ctx.JSON(http.StatusOK, response)
}
