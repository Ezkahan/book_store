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
	Get(slug string) (entity.Author, error)
	List() entity.AuthorList
	Delete(slug string) error
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

func (service *authorService) Get(slug string) (entity.Author, error) {
	author, err := service.authorRepository.FindBySlug(slug)
	return author, err
}

func (service *authorService) Delete(slug string) error {
	err := service.authorRepository.Delete(slug)
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
