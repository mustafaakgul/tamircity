package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

type UserStore interface {
	Create(model *db.User) error
	FindByEmail(email string) (*db.User, error)
	FindById(id int) (*db.User, error)
	//Login(model *db.User) error
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db: db}
}

func (u *userStore) Create(model *db.User) error {
	return u.db.Create(model).Error
}

func (u *userStore) FindByEmail(email string) (*db.User, error) {
	var user *db.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *userStore) FindById(id int) (*db.User, error) {
	var user *db.User
	err := u.db.Where("id = ?", id).First(&user).Error
	return user, err
}
