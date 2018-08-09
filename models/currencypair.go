package models

import "time"

// CurrencyPair is a base currency exchange
type CurrencyPair struct {
	ID              int64      `gorm:"primary_key" json:"id"`
	BaseCurrency    string     `gorm:"column:base_currency" json:"base_currency"`
	CounterCurrency string     `gorm:"column:counter_currency" json:"counter_currency"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

//TableName get current table name
func (CurrencyPair) TableName() string {
	return "currency_pair"
}
