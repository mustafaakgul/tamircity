package service

import (
	"time"

	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type reservationService struct {
	reservationStore repositories.ReservationStore
}

type ReservationService interface {
	Create(*web.ReservationCreateRequest) error
	FindByID(id int) (*db.Reservation, error)
	GetReservationDetail(reservationId int) (response *web.ReservationDetail, err error)
	GetPendingListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationPendingResponse, err error)
	GetCompletedListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationCompletedResponse, err error)
	GetCancelledListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationCancelledResponse, err error)
	GetApprovedListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationApprovedResponse, err error)
	GetApprovedListByExpertiseServiceIdAndDatetime(expertiseServiceId int, reservationDate time.Time) (response []web.ReservationApprovedResponse, err error)
	GetPendingAndCompletedReservationCount(expertiseServiceId int) (web.ReservationPendingAndCompletedCountResponse, error)
	UpdateReservationStatus(int, db.ReservationStatus) error
	ChangeOperationStatus(int, db.OperationStatus) error
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
	reservation.ServiceTypeId = reservationReq.ServiceTypeId
	reservation.ExpertiseServiceId = reservationReq.ExpertiseServiceId
	reservation.Date = reservationReq.ReservationDate
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

func (r *reservationService) FindByID(id int) (*db.Reservation, error) {
	return r.reservationStore.FindByID(id)
}

func (r *reservationService) GetReservationDetail(reservationId int) (*web.ReservationDetail, error) {
	reservation, err := r.reservationStore.GetReservationDetail(reservationId)

	if err != nil {
		return nil, err
	}
	var response web.ReservationDetail
	response.ReservationId = int(reservation.Model.ID)
	response.Email = reservation.Email
	response.FullName = reservation.FullName
	response.BrandId = reservation.BrandId
	response.BrandName = reservation.Brand.Name
	response.DeviceTypeId = reservation.DeviceTypeId
	response.ModelId = reservation.ModelId
	response.ModelName = reservation.ModelEntity.Name

	return &response, nil
}

func (r *reservationService) GetPendingListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationPendingResponse, err error) {
	reservations, err := r.reservationStore.GetPendingListByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationPendingResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.Date
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) GetCompletedListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationCompletedResponse, err error) {
	reservations, err := r.reservationStore.GetCompletedListByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationCompletedResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.Date
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) GetApprovedListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationApprovedResponse, err error) {
	reservations, err := r.reservationStore.GetApprovedListByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationApprovedResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.Date
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) GetApprovedListByExpertiseServiceIdAndDatetime(expertiseServiceId int, reservationDate time.Time) (response []web.ReservationApprovedResponse, err error) {
	reservations, err := r.reservationStore.GetApprovedListByExpertiseServiceIdAndDatetime(expertiseServiceId, reservationDate)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationApprovedResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.Date
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
		reservationResponse.FullName = reservation.FullName
		reservationResponse.Email = reservation.Email
		reservationResponse.PhoneNumber = reservation.PhoneNumber

		response = append(response, reservationResponse)
	}

	return response, nil
}

func (r *reservationService) GetCancelledListByExpertiseServiceId(expertiseServiceId int) (response []web.ReservationCancelledResponse, err error) {
	reservations, err := r.reservationStore.GetCancelledListByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		var reservationResponse web.ReservationCancelledResponse
		reservationResponse.ReservationId = int(reservation.ID)
		reservationResponse.ReservationDate = reservation.Date
		reservationResponse.DeviceTypeName = reservation.DeviceType.Name
		reservationResponse.BrandName = reservation.Brand.Name
		reservationResponse.ModelName = reservation.ModelEntity.Name
		reservationResponse.ServiceTypeName = reservation.ServiceType.Description
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

func (r *reservationService) ChangeOperationStatus(reservationId int, operationStatus db.OperationStatus) error {
	return r.reservationStore.ChangeOperationStatus(reservationId, operationStatus)
}

func (r *reservationService) GetPendingAndCompletedReservationCount(expertiseServiceId int) (web.ReservationPendingAndCompletedCountResponse, error) {
	var response web.ReservationPendingAndCompletedCountResponse
	a, err := r.reservationStore.GetReservationCountWithStatus(expertiseServiceId, db.ReservationStatus(0))
	if err != nil {
		return response, err
	}
	response.PendingCount = a
	b, err := r.reservationStore.GetReservationCountWithStatus(expertiseServiceId, db.ReservationStatus(1))
	if err != nil {
		return response, err
	}

	response.CompletedCount = b
	return response, nil
}
