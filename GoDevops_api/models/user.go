package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// User xxx
type User struct {
	BaseModel
	Name      string `json:"name,omitempty"`
	Telephone string `json:"telephone,omitempty" gorm:"unique_index"`
	Password  string `json:"password,omitempty"`
}

// IsTelephoneExist xxx
func IsTelephoneExist(DB *gorm.DB, user User, telephone string) bool {
	DB.Where("telephone=?", telephone).First(&user)
	log.Println(user.ID)
	if user.ID > 0 {
		return true
	}
	return false
}
