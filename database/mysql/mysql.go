package mysql

import (
	"log"
	"os"

	authorEntity "github.com/ezkahan/book_store/src/modules/author/entity"
	bannerEntity "github.com/ezkahan/book_store/src/modules/banner/entity"
	bookEntity "github.com/ezkahan/book_store/src/modules/book/entity"
	categoryEntity "github.com/ezkahan/book_store/src/modules/category/entity"
	commentEntity "github.com/ezkahan/book_store/src/modules/comment/entity"
	favoriteEntity "github.com/ezkahan/book_store/src/modules/favorite/entity"
	imageEntity "github.com/ezkahan/book_store/src/modules/image/entity"
	userEntity "github.com/ezkahan/book_store/src/modules/user/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	}); err != nil {
		panic("Failed to connect to MySQL")
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}

func SyncDatabase() {
	db := Connection()

	db.AutoMigrate(
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
