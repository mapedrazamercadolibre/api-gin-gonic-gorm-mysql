package Configuration

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfiguration struct {
	host     string
	port     int
	user     string
	dbName   string
	password string
}

func BuildDBConfiguration() *DBConfiguration {
	dbConfiguration := DBConfiguration{
		host:     "localhost",
		port:     3306,
		user:     "usrtest",
		password: "Qwerty",
		dbName:   "Clientes",
	}

	return &dbConfiguration
}

func DBURL(dbConfiguration *DBConfiguration) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		dbConfiguration.user,
		dbConfiguration.password,
		dbConfiguration.host,
		dbConfiguration.port,
		dbConfiguration.dbName,
	)

}
