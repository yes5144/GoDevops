package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Asset struct {
	Model `json:"model,omitempty"`

	Hostname   string `json:"hostname,omitempty"`
	AssetType  string `json:"asset_type,omitempty"`
	PubIP01    string `json:"pub_ip_01,omitempty"`
	PubIP02    string `json:"pub_ip_02,omitempty"`
	PriIP01    string `json:"pri_ip_01,omitempty"`
	PriIP02    string `json:"pri_ip_02,omitempty"`
	CPU        string `json:"cpu,omitempty"`
	Mem        string `json:"mem,omitempty"`
	Disk       string `json:"disk,omitempty"`
	Memo       string `json:"memo,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	Isdelete   int    `json:"isdelete,omitempty"`
}

func ExistArtcleByID(id int) bool {
	var asset Asset
	db.Select("id").Where("id=?", id).First(&asset)

	if asset.ID > 0 {
		return true
	}
	return false
}

func GetAssetTotal(maps interface{}) (count int) {
	db.Model(&Asset{}).Where(maps).Count(&count)
	return
}

func GetAssets(pageNum int, pageSize int, maps interface{}) (assets []Asset) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&assets)
	return
}

func GetAsset(id int) (asset Asset) {
	db.Where("id=?", id).First(&asset)
	// db.Model(&asset).Related(&asset.Tag)
	return
}

func EditAsset(id int, data interface{}) bool {
	db.Model(&Asset{}).Where("id=?", id).Update(data)
	return true
}

func AddAsset(data map[string]interface{}) bool {
	db.Create(&Asset{
		Hostname:  data["hostname"].(string),
		AssetType: data["asset_type"].(string),
		PubIP01:   data["pur_ip_01"].(string),
		PubIP02:   data["pur_ip_02"].(string),
		PriIP01:   data["pri_ip_01"].(string),
		PriIP02:   data["pri_ip_02"].(string),
		CPU:       data["cpu"].(string),
		Mem:       data["mem"].(string),
		Disk:      data["disk"].(string),
		Memo:      data["memo"].(string),
		CreatedBy: data["created_by"].(string),
	})

	return true
}

func DeleteAsset(id int) bool {
	db.Where("id=?", id).Delete(Asset{})
	return true
}

func (asset *Asset) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (asset *Asset) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
