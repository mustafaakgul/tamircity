package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type fixTypeHandler struct {
	fixTypeService service.FixTypeService
}

type FixTypeService interface {
	GetAll(ctx *gin.Context)
}

func NewFixTypeHandler(fixTypeService service.FixTypeService) FixTypeService {
	return &fixTypeHandler{
		fixTypeService: fixTypeService,
	}
}

func (f *fixTypeHandler) GetAll(ctx *gin.Context) {
	fixTypes, err := f.fixTypeService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, fixTypes)
	ctx.JSON(http.StatusOK, response)
}
