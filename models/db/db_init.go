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
//	DB_HOST         = "54.93.171.33" // for development
	DB_HOST         = "localhost" // for production
)

var Connection gorm.DB

func init() {
	var err error
	Connection, err = gorm.Open("postgres", "user="+DB_USER+" dbname="+DB_DBNAME+" host="+DB_HOST+" password="+DB_USERPASSWORD)
	if err != nil {
		fmt.Println(err.Error())
	}
}
