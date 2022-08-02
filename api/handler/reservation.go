package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/models/web"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"net/http"
)

type reservationHandler struct {
	reservationService service.ReservationService
}

type ReservationHandler interface {
	Create(ctx *gin.Context)
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
