package shortlink

import "goShortURL/src/database"

func SyncModels() {
	err := database.DB.AutoMigrate(&Shortlink{})
	if err != nil {
		panic("failed to sync models")
	}
}
