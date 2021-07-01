package main

import (
	"api-gin-gionic-gorm-mysql/Configuration"
	"api-gin-gionic-gorm-mysql/Models"
	"api-gin-gionic-gorm-mysql/Routers"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	//Configuration.DB, err = gorm.Open("mysql", Configuration.DBURL(Configuration.BuildDBConfiguration()))
	Configuration.DB, err = gorm.Open(mysql.Open(Configuration.DBURL(Configuration.BuildDBConfiguration())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Configuration.DB.AutoMigrate(&Models.User{})
	r := Routers.SetupRouter()
	//running
	r.Run()

}
