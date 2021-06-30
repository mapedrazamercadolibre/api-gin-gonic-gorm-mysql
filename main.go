package main

import (
	"api-gin-gionic-gorm-mysql/Configuration"
	"api-gin-gionic-gorm-mysql/Models"
	"api-gin-gionic-gorm-mysql/Routers"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Configuration.DB, err = gorm.Open("mysql", Configuration.DBURL(Configuration.BuildDBConfiguration()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	//defer Configuration.DB.Close()
	Configuration.DB.AutoMigrate(&Models.User{})
	r := Routers.SetupRouter()
	//running
	r.Run()

}
