package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/pkg/constants"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MYSQL_DEFAULT_DSN), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}
}
