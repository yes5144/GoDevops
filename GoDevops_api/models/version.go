package models

import "github.com/jinzhu/gorm"

// Version xxx
type Version struct {
	gorm.Model

	Project    string `json:"project,omitempty"`
	Zone       string `json:"zone,omitempty"`
	Version    string `json:"version,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	IsDeleted  bool   `json:"is_deleted,omitempty"`
}
