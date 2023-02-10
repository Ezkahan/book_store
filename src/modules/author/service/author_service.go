package service

import (
	"path/filepath"

	"github.com/ezkahan/book_store/src/modules/author/entity"
	"github.com/ezkahan/book_store/src/modules/author/repository"
	"github.com/ezkahan/book_store/src/modules/author/request"
	"github.com/gin-gonic/gin"
)

type AuthorService interface {
	Add(ctx *gin.Context, request request.AddAuthorRequest) (entity.Author, error)
	Update()
	Get(id uint64) (entity.Author, error)
	List() entity.AuthorList
	Delete(id uint64) error
}

type authorService struct {
	authorRepository repository.AuthorRepository
}

func NewAuthorService(authorRepository repository.AuthorRepository) AuthorService {
	return &authorService{
		authorRepository,
	}
}

func (service *authorService) Add(ctx *gin.Context, request request.AddAuthorRequest) (entity.Author, error) {
	image_url := service.SaveImage(ctx)

	author := entity.Author{
		Slug:  request.Slug,
		Name:  request.Name,
		Image: image_url,
	}

	author, err := service.authorRepository.Save(author)
	return author, err
}

func (service *authorService) Update() {
	//
}

func (service *authorService) List() entity.AuthorList {
	authors := service.authorRepository.List()
	return authors
}

func (service *authorService) Get(id uint64) (entity.Author, error) {
	author, err := service.authorRepository.FindByID(id)
	return author, err
}

func (service *authorService) Delete(id uint64) error {
	err := service.authorRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *authorService) SaveImage(ctx *gin.Context) string {
	file, _ := ctx.FormFile("image")
	name := filepath.Base(file.Filename)
	ctx.SaveUploadedFile(file, "assets/images/authors/"+name)

	return "assets/images/authors/" + name
}
