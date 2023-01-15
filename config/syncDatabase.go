package config

import (
	"github.com/ezkahan/golab/src/modules/user/entities"
)

func SyncDatabase() {
	DB.AutoMigrate(&entities.User{})
}
