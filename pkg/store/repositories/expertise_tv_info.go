package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseTvInfoStore struct {
	db *gorm.DB
}

type ExpertiseTvInfoStore interface {
	Create(model *db.ExpertiseTvInfo) error
	GetByID(id int) (*db.ExpertiseTvInfo, error)
}

func NewExpertiseTvInfoStore(db *gorm.DB) ExpertiseTvInfoStore {
	return &expertiseTvInfoStore{db: db}
}

func (e *expertiseTvInfoStore) GetByID(id int) (*db.ExpertiseTvInfo, error) {
	var expertiseTvInfo db.ExpertiseTvInfo
	err := e.db.First(&expertiseTvInfo, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservation).Error
	return &expertiseTvInfo, err
}

func (c *expertiseTvInfoStore) Create(model *db.ExpertiseTvInfo) error {
	return c.db.Create(model).Error
}
