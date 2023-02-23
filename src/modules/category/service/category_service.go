package service

import (
	"path/filepath"

	"github.com/ezkahan/book_store/src/modules/category/entity"
	"github.com/ezkahan/book_store/src/modules/category/repository"
	"github.com/ezkahan/book_store/src/modules/category/request"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	Create(ctx *gin.Context, request request.CategoryRequest) (entity.Category, error)
	Update(ctx *gin.Context, request request.CategoryRequest) (entity.Category, error)
	Get(slug string) (entity.Category, error)
	List() entity.CategoryList
	Delete(slug string) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository,
	}
}

func (service *categoryService) Create(ctx *gin.Context, request request.CategoryRequest) (entity.Category, error) {
	icon_url := service.SaveIcon(ctx)

	category := entity.Category{
		Title:       request.Title,
		Description: request.Description,
		Icon:        icon_url,
	}

	return service.categoryRepository.Create(category)
}

func (service *categoryService) Update(ctx *gin.Context, request request.CategoryRequest) (entity.Category, error) {
	icon_url := service.SaveIcon(ctx)
	slug := ctx.Param("slug")

	category := entity.Category{
		Title:       request.Title,
		Description: request.Description,
		Icon:        icon_url,
	}

	return service.categoryRepository.Update(category, slug)
}

func (service *categoryService) Get(slug string) (entity.Category, error) {
	category, err := service.categoryRepository.FindBySlug(slug)
	return category, err
}

func (service *categoryService) List() entity.CategoryList {
	categories := service.categoryRepository.List()
	return categories
}

func (service *categoryService) Delete(slug string) error {
	err := service.categoryRepository.Delete(slug)
	if err != nil {
		return err
	}
	return nil
}

func (service *categoryService) SaveIcon(ctx *gin.Context) string {
	file, _ := ctx.FormFile("image")
	name := filepath.Base(file.Filename)
	path := "assets/images/categories/"
	ctx.SaveUploadedFile(file, path+name)

	return path + name
}
