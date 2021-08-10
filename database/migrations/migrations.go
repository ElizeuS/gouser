package migrations

import (
	"github.com/ElizeuS/gouser/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
