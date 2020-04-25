package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var err error
	// some variables
	dbType := "mysql"
	connInfo := "root:channel@tcp(192.168.204.222:3306)/godevops?charset=utf8&parseTime=True&loc=Local"
	tablePrefix := "wk"

	db, err = gorm.Open(dbType, connInfo)
	if err != nil {
		log.Fatalf("gorm.Open db err, Failed code %#v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	defer db.Close()
}
