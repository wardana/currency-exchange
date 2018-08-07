package models

import "time"

//Rate is a data strucrture for currency_rate entities
type Rate struct {
	ID           uint64     `gorm:"primary_key" gorm:"column:id" json:"id"`
	FromCode     string     `gorm:"column:from_code" json:"from_code"`
	ToCode       string     `gorm:"column:to_code" json:"to_code"`
	Rate         float64    `gorm:"column:rate" json:"rate"`
	Average      float64    `gorm:"column:average" json:"average"`
	EndOfDayDate float64    `gorm:"column:end_of_day_rate" json:"endOfDayRate"`
	CreatedAt    time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt    *time.Time `gorm:"column:deletedAt" json:"-"`
}

//TableName get current table name
func (Rate) TableName() string {
	return "currency_rate"
}
