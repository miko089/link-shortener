package database

import (
	"gorm.io/gorm"
	"linkShortener/utils"
)

type Link struct {
	Short string `gorm:"type:text"`
	Full  string `gorm:"type:text"`
}

func CreateLink(db *gorm.DB, full string) (*Link, error) {
	var short string
	for {
		short = utils.RandomLink()
		_, err := GetLink(db, short)
		if err != nil {
			break
		}
	}
	err := db.Create(&Link{Short: short, Full: full}).Error
	if err != nil {
		return nil, err
	}
	return GetLink(db, short)
}

func GetLink(db *gorm.DB, short string) (*Link, error) {
	var link Link
	b := short == ""
	if b {
		return nil, gorm.ErrRecordNotFound
	}
	err := db.Where("short = ?", short).First(&link).Error
	return &link, err
}
