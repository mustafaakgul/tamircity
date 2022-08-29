package handler

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/models/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type brandHandler struct {
	brandService service.BrandService
}

type BrandHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByDeviceTypeId(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewBrandHandler(brandService service.BrandService) BrandHandler {
	return &brandHandler{
		brandService: brandService,
	}
}

func (b *brandHandler) GetAll(ctx *gin.Context) {
	brands, err := b.brandService.FindAll()
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, brands)
	ctx.JSON(http.StatusOK, response)
}

func (b *brandHandler) GetAllByDeviceTypeId(ctx *gin.Context) {

	deviceTypeId, err := strconv.Atoi(ctx.Param("device_type_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	brands, err := b.brandService.FindByDeviceTypeId(deviceTypeId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, brands)
	ctx.JSON(http.StatusOK, response)
}

func (b *brandHandler) Create(ctx *gin.Context) {
	var brand web.BrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var brandModel db.Brand
	brandModel.Name = brand.Name
	brandModel.IsActive = brand.IsActive

	err := b.brandService.Create(&brandModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, brandModel)
	ctx.JSON(http.StatusOK, response)
}
