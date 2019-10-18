package models

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var mysql *gorm.DB

func init(){
	print("Connecting to db...")
	e := godotenv.Load()

	if e!=nil {
		fmt.Print("Error in Loading file", e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	dbName := os.Getenv("db_name")
	dbType := os.Getenv("db_type")

	dbUri := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", username, password, dbHost, dbName)

	conn, err := gorm.Open(dbType, dbUri)

	if err!=nil{
		fmt.Print("Error in connecting db", err)
	}else {
		fmt.Print("DB Connection successfull")
	}

	mysql = conn

}

func GetDB() *gorm.DB {
	return mysql
}
