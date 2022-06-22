package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"net/http"
)

type TechnicalServiceHandler interface {
	GetAll(ctx *gin.Context)
}

type technicalServiceHandler struct {
	technicalServiceService service.TechnicalServiceService
}

func NewTechnicalServiceHandler(technicalServiceService service.TechnicalServiceService) TechnicalServiceHandler {
	return &technicalServiceHandler{
		technicalServiceService: technicalServiceService,
	}
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
