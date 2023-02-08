package banner

import (
	"net/http"
	"strconv"

	"github.com/ezkahan/book_store/src/modules/banner/request"
	"github.com/ezkahan/book_store/src/modules/banner/service"
	"github.com/ezkahan/book_store/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BannerHandler interface {
	GetBanners(ctx *gin.Context)
	Add(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bannerHandler struct {
	bannerService service.BannerService
}

func NewBannerHandler(bannerService service.BannerService) BannerHandler {
	return &bannerHandler{
		bannerService,
	}
}

func (h *bannerHandler) GetBanners(ctx *gin.Context) {
	res := h.bannerService.List()

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (h *bannerHandler) Add(ctx *gin.Context) {
	var (
		request  request.AddBannerRequest
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

	banner := h.bannerService.Add(ctx, request)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": banner,
	})
}

func (h *bannerHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	err := h.bannerService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "Success deleted",
	})
}
