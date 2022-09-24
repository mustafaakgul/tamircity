package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type reservationService struct {
	reservationStore repositories.ReservationStore
}

type ReservationService interface {
	Create(*web.ReservationCreateRequest) error
	GetPendingListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationPendingResponse, err error)
	GetCompletedListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationCompletedResponse, err error)
	GetCancelledListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationCancelledResponse, err error)
	GetPendingAndCompletedReservationCount(technicalServiceId int) (web.ReservationPendingAndCompletedCountResponse, error)
	UpdateReservationStatus(int, db.ReservationStatus) error
}

func NewReservationService(reservationStore repositories.ReservationStore) ReservationService {
	return &reservationService{
		reservationStore: reservationStore,
	}
}

// TO DO : Gerekli kontroller sağlanmalı her alan için.
func (r *reservationService) Create(reservationReq *web.ReservationCreateRequest) error {
	var reservation db.Reservation
	reservation.DeviceTypeId = reservationReq.DeviceTypeId
	reservation.BrandId = reservationReq.BrandId
	reservation.ModelId = reservationReq.ModelId
	reservation.FixTypeId = reservationReq.FixTypeId
	reservation.ServiceTypeId = reservationReq.ServiceTypeId
	reservation.ExtraServiceId = reservationReq.ExtraServiceId
	reservation.TechnicalServiceId = reservationReq.TechnicalServiceId
	reservation.ReservationDate = reservationReq.ReservationDate
	reservation.StartOfHour = reservationReq.StartOfHour
	reservation.EndOfHour = reservationReq.EndOfHour
	reservation.Price = reservationReq.Price
	reservation.FullName = reservationReq.FullName
	reservation.PhoneNumber = reservationReq.PhoneNumber
	reservation.SecondPhoneNumber = reservationReq.SecondPhoneNumber
	reservation.Description = reservationReq.Description

	return r.reservationStore.Create(&reservation)
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
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) GetCompletedListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationCompletedResponse, err error) {
	reservations, err := r.reservationStore.GetCompletedListByTechnicalServiceId(technicalServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationCompletedResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.ReservationDate
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.FixTypeName = reservation.FixType.Description // ?
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.ExtraServiceName = reservation.ExtraService.Description
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) GetCancelledListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationCancelledResponse, err error) {
	reservations, err := r.reservationStore.GetCancelledListByTechnicalServiceId(technicalServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationCancelledResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.ReservationDate
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.FixTypeName = reservation.FixType.Description // ?
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.ExtraServiceName = reservation.ExtraService.Description
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) UpdateReservationStatus(reservationId int, status db.ReservationStatus) error {
	return r.reservationStore.UpdateReservationStatus(reservationId, status)
}

func (r *reservationService) GetPendingAndCompletedReservationCount(technicalServiceId int) (web.ReservationPendingAndCompletedCountResponse, error) {
	var response web.ReservationPendingAndCompletedCountResponse
	a, err := r.reservationStore.GetReservationCountWithStatus(technicalServiceId, db.ReservationStatus(0))
	if err != nil {
		return response, err
	}
	response.PendingCount = a
	b, err := r.reservationStore.GetReservationCountWithStatus(technicalServiceId, db.ReservationStatus(1))
	if err != nil {
		return response, err
	}

	response.CompletedCount = b
	return response, nil
}
