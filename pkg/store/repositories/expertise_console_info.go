package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseConsoleInfoStore struct {
	db *gorm.DB
}

type ExpertiseConsoleInfoStore interface {
	Create(model *db.ExpertiseConsoleInfo) error
	GetByID(id int) (*db.ExpertiseConsoleInfo, error)
}

func NewExpertiseConsoleInfoStore(db *gorm.DB) ExpertiseConsoleInfoStore {
	return &expertiseConsoleInfoStore{db: db}
}

func (e *expertiseConsoleInfoStore) GetByID(id int) (*db.ExpertiseConsoleInfo, error) {
	var expertiseConsoleInfo db.ExpertiseConsoleInfo
	err := e.db.First(&expertiseConsoleInfo, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservation).Error
	return &expertiseConsoleInfo, err
}

func (c *expertiseConsoleInfoStore) Create(model *db.ExpertiseConsoleInfo) error {
	return c.db.Create(model).Error
}
