package handler

import (
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type commentHandler struct {
	commentService service.CommentService
}

type CommentHandler interface {
	Create(ctx *gin.Context)
}

func NewCommentHandler(commentService service.CommentService) CommentHandler {
	return &commentHandler{
		commentService: commentService,
	}
}

func (c *commentHandler) Create(ctx *gin.Context) {
	var commentRequest web.CommentRequest
	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		response := utils.HandleResponseModel(false, "Wrong Request", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.commentService.Create(&commentRequest)
	if err != nil {
		response := utils.HandleResponseModel(false, "Comment could not be created", err, nil)
		ctx.JSON(http.StatusInternalServerError, response)
	}
	response := utils.HandleResponseModel(true, "Comment created successfully", nil, nil)
	ctx.JSON(http.StatusOK, response)
}
