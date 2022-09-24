package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/anthophora/tamircity/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	LoginJwt(ctx *gin.Context)
	Validate(ctx *gin.Context)
	GetCurrentUser(ctx *gin.Context)
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) Login(ctx *gin.Context) {
	var user web.UserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userResult, _ := u.userService.FindByEmail(user.Email)

	if userResult.ID == 0 {
		response := utils.HandleResponseModel(false, "User not found", nil, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if userResult.Password != user.Password {
		response := utils.HandleResponseModel(false, "Password is incorrect", nil, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userResult.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response := utils.HandleResponseModel(false, "Failed to create token", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  userResult,
	})

}

func (u *userHandler) Create(ctx *gin.Context) {
	var user web.UserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var userModel db.User
	userModel.Email = user.Email

	hashedPassword, err_hash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err_hash != nil {
		return
	}
	userModel.Password = string(hashedPassword)

	err := u.userService.Create(&userModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, userModel)
	ctx.JSON(http.StatusOK, response)
}

func (u *userHandler) Validate(ctx *gin.Context) {
	/*user, _ := ctx.Get("user")*/
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (u *userHandler) LoginJwt(ctx *gin.Context) {
	var user web.UserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResult, _ := u.userService.FindByEmail(user.Email)

	token, err := u.userService.LoginCheck(userResult, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *userHandler) GetCurrentUser(ctx *gin.Context) {
	/*user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})*/
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResult, _ := u.userService.FindById(int(user_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResult.Password = ""

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": userResult})
}
