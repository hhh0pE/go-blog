package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	DB_USER         = "blog"
	DB_USERPASSWORD = "blogblog"
	DB_DBNAME       = "blog"
	DB_HOST         = "lesnoy.name" // for development
	//DB_HOST         = "localhost" // for production
)

var Connection gorm.DB

func init() {
	var err error
	Connection, err = gorm.Open("postgres", "user="+DB_USER+" dbname="+DB_DBNAME+" host="+DB_HOST+" password="+DB_USERPASSWORD)
	if err != nil {
		fmt.Println(err.Error())
	}
}
