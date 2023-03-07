package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.Engine, commentHandler handler.CommentHandler) {
	route := router.Group("api/v1/comments")
	{
		route.POST("", commentHandler.Create)
	}
}
