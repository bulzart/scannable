package initializers

import "artwear/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.QR_{})
	DB.AutoMigrate(&models.QR_redirect{})

}
