package models

import (
	"encoding/json"
	"time"
)

//Rate is a data strucrture for exchange_rate entities
type Rate struct {
	ID              uint64     `gorm:"primary_key" gorm:"column:id" json:"id"`
	BaseCurrency    string     `gorm:"column:base_currency" json:"base_currency"`
	CounterCurrency string     `gorm:"column:counter_currency" json:"counter_currency"`
	ExchangeRate    float64    `gorm:"column:exchange_rate" json:"exchange_rate"`
	ExchangeDate    time.Time  `gorm:"column:exchange_date" json:"exchange_date"`
	CreatedAt       time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt       time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt       *time.Time `gorm:"column:deletedAt" json:"-"`
}

//UnmarshalJSON change data type from string to time
func (r *Rate) UnmarshalJSON(b []byte) error {
	var dateFormat = "2006-01-02" //YYYY-MM-DD
	type Alias Rate
	aux := &struct {
		ExchangeDate string `json:"exchange_date"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	t, err := time.Parse(dateFormat, aux.ExchangeDate)
	if err == nil {
		r.ExchangeDate = t
	}

	return nil
}

//TableName get current table name
func (Rate) TableName() string {
	return "rate"
}
