package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (err error) {
	dsn := "root:123456@tcp(101.89.162.244:35784)/community?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func Migrate(model interface{}) error {
	err := DB.AutoMigrate(&model)
	return err
}
