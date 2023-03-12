package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type paymentInfoService struct {
	paymentInfoStore repositories.PaymentInfoStore
}

type PaymentInfoService interface {
	Create(request *web.PaymentInfoRequest) error
	FindAllByExpertiseServiceId(expertiseServiceId int) ([]web.PaymentInfoResponse, error)
	FindByID(id int) (res *web.PaymentInfoResponse, err error)
}

func NewPaymentInfoService(paymentInfoStore repositories.PaymentInfoStore) PaymentInfoService {
	return &paymentInfoService{
		paymentInfoStore: paymentInfoStore,
	}
}

func (p *paymentInfoService) Create(request *web.PaymentInfoRequest) error {
	var paymentInfo db.PaymentInfo
	paymentInfo.Amount = request.Amount
	paymentInfo.ReservationID = request.ReservationID
	paymentInfo.PaymentDate = request.PaymentDate
	paymentInfo.PaymentType = request.PaymentType

	return p.paymentInfoStore.Create(&paymentInfo)
}

func (p *paymentInfoService) FindAllByExpertiseServiceId(expertiseServiceId int) (res []web.PaymentInfoResponse, err error) {
	paymentInfos, err := p.paymentInfoStore.FindAllByExpertiseServiceId(expertiseServiceId)

	if err != nil {
		return nil, err
	}
	for _, paymentInfo := range paymentInfos {
		var paymentInfoResponse web.PaymentInfoResponse
		paymentInfoResponse.Id = paymentInfo.ID
		paymentInfoResponse.PaymentType = paymentInfo.PaymentType
		paymentInfoResponse.PaymentDate = paymentInfo.PaymentDate
		paymentInfoResponse.Amount = paymentInfo.Amount

		res = append(res, paymentInfoResponse)
	}

	return res, nil

}

func (p *paymentInfoService) FindByID(id int) (res *web.PaymentInfoResponse, err error) {
	paymentInfo, err := p.paymentInfoStore.FindByID(id)
	if err != nil {
		return nil, err
	}

	res.Id = paymentInfo.ID
	res.PaymentType = paymentInfo.PaymentType
	res.PaymentDate = paymentInfo.PaymentDate
	res.Amount = paymentInfo.Amount
	return res, nil
}
