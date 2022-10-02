package service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
	"time"
)

type reservationService struct {
	reservationStore tech_service2.ReservationStore
}

type ReservationService interface {
	Create(*web.ReservationCreateRequest) error
	FindByID(id int) (*tech_service.Reservation, error)
	GetPendingListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationPendingResponse, err error)
	GetCompletedListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationCompletedResponse, err error)
	GetCancelledListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationCancelledResponse, err error)
	GetApprovedListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationApprovedResponse, err error)
	GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId int, reservationDate time.Time) (response []web.ReservationApprovedResponse, err error)
	GetPendingAndCompletedReservationCount(technicalServiceId int) (web.ReservationPendingAndCompletedCountResponse, error)
	UpdateReservationStatus(int, tech_service.ReservationStatus) error
	ChangeOperationStatus(int, tech_service.OperationStatus) error
}

func NewReservationService(reservationStore tech_service2.ReservationStore) ReservationService {
	return &reservationService{
		reservationStore: reservationStore,
	}
}

// TO DO : Gerekli kontroller sağlanmalı her alan için.
func (r *reservationService) Create(reservationReq *web.ReservationCreateRequest) error {
	var reservation tech_service.Reservation
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
	reservation.Email = reservationReq.Email
	reservation.PhoneNumber = reservationReq.PhoneNumber
	reservation.SecondPhoneNumber = reservationReq.SecondPhoneNumber
	reservation.Description = reservationReq.Description

	return r.reservationStore.Create(&reservation)
}

func (r *reservationService) FindByID(id int) (*tech_service.Reservation, error) {
	return r.reservationStore.FindByID(id)
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

func (r *reservationService) GetApprovedListByTechnicalServiceId(technicalServiceId int) (response []web.ReservationApprovedResponse, err error) {
	reservations, err := r.reservationStore.GetApprovedListByTechnicalServiceId(technicalServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationApprovedResponse
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

func (r *reservationService) GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId int, reservationDate time.Time) (response []web.ReservationApprovedResponse, err error) {
	reservations, err := r.reservationStore.GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId, reservationDate)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationApprovedResponse
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

func (r *reservationService) UpdateReservationStatus(reservationId int, status tech_service.ReservationStatus) error {
	return r.reservationStore.UpdateReservationStatus(reservationId, status)
}

func (r *reservationService) ChangeOperationStatus(reservationId int, operationStatus tech_service.OperationStatus) error {
	return r.reservationStore.ChangeOperationStatus(reservationId, operationStatus)
}

func (r *reservationService) GetPendingAndCompletedReservationCount(technicalServiceId int) (web.ReservationPendingAndCompletedCountResponse, error) {
	var response web.ReservationPendingAndCompletedCountResponse
	a, err := r.reservationStore.GetReservationCountWithStatus(technicalServiceId, tech_service.ReservationStatus(0))
	if err != nil {
		return response, err
	}
	response.PendingCount = a
	b, err := r.reservationStore.GetReservationCountWithStatus(technicalServiceId, tech_service.ReservationStatus(1))
	if err != nil {
		return response, err
	}

	response.CompletedCount = b
	return response, nil
}
