package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertisePhoneInfoStore struct {
	db *gorm.DB
}

type ExpertisePhoneInfoStore interface {
	Create(model *db.ExpertisePhoneInfo) error
	GetByID(id int) (*db.ExpertisePhoneInfo, error)
}

func NewExpertisePhoneInfoStore(db *gorm.DB) ExpertisePhoneInfoStore {
	return &expertisePhoneInfoStore{db: db}
}

func (e *expertisePhoneInfoStore) GetByID(id int) (*db.ExpertisePhoneInfo, error) {
	var expertisePhoneInfo db.ExpertisePhoneInfo
	err := e.db.First(&expertisePhoneInfo, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservation).Error
	return &expertisePhoneInfo, err
}

func (c *expertisePhoneInfoStore) Create(model *db.ExpertisePhoneInfo) error {
	return c.db.Create(model).Error
}
