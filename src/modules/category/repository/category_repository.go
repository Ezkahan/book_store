package repository

import (
	"errors"

	"github.com/ezkahan/book_store/src/modules/category/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category entity.Category) (entity.Category, error)
	Update(category entity.Category, slug string) (entity.Category, error)
	FindByID(id uint64) (entity.Category, error)
	FindBySlug(slug string) (entity.Category, error)
	List() entity.CategoryList
	Delete(slug string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db,
	}
}

func (repo *categoryRepository) Create(category entity.Category) (entity.Category, error) {
	res := repo.db.Create(&category)
	return category, res.Error
}

func (repo *categoryRepository) Update(category entity.Category, slug string) (entity.Category, error) {
	res, err := repo.FindBySlug(slug)
	if err != nil {
		return res, err
	}

	res.Title = category.Title
	res.Description = category.Description
	repo.db.Save(&res)

	return category, nil
}

func (repo *categoryRepository) FindByID(id uint64) (entity.Category, error) {
	var category entity.Category
	res := repo.db.Find(&category, id)
	return category, res.Error
}

func (repo *categoryRepository) FindBySlug(slug string) (entity.Category, error) {
	var category entity.Category
	res := repo.db.Where("slug = ?", slug).First(&category)
	return category, res.Error
}

func (repo *categoryRepository) List() entity.CategoryList {
	var categories entity.CategoryList
	repo.db.Find(&categories)
	return categories
}

func (repo *categoryRepository) Delete(slug string) error {
	var category entity.Category
	res := repo.db.Where("slug = ?", slug).Delete(&category)
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
