package migrations

import (
	"goShortURL/internal/database"
)

func SyncModels() {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		panic("failed to sync models")
	}
}
