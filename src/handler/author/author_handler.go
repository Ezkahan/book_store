package author

import (
	"net/http"

	"github.com/ezkahan/book_store/src/modules/author/request"
	"github.com/ezkahan/book_store/src/modules/author/service"
	"github.com/ezkahan/book_store/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthorHandler interface {
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type authorHandler struct {
	service service.AuthorService
}

func NewAuthorHandler(service service.AuthorService) AuthorHandler {
	return &authorHandler{
		service,
	}
}

func (h *authorHandler) Add(ctx *gin.Context) {
	var (
		request  request.AddAuthorRequest
		validate *validator.Validate = validator.New()
	)

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := validate.Struct(request)
	if err != nil {
		errors := utils.ParseValidationError(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	author, err := h.service.Add(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"data": author,
	})
}

func (h *authorHandler) Update(ctx *gin.Context) {
	//
}

func (h *authorHandler) Get(ctx *gin.Context) {
	slug := ctx.Param("slug")
	author, err := h.service.Get(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": author,
	})
}

func (h *authorHandler) List(ctx *gin.Context) {
	authors := h.service.List()

	ctx.JSON(http.StatusOK, gin.H{
		"data": authors,
	})
}

func (h *authorHandler) Delete(ctx *gin.Context) {
	slug := ctx.Param("slug")
	err := h.service.Delete(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success deleted",
	})
}
