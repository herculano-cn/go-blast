package data

import (
	"fmt"

	"blast/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	config, err := config.LoadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.SSLMode,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// err = database.AutoMigrate(
	// 	&models.SomeModel{},
	// )
	// if err != nil {
	// 	return
	// }

	DB = database
}
