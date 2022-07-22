package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"net/http"
)

type ServiceTypeHandler interface {
	GetAll(ctx *gin.Context)
}

type serviceTypeHandler struct {
	serviceTypeService service.ServiceTypeService
}

func NewServiceTypeHandler(serviceTypeService service.ServiceTypeService) ServiceTypeHandler {
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
