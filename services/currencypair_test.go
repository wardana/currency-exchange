package services

import (
	"reflect"
	"testing"

	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/repositories"
	"github.com/wardana/currency-exchange/repositories/mocks"
)

func TestCurrencyPair_Create(t *testing.T) {

	type fields struct {
		CurrencyPairRepository repositories.CurrencyPairInterface
		RateRepository         repositories.RateInterface
	}
	type args struct {
		params models.CurrencyPair
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CurrencyPair
		wantErr bool
	}{
		{
			name:    "Positive test case",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{params: models.CurrencyPair{BaseCurrency: "IDR", CounterCurrency: "USD"}},
			wantErr: false,
		},
		{
			name:    "Duplicate currency exchange code",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{params: models.CurrencyPair{BaseCurrency: "IDR", CounterCurrency: "SGD"}},
			wantErr: true,
		},
		{
			name:    "Failed to store into db",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{params: models.CurrencyPair{BaseCurrency: "IDR", CounterCurrency: "IDR"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CurrencyPair{
				CurrencyPairRepository: tt.fields.CurrencyPairRepository,
				RateRepository:         tt.fields.RateRepository,
			}
			got, err := c.Create(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrencyPair.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyPair.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestCurrencyPair_FindAll(t *testing.T) {
// 	type fields struct {
// 		CurrencyPairRepository repositories.CurrencyPairInterface
// 		RateRepository         repositories.RateInterface
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []models.CurrencyPair
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Positive test case",
// 			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
// 			wantErr: false,
// 		},
// 		// {
// 		// 	name:    "Failed to get any data err at database",
// 		// 	fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
// 		// 	wantErr: true,
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &CurrencyPair{
// 				CurrencyPairRepository: tt.fields.CurrencyPairRepository,
// 				RateRepository:         tt.fields.RateRepository,
// 			}
// 			got, err := c.FindAll()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CurrencyPair.FindAll() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CurrencyPair.FindAll() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestCurrencyPair_FindOne(t *testing.T) {
	type fields struct {
		CurrencyPairRepository repositories.CurrencyPairInterface
		RateRepository         repositories.RateInterface
	}
	type args struct {
		params models.CurrencyPair
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CurrencyPair
		wantErr bool
	}{
		{
			name:    "Positive test case",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{params: models.CurrencyPair{BaseCurrency: "IDR", CounterCurrency: "SGD"}},
			wantErr: false,
		},
		{
			name:    "Can't found any data",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{params: models.CurrencyPair{BaseCurrency: "IDR", CounterCurrency: "USD"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CurrencyPair{
				CurrencyPairRepository: tt.fields.CurrencyPairRepository,
				RateRepository:         tt.fields.RateRepository,
			}
			got, err := c.FindOne(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrencyPair.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyPair.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyPair_Update(t *testing.T) {
	type fields struct {
		CurrencyPairRepository repositories.CurrencyPairInterface
		RateRepository         repositories.RateInterface
	}
	type args struct {
		id     int64
		params models.CurrencyPair
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CurrencyPair
		wantErr bool
	}{
		{
			name:    "Positive test case",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 1, params: models.CurrencyPair{}},
			wantErr: false,
		},
		{
			name:    "Database err",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 3, params: models.CurrencyPair{}},
			wantErr: true,
		},
		{
			name:    "Invalid id",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 3, params: models.CurrencyPair{BaseCurrency: "IDR", CounterCurrency: "GBR"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CurrencyPair{
				CurrencyPairRepository: tt.fields.CurrencyPairRepository,
				RateRepository:         tt.fields.RateRepository,
			}
			got, err := c.Update(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrencyPair.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyPair.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyPair_Delete(t *testing.T) {
	type fields struct {
		CurrencyPairRepository repositories.CurrencyPairInterface
		RateRepository         repositories.RateInterface
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Positive test case",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 1},
			wantErr: false,
		},
		{
			name:    "Can't found any data",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 1},
			wantErr: false,
		},
		{
			name:    "Can't found any data",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 2},
			wantErr: true,
		},
		{
			name:    "Failed to update current data",
			fields:  fields{CurrencyPairRepository: mocks.MockCurrencyPair{}, RateRepository: mocks.MockRateRepository{}},
			args:    args{id: 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CurrencyPair{
				CurrencyPairRepository: tt.fields.CurrencyPairRepository,
				RateRepository:         tt.fields.RateRepository,
			}
			if err := c.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CurrencyPair.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
