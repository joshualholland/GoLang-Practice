package initializers

import "example/hello/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
