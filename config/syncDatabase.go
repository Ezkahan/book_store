package config

import (
	authorEntity "github.com/ezkahan/book_store/src/modules/author/entities"
	bannerEntity "github.com/ezkahan/book_store/src/modules/banner/entities"
	bookEntity "github.com/ezkahan/book_store/src/modules/book/entities"
	categoryEntity "github.com/ezkahan/book_store/src/modules/category/entities"
	imageEntity "github.com/ezkahan/book_store/src/modules/image/entities"
	userEntity "github.com/ezkahan/book_store/src/modules/user/entities"
)

func SyncDatabase() {
	DB.AutoMigrate(
		&userEntity.User{},
		&userEntity.Comment{},
		&userEntity.Favorite{},
		&bannerEntity.Banner{},
		&bookEntity.Book{},
		&categoryEntity.Category{},
		&imageEntity.Image{},
		&authorEntity.Author{},
	)
}
