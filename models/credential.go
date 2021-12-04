package models

import (
	"gorm.io/gorm"
)

type Credential struct {
	ID       uint
	Host     string `gorm:"default:nil"`
	App      string `gorm:"default:nil"`
	User     string `gorm:"default:nil"`
	Password string `gorm:"default:nil"`
}

func (c *Credential) Get(db *gorm.DB) []*Credential {
	var creds []*Credential
	if c.Host == "*" {
		db.Find(&creds)
	} else {
		db.Where("host = ? OR app = ? OR user = ?", c.Host, c.App, c.User).Find(&creds)
	}
	return creds
}

func (c *Credential) Save(db *gorm.DB) {
	db.Create(&c)
}

func (c *Credential) Update(id uint, db *gorm.DB) {
	db.Model(&c).Where("id = ?", c.ID).Updates(&c)
}

func (c *Credential) Delete(db *gorm.DB) {
	db.Delete(&Credential{}, c.ID)
}

func (c *Credential) Purge(db *gorm.DB, toPurge string) {
	if toPurge == "*" {
		db.Where("host like ?", "%%").Delete(&Credential{})
	} else {
		db.Where("host like ?", toPurge).Delete(&Credential{})
	}
}
