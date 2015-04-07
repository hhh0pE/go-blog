package models

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
    DB_USER = "blog"
    DB_USERPASSWORD = "blogblog"
    DB_DBNAME = "blog"
    DB_HOST = "lesnoy.name"
)

var DB *sql.DB

func init() {
    var err error
    DB, err = sql.Open("postgres", "user="+DB_USER+" dbname="+DB_DBNAME+" host="+DB_HOST+" password="+DB_USERPASSWORD)
    if err != nil {
        fmt.Println(err.Error())
    }
}