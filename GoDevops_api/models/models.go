package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

//  BaseModel xxx
type BaseModel struct {
	// gorm.Model
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt string `gorm:"column:create_time" json:"create_time"`
	UpdatedAt string `gorm:"column:update_time" json:"update_time"`
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
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbuser, dbpass, dbhost, dbport, dbname, charset)

	log.Println(dbType, connInfo)
	// db, err = gorm.Open(dbType, connInfo)
	db, err = gorm.Open("sqlite3", "sqlite3.db")
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
	db.AutoMigrate(&User{}, &Version{})
	DB = db
	// InitVersion()
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

// NowTime xxx
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//
func (v BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", NowTime())
	scope.SetColumn("update_time", NowTime())
	return nil
}

//
func (v BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", NowTime())
	return nil
}
