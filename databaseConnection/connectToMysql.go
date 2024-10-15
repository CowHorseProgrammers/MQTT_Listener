package databaseConnection

import (
	"MQTT_Middleware/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var MysqlClient *gorm.DB

func MysqlInit() {
	db, err := gorm.Open("mysql", config.GlobalConfig.Mysql.Url)
	if err != nil {
		log.Println("ERROR : Connect fail")
		panic("Mysql Connect fail")
	}
	log.Println("connect to mysql successfully")
	MysqlClient = db
}
