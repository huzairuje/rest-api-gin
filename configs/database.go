package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func Database() *gorm.DB {

	//-----THIS CONFIGURATION JUST FOR POSTGRESQL
	connection := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USERNAME") +
		" dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" sslmode=" + os.Getenv("DB_SSL")
	db, err := gorm.Open(os.Getenv("DB_DRIVER"), connection)
	//-----THIS CONFIGURATION JUST FOR POSTGRESQL

	//-----THIS CONFIGURATION JUST FOR MYSQL
	//db, err := gorm.Open("mysql", "root:root@/rest_api_golang?charset=utf8&parseTime=True&loc=Local")
	//-----THIS CONFIGURATION JUST FOR MYSQL

	if err != nil {
		log.Fatal(err)
	}
	return db
}