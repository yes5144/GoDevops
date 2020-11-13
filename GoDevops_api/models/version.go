package models

import (
	"log"
	"strings"
)

// Version xxx
type Version struct {
	// gorm.Model
	BaseModel

	Project    string `json:"project,omitempty"`
	Zone       string `json:"zone,omitempty"`
	Version    string `json:"version,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	IsDeleted  bool   `json:"is_deleted,omitempty"`
}

func (v *Version) GetAll() ([]Version, error) {
	var version []Version
	DB.Select([]string{"id", "project", "zone", "version", "update_time"}).Find(&version)
	log.Printf("from model getall: %#v ", version)
	return version, nil
}

func (v *Version) GetByIds(ids string) ([]Version, error) {
	var version []Version
	log.Printf("------%#v", ids)
	DB.Select([]string{"id", "project", "zone", "version", "update_time"}).Where("id in (?)", strings.Split(ids, ",")).Find(&version)
	log.Printf("from model getbyids: %#v ", version)
	return version, nil
}

func (v *Version) Create() error {
	return nil
}

func (v *Version) Update() error {
	return nil
}

func (v *Version) Delete() error {
	return nil
}

func InitVersion() {
	for _, p := range []string{"nz", "xkx2", "dj"} {
		log.Println("project:", p)
		for _, z := range []string{"wx", "hf", "qqgame"} {
			log.Println("zone:", z)
			DB.Create(&Version{
				Project: p,
				Zone:    z,
				Version: "1.0.23.2",
			})
		}
	}
}
