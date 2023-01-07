package configs

import "github.com/ezkahan/golab/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
