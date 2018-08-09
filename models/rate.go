package models

import (
	"encoding/json"
	"time"
)

//Rate is a data strucrture for exchange_rate entities
type Rate struct {
	ID             int64        `gorm:"primary_key" json:"id"`
	CurrencyPair   CurrencyPair `gorm:"foreignkey:CurrencyPairID"` // use ProfileRefer as foreign key
	CurrencyPairID int64
	ExchangeRate   float64    `gorm:"column:exchange_rate" json:"exchange_rate"`
	ExchangeDate   time.Time  `gorm:"column:exchange_date" json:"exchange_date"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

//RatePayload is a data structure format for exchange_rate request and response api
type RatePayload struct {
	ExchangeRate    float64   `gorm:"column:exchange_rate" json:"exchange_rate"`
	ExchangeDate    time.Time `gorm:"column:exchange_date" json:"exchange_date"`
	BaseCurrency    string    `gorm:"column:base_currency" json:"base_currency"`
	CounterCurrency string    `gorm:"column:counter_currency" json:"counter_currency"`
	WeeklyAverage   float64   `gorm:"column:7day_avg" json:"weekly_average"`
}

//ExchangeData is a data structure format for exchange_rate request and response api
type ExchangeData struct {
	ExchangeRate float64   `gorm:"column:exchange_rate" json:"exchange_rate"`
	ExchangeDate time.Time `gorm:"column:exchange_date" json:"exchange_date"`
}

//UnmarshalJSON change data type from string to time
func (rp *RatePayload) UnmarshalJSON(b []byte) error {
	var dateFormat = "2006-01-02" //YYYY-MM-DD
	type Alias RatePayload
	aux := &struct {
		ExchangeDate string `json:"exchange_date"`
		*Alias
	}{
		Alias: (*Alias)(rp),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	t, err := time.Parse(dateFormat, aux.ExchangeDate)
	if err == nil {
		rp.ExchangeDate = t
	}

	return nil
}

//TableName get current table name
func (Rate) TableName() string {
	return "rate"
}
