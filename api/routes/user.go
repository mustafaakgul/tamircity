package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/anthophora/tamircity/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, userHandler handler.UserHandler) {
	userRoute := router.Group("api/v1/user")
	{
		userRoute.POST("", userHandler.Create)
		userRoute.POST("/login-cookie", userHandler.Login)
		userRoute.POST("/login-jwt", userHandler.LoginJwt)
		userRoute.GET("/validate", middleware.RequireAuth, userHandler.Validate)

		userRouteProtected := router.Group("/protected")
		userRouteProtected.Use(middleware.JwtAuthMiddleware())
		userRouteProtected.GET("/user", userHandler.GetCurrentUser)
	}

	/*protected := api.Group("/protected").Use(middlewares.Authz())
	{
		protected.GET("/profile", controllers.Profile)
	}*/
}
