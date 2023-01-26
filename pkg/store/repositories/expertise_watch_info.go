package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseWatchInfoStore struct {
	db *gorm.DB
}

type ExpertiseWatchInfoStore interface {
	Create(model *db.ExpertiseWatchInfo) error
	GetByID(id int) (*db.ExpertiseWatchInfo, error)
}

func NewExpertiseWatchInfoStoreStore(db *gorm.DB) ExpertiseWatchInfoStore {
	return &expertiseWatchInfoStore{db: db}
}

func (e *expertiseWatchInfoStore) GetByID(id int) (*db.ExpertiseWatchInfo, error) {
	var expertiseWatchInfo db.ExpertiseWatchInfo
	err := e.db.First(&expertiseWatchInfo, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservation).Error
	return &expertiseWatchInfo, err
}

func (c *expertiseWatchInfoStore) Create(model *db.ExpertiseWatchInfo) error {
	return c.db.Create(model).Error
}
