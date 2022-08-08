package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/models/web"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"net/http"
	"strconv"
)

type reservationHandler struct {
	reservationService service.ReservationService
}

type ReservationHandler interface {
	Create(ctx *gin.Context)
	GetPendingList(ctx *gin.Context)
}

func NewReservationHandler(reservationService service.ReservationService) ReservationHandler {
	return &reservationHandler{
		reservationService: reservationService,
	}
}

func (r *reservationHandler) Create(ctx *gin.Context) {
	var reservationReq web.ReservationCreateRequest
	if err := ctx.ShouldBindJSON(&reservationReq); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := r.reservationService.Create(&reservationReq)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "Rezervasyon başarı ile oluşturulmuştur.", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) GetPendingList(ctx *gin.Context) {
	technicalServiceId, err := strconv.Atoi(ctx.Param("technical_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservations, err := r.reservationService.GetPendingListByTechnicalServiceId(technicalServiceId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservations)
	ctx.JSON(http.StatusOK, response)

}
