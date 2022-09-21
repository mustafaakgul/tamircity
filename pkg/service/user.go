package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
	"github.com/anthophora/tamircity/pkg/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userStore repositories.UserStore
}

type UserService interface {
	Create(model *db.User) error
	FindByEmail(email string) (*db.User, error)
	FindById(id int) (*db.User, error)
	LoginCheck(models *db.User, password string) (string, error)
}

func NewUserService(userStore repositories.UserStore) UserService {
	return &userService{
		userStore: userStore,
	}
}

func (u *userService) Create(model *db.User) error {
	return u.userStore.Create(model)
}

func (u *userService) FindByEmail(email string) (*db.User, error) {
	return u.userStore.FindByEmail(email)
}

func (u *userService) FindById(id int) (*db.User, error) {
	return u.userStore.FindById(id)
}

func (u *userService) LoginCheck(models *db.User, password string) (string, error) {
	var err error

	/*result := VerifyPassword(password, models.Password)
	if result != true {
		return "Passwords not matched", err
	}*/

	err = VerifyPasswordHash(password, models.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(models.ID)
	if err != nil {
		return "Token was not created successfully.", err
	}

	return token, nil
}

func VerifyPassword(password, hashedPassword string) bool {
	if password != hashedPassword {
		return false
	} else {
		return true
	}
}

func VerifyPasswordHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
