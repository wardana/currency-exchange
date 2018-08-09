package repositories

import (
	"time"

	"github.com/wardana/currency-exchange/models"
)

type (
	//RateInterface is an interface for rate entities
	RateInterface interface {
		Create(params models.Rate) (models.Rate, error)
		Find(params *models.Rate) ([]models.Rate, error)
		Update(id int64, params models.Rate) (models.Rate, error)
		RemoveByPairID(id int64) error
		TrendDataByCurrency(base, counter string) ([]models.ExchangeData, error)
		ExchangeDataByDate(date time.Time) ([]models.RatePayload, error)
	}

	//CurrencyPairInterface is an interface for CurrencyPair entities
	CurrencyPairInterface interface {
		Create(params models.CurrencyPair) (models.CurrencyPair, error)
		Find(params *models.CurrencyPair) ([]models.CurrencyPair, error)
		Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error)
	}
)
