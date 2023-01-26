package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertisePcInfoStore struct {
	db *gorm.DB
}

type ExpertisePcInfoStore interface {
	Create(model *db.ExpertisePcInfo) error
	GetByID(id int) (*db.ExpertisePcInfo, error)
}

func NewExpertisePcInfoStoreStore(db *gorm.DB) ExpertisePcInfoStore {
	return &expertisePcInfoStore{db: db}
}

func (e *expertisePcInfoStore) GetByID(id int) (*db.ExpertisePcInfo, error) {
	var expertisePcInfo db.ExpertisePcInfo
	err := e.db.First(&expertisePcInfo, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservation).Error
	return &expertisePcInfo, err
}

func (c *expertisePcInfoStore) Create(model *db.ExpertisePcInfo) error {
	return c.db.Create(model).Error
}
