package repositories

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wardana/currency-exchange/models"
)

type (
	//Rate initialize exchange rate class
	Rate struct {
		DB *gorm.DB
	}
	//RateInterface is an interface for rate entities
	RateInterface interface {
		Create(params models.Rate) (models.Rate, error)
		Find(params *models.Rate) ([]models.Rate, error)
		Update(id int64, params models.Rate) (models.Rate, error)
		RemoveByPairID(id int64) error
		HistoricalDataByDate(date time.Time) (*[]models.RatePayload, error)
	}
)

//HistoricalDataByDate get historical data using specify date
func (r *Rate) HistoricalDataByDate(date time.Time) (*[]models.RatePayload, error) {

	data := &[]models.RatePayload{}
	query := `select cp.base_currency, cp.counter_currency, r.exchange_date "exchange_date", COALESCE(r.exchange_rate, 0) "exchange_rate", a.7day_avg
		from currency_pair cp
		left join rate r on cp.id = r.currency_pair_id and date(exchange_date) = ?
		left join (
			select cp.id, AVG(r.exchange_rate) "7day_avg"
			from currency_pair cp
			join rate r on cp.id = r.currency_pair_id 
			where r.deleted_at is null and cp.deleted_at is null
			and date(?) between (? - interval 7 day) and current_date
			group by 1
		) a on a.id = cp.id
		where r.deleted_at is null and cp.deleted_at is null
	`

	if errs := r.DB.Raw(query, date, date, date).Scan(&data).GetErrors(); len(errs) > 0 {
		fmt.Println(data)
		return data, errs[0]
	}
	return data, nil

}

//Create is a function to create new record
func (r *Rate) Create(params models.Rate) (models.Rate, error) {
	if errs := r.DB.Create(&params).GetErrors(); len(errs) > 0 {
		return params, errs[0]
	}
	return params, nil
}

//Find is a function to search rate data using currency code
func (r *Rate) Find(params *models.Rate) ([]models.Rate, error) {
	data := []models.Rate{}
	if errs := r.DB.Where(params).Find(&data).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}

//Update is a function to update currency data
func (r *Rate) Update(id int64, params models.Rate) (models.Rate, error) {
	data := models.Rate{
		ID: id,
	}
	if errs := r.DB.Model(&data).UpdateColumns(&params).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}

//RemoveByPairID is a function for delete exchange rate data using it's currency pair id
func (r *Rate) RemoveByPairID(id int64) error {

	currentDate := time.Now()
	if errs := r.DB.Model(&models.Rate{}).UpdateColumns(models.Rate{DeletedAt: &currentDate}).GetErrors(); len(errs) > 0 {
		return errs[0]
	}
	return nil
}
