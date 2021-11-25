package configs

import (
	"arisan.com/arisan/models/anggota"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/arisan?charset=utf8mb4&parseTime=True&loc=Local"
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database gagal terkoneksi")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(anggota.Anggota{})
}
