package models

import (
    "github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER         = "blog"
	DB_USERPASSWORD = "blogblog"
	DB_DBNAME       = "blog"
    //DB_HOST         = "lesnoy.name" // for development
	DB_HOST         = "localhost" // for production
)

var DB gorm.DB

func init() {
    var err error
	DB, err = gorm.Open("postgres", "user="+DB_USER+" dbname="+DB_DBNAME+" host="+DB_HOST+" password="+DB_USERPASSWORD)
	if err != nil {
		fmt.Println(err.Error())
	}
}
