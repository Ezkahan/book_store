package repository

import (
	"errors"

	"github.com/ezkahan/book_store/src/modules/author/entity"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	Save(author entity.Author) (entity.Author, error)
	FindByID(id uint64) (entity.Author, error)
	FindBySlug(slug string) (entity.Author, error)
	List() entity.AuthorList
	Delete(slug string) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{
		db,
	}
}

func (repo *authorRepository) Save(author entity.Author) (entity.Author, error) {
	res := repo.db.Save(&author)
	return author, res.Error
}

func (repo *authorRepository) FindByID(id uint64) (entity.Author, error) {
	var author entity.Author
	res := repo.db.Find(&author, id)
	return author, res.Error
}

func (repo *authorRepository) FindBySlug(slug string) (entity.Author, error) {
	var author entity.Author
	res := repo.db.Where("slug = ?", slug).First(&author)
	return author, res.Error
}

func (repo *authorRepository) List() entity.AuthorList {
	var authors entity.AuthorList
	repo.db.Find(&authors)
	return authors
}

func (repo *authorRepository) Delete(slug string) error {
	var author entity.Author
	res := repo.db.Where("slug = ?", slug).Delete(&author)
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
