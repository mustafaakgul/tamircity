package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/models/web"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type reservationService struct {
	reservationStore repositories.ReservationStore
}

type ReservationService interface {
	Create(*web.ReservationCreateRequest) error
	GetPendingListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationPendingResponse, err error)
	UpdateReservationStatus(int, db.ReservationStatus) error
}

func NewReservationService(reservationStore repositories.ReservationStore) ReservationService {
	return &reservationService{
		reservationStore: reservationStore,
	}
}

// TO DO : Gerekli kontroller sağlanmalı her alan için.
func (r *reservationService) Create(reservationReq *web.ReservationCreateRequest) error {
	var reservation *db.Reservation
	reservation.DeviceTypeId = reservationReq.DeviceTypeId
	reservation.BrandId = reservationReq.BrandId
	reservation.ModelId = reservationReq.ModelId
	reservation.FixTypeId = reservationReq.FixTypeId
	reservation.ServiceTypeId = reservationReq.ServiceTypeId
	reservation.ExtraServiceId = reservationReq.ExtraServiceId
	reservation.TechnicalServiceId = reservationReq.TechnicalServiceId
	reservation.ReservationDate = reservationReq.ReservationDate
	reservation.Price = reservationReq.Price
	reservation.FullName = reservationReq.FullName
	reservation.PhoneNumber = reservationReq.PhoneNumber
	reservation.SecondPhoneNumber = reservationReq.SecondPhoneNumber
	reservation.Description = reservationReq.Description

	return r.reservationStore.Create(reservation)
}

func (r *reservationService) GetPendingListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationPendingResponse, err error) {
	reservations, err := r.reservationStore.GetPendingListByTechnicalServiceId(technicalServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationPendingResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.ReservationDate
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.FixTypeName = reservation.FixType.Description // ?
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.ExtraServiceName = reservation.ExtraService.Description

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) UpdateReservationStatus(reservationId int, status db.ReservationStatus) error {
	return r.reservationStore.UpdateReservationStatus(reservationId, status)
}
