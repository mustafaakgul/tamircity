package db

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	DeviceTypeId       int          `gorm:"not null"`
	DeviceType         *DeviceType  `gorm:"foreignkey:DeviceTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BrandId            int          `gorm:"not null"`
	Brand              *Brand       `gorm:"foreignkey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ModelId            int          `gorm:"not null"`
	ModelEntity        *Model       `gorm:"foreignkey:ModelId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ServiceTypeId      int          `gorm:"not null"`
	ServiceType        *ServiceType `gorm:"foreignkey:ServiceTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExpertiseServiceId int
	ExpertiseService   *ExpertiseService `gorm:"foreignkey:ExpertiseServiceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentInfo        *PaymentInfo
	Date               time.Time
	WeekDay            time.Weekday
	StartOfHour        int
	EndOfHour          int
	Price              int
	Status             ReservationStatus
	OperationalStatus  OperationStatus
	FullName           string `gorm:"not null"`
	Email              string `gorm:"not null"`
	PhoneNumber        string `gorm:"not null"`
	SecondPhoneNumber  string `gorm:"not null"`
	Description        string
}

type ReservationStatus int

const (
	Pending   ReservationStatus = 0 // Beklemede
	Cancelled                   = 1 // İptal Edildi
	Approved                    = 2 // Onaylandı
	Completed                   = 3 // Tamamlandı
)

type OperationStatus string

const (
	OperationStatus_WAITING_FOR_REPAIR OperationStatus = "WaitingforRepair" // İşlem Bekliyor
	OperationStatus_IN_PROGRESS        OperationStatus = "InProgress"       // İşleme Alındı
	OperationStatus_COMPLETED          OperationStatus = "Completed"        // Hazır Tamamlandı
	//OperationStatus_PRICE_OFFER_WILL_BE_GIVEN               OperationStatus = "PriceWillbeGiven"                  // Fiyat Verilecek
	//OperationStatus_SMS_WILL_BE_SENT_FOR_PRICE_CONFIRMATION OperationStatus = "SmsWillBeSentForPriceConfirmation" // Fiyat Onayı İçin SMS Gönderildi
	//OperationStatus_PENDING_PRICE_CONFIRMATION              OperationStatus = "PendingPriceConfirmation"          // Fiyat Onay Bekliyor
	//OperationStatus_PRICE_OFFER_NOT_APPROVED                OperationStatus = "PriceOfferNotApproved"             // Onaylanmadı
	//OperationStatus_WAITING_FOR_SPARE_PARTS                 OperationStatus = "WaitingforSpareParts"              // Parça Bekliyor
	//OperationStatus_THE_DEVICE_CANNOT_BE_REPAIRED           OperationStatus = "TheDeviceCannotBeRepaired"         // Onarım Yapılamıyor
	//OperationStatus_DEVICE_WILL_BE_RETURNED                 OperationStatus = "DeviceWillBeReturned"              // Cihaz İade Edilecek
	//OperationStatus_DEVICE_HAS_BEEN_RETURNED                OperationStatus = "DeviceHasBeenReturned"             // Cihaz İade Edildi
	OperationStatus_DEVICE_HAS_BEEN_DELIVERED OperationStatus = "DeviceHasBeenDelivered" // Cihaz Teslim Edildi
	//OperationStatus_OPERATION_HAS_BEEN_CANCELLED            OperationStatus = "OperationHasBeenCancelled"         // Servis İşlem İptal Edildi
)
