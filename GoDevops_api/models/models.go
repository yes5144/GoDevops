package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB
var DB = db

// Model xxx
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// InitDB xxx
func InitDB() *gorm.DB {
	var err error
	// some variables
	dbType := viper.GetString("datasource.driverName")
	dbuser := viper.GetString("datasource.dbuser")
	dbpass := viper.GetString("datasource.dbpass")
	dbhost := viper.GetString("datasource.dbhost")
	dbport := viper.GetString("datasource.dbport")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	tablePrefix := viper.GetString("datasource.tablePrefix")

	// connInfo := "root:channel@tcp(192.168.204.222:3306)/godevops?charset=utf8&parseTime=True&loc=Local"
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbuser, dbpass, dbhost, dbport, dbname, charset)

	db, err = gorm.Open(dbType, connInfo)
	if err != nil {
		log.Printf("gorm.Open db err, Failed code %#v", err)
		panic(fmt.Sprintf("gorm.Open db err, Failed code %v", err))
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&User{})
	DB = db
	return DB
}

// GetDB xxx
func GetDB() *gorm.DB {
	DB = db
	return DB
}

// CloseDB xxx
func CloseDB() {
	defer DB.Close()
}
