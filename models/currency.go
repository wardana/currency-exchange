package models

import "time"

//Currency is a data strucrture for currency entities
type Currency struct {
	ID        uint64     `gorm:"primary_key" gorm:"column:id" json:"id"`
	ISOCode   string     `gorm:"column:iso_code" json:"code"`
	Name      string     `gorm:"column:name" json:"name"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" json:"-"`
}

//TableName get current table name
func (Currency) TableName() string {
	return "currency"
}
