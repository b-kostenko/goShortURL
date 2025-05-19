package auth

import "goShortURL/src/database"

func SyncModels() {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		panic("failed to sync models")
	}
}
