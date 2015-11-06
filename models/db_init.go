package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

type Configuration struct {
	DB struct {
		User, Dbname, Password, Host string
	}
    Environment     string
}

var Conf Configuration
var Connection gorm.DB

func IsProduction() bool {
    return Conf.Environment == "production"
}

func init() {

	conf_file, err := os.Open("conf.json")
	if err != nil {
		panic("Error when reading configuration file")
	}

	decoder := json.NewDecoder(conf_file)
	decode_error := decoder.Decode(&Conf)

	if decode_error != nil {
		panic(decode_error.Error())
	}

	Connection, err = gorm.Open("postgres", "user="+Conf.DB.User+" dbname="+Conf.DB.Dbname+" host="+Conf.DB.Host+" password="+Conf.DB.Password)
	if err != nil {
		fmt.Println(err.Error())
	}
    if !IsProduction() {
        Connection.LogMode(true)
    }
}
