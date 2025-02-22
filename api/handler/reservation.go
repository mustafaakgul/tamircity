package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
)

type reservationHandler struct {
	reservationService service.ReservationService
}

type ReservationHandler interface {
	Create(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	GetPendingList(ctx *gin.Context)
	GetCompletedList(ctx *gin.Context)
	GetCancelledList(ctx *gin.Context)
	GetApprovedList(ctx *gin.Context)
	GetApprovedListByExpertiseServiceIdAndDatetime(ctx *gin.Context)
	UpdateReservationStatus(ctx *gin.Context)
	GetPendingAndCompletedReservationCount(ctx *gin.Context)
	ChangeOperationStatus(ctx *gin.Context)
	GetReservationDetail(ctx *gin.Context)
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

func (r *reservationHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := utils.HandleResponseModel(false, "Wrong Params", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	reservation, err := r.reservationService.FindByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Reservation could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleResponseModel(true, "Reservation Candidate found successfully", nil, reservation)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) GetReservationDetail(ctx *gin.Context) {
	reservationId, err := strconv.Atoi(ctx.Query("reservation_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservation, err := r.reservationService.GetReservationDetail(reservationId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "Reservation not found", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservation)
	ctx.JSON(http.StatusOK, response)

}

func (r *reservationHandler) GetPendingList(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservations, err := r.reservationService.GetPendingListByExpertiseServiceId(expertiseServiceId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservations)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) GetCompletedList(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservations, err := r.reservationService.GetCompletedListByExpertiseServiceId(expertiseServiceId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservations)
	ctx.JSON(http.StatusOK, response)

}

func (r *reservationHandler) GetCancelledList(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservations, err := r.reservationService.GetCancelledListByExpertiseServiceId(expertiseServiceId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservations)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) GetApprovedList(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservations, err := r.reservationService.GetApprovedListByExpertiseServiceId(expertiseServiceId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservations)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) GetApprovedListByExpertiseServiceIdAndDatetime(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservationDateQuery := ctx.Query("reservation_date") + " " + "00:00"
	//reservationDateQuery := "2018-01-20 04:35"
	reservationDate, err := time.Parse("2006-01-02 15:04", reservationDateQuery)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservations, err := r.reservationService.GetApprovedListByExpertiseServiceIdAndDatetime(expertiseServiceId, reservationDate)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, reservations)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) UpdateReservationStatus(ctx *gin.Context) {
	reservationId, err := strconv.Atoi(ctx.Query("reservation_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	reservationStatus, err := strconv.Atoi(ctx.Query("reservation_status"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	if err := r.reservationService.UpdateReservationStatus(reservationId, db.ReservationStatus(reservationStatus)); err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "İşlem başarıyla gerçekleşmiştir.", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) GetPendingAndCompletedReservationCount(ctx *gin.Context) {
	expertiseServiceId, err := strconv.Atoi(ctx.Query("expertise_service_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	res, err := r.reservationService.GetPendingAndCompletedReservationCount(expertiseServiceId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, res)
	ctx.JSON(http.StatusOK, response)
}

func (r *reservationHandler) ChangeOperationStatus(ctx *gin.Context) {
	id, err := ctx.GetQuery("id")
	if err != true {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	operationStatus, err := ctx.GetQuery("operation_status")
	if err != true {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	idInt, _ := strconv.Atoi(id)

	fmt.Println(idInt, operationStatus)

	if err := r.reservationService.ChangeOperationStatus(idInt, db.OperationStatus(operationStatus)); err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "İşlem başarıyla gerçekleşmiştir.", nil, nil)
	ctx.JSON(http.StatusOK, response)
}
