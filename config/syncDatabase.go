package config

import (
	"github.com/ezkahan/book_store/database/mysql"
	authorEntity "github.com/ezkahan/book_store/src/modules/author/entity"
	bannerEntity "github.com/ezkahan/book_store/src/modules/banner/entity"
	bookEntity "github.com/ezkahan/book_store/src/modules/book/entity"
	categoryEntity "github.com/ezkahan/book_store/src/modules/category/entity"
	commentEntity "github.com/ezkahan/book_store/src/modules/comment/entity"
	favoriteEntity "github.com/ezkahan/book_store/src/modules/favorite/entity"
	imageEntity "github.com/ezkahan/book_store/src/modules/image/entity"
	userEntity "github.com/ezkahan/book_store/src/modules/user/entity"
)

func SyncDatabase() {
	mysql.DB.AutoMigrate(
		&userEntity.User{},
		&commentEntity.Comment{},
		&favoriteEntity.Favorite{},
		&bannerEntity.Banner{},
		&bookEntity.Book{},
		&categoryEntity.Category{},
		&imageEntity.Image{},
		&authorEntity.Author{},
	)
}
