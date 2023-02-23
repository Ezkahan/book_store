package category

import (
	"net/http"

	"github.com/ezkahan/book_store/src/modules/category/request"
	"github.com/ezkahan/book_store/src/modules/category/service"
	"github.com/ezkahan/book_store/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService,
	}
}

func (h *categoryHandler) Create(ctx *gin.Context) {
	var (
		request  request.CategoryRequest
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

	category, err := h.categoryService.Create(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"data": category,
	})
}

func (h *categoryHandler) Update(ctx *gin.Context) {
	var (
		request  request.CategoryRequest
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

	category, err := h.categoryService.Update(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"data": category,
	})
}

func (h *categoryHandler) Get(ctx *gin.Context) {
	slug := ctx.Param("slug")
	category, err := h.categoryService.Get(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (h *categoryHandler) Delete(ctx *gin.Context) {
	slug := ctx.Param("slug")
	err := h.categoryService.Delete(slug)
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

func (h *categoryHandler) List(ctx *gin.Context) {
	categories := h.categoryService.List()

	ctx.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
