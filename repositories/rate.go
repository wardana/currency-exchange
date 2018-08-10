package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wardana/currency-exchange/models"
)

//Rate initialize exchange rate class
type Rate struct {
	DB *gorm.DB
}

//TrendDataByCurrency get 7 latest currency rate data using currency code
func (r *Rate) TrendDataByCurrency(base, counter string) ([]models.ExchangeData, error) {
	data := []models.ExchangeData{}
	query := `select r.exchange_date, r.exchange_rate from currency_pair cp
		left join rate r ON cp.id = r.currency_pair_id
		where cp.base_currency = ? and cp.counter_currency = ? ORDER BY r.exchange_date DESC LIMIT 7`

	if err := r.DB.Raw(query, base, counter).Scan(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

//ExchangeDataByDate get historical data using specify date
func (r *Rate) ExchangeDataByDate(date time.Time) ([]models.RatePayload, error) {
	data := []models.RatePayload{}
	query := `select cp.base_currency, cp.counter_currency, r.exchange_date as "exchange_date", COALESCE(r.exchange_rate, 0) as "exchange_rate", a.7day_avg
		from currency_pair cp
		left join rate r on cp.id = r.currency_pair_id and r.exchange_date = ?
		left join (
			select cp.id, AVG(r.exchange_rate) as  "7day_avg"
			from currency_pair cp
			join rate r on cp.id = r.currency_pair_id
			where r.deleted_at is null and cp.deleted_at is null
			and date(r.exchange_date) between (? - interval 7 day) and current_date
			group by 1
		) a on a.id = cp.id
		where r.deleted_at is null and cp.deleted_at is null
		`

	if err := r.DB.Raw(query, date, date).Scan(&data).Error; err != nil {
		return data, err
	}
	return data, nil

}

//Create is a function to create new record
func (r *Rate) Create(params models.Rate) (models.Rate, error) {
	if err := r.DB.Create(&params).Error; err != nil {
		return params, err
	}
	return params, nil
}

//Find is a function to search rate data using currency code
func (r *Rate) Find(params *models.Rate) ([]models.Rate, error) {
	data := []models.Rate{}
	if err := r.DB.Where(params).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

//Update is a function to update currency data
func (r *Rate) Update(id int64, params models.Rate) (models.Rate, error) {
	data := models.Rate{
		ID: id,
	}
	if err := r.DB.Model(&data).UpdateColumns(&params).Error; err != nil {
		return data, err
	}
	return data, nil
}

//RemoveByPairID is a function for delete exchange rate data using it's currency pair id
func (r *Rate) RemoveByPairID(id int64) error {

	currentDate := time.Now()
	err := r.DB.Model(&models.Rate{}).UpdateColumns(models.Rate{DeletedAt: &currentDate}).Error
	if err != nil {
		return err
	}
	return nil
}
