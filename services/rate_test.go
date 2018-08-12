package services

// func TestRate_Create(t *testing.T) {
// 	type fields struct {
// 		RateRepository repositories.RateInterface
// 	}
// 	type args struct {
// 		params models.Rate
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    models.Rate
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &Rate{
// 				RateRepository: tt.fields.RateRepository,
// 			}
// 			got, err := r.Create(tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Rate.Create() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Rate.Create() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRate_FindAll(t *testing.T) {
// 	type fields struct {
// 		RateRepository repositories.RateInterface
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []models.Rate
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &Rate{
// 				RateRepository: tt.fields.RateRepository,
// 			}
// 			got, err := r.FindAll()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Rate.FindAll() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Rate.FindAll() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRate_Update(t *testing.T) {
// 	type fields struct {
// 		RateRepository repositories.RateInterface
// 	}
// 	type args struct {
// 		id     int64
// 		params models.Rate
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    models.Rate
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &Rate{
// 				RateRepository: tt.fields.RateRepository,
// 			}
// 			got, err := r.Update(tt.args.id, tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Rate.Update() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Rate.Update() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRate_Delete(t *testing.T) {
// 	type fields struct {
// 		RateRepository repositories.RateInterface
// 	}
// 	type args struct {
// 		id int64
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &Rate{
// 				RateRepository: tt.fields.RateRepository,
// 			}
// 			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
// 				t.Errorf("Rate.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestRate_ExchangeDataByDate(t *testing.T) {
// 	type fields struct {
// 		RateRepository repositories.RateInterface
// 	}
// 	type args struct {
// 		date time.Time
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    []models.RatePayload
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &Rate{
// 				RateRepository: tt.fields.RateRepository,
// 			}
// 			got, err := r.ExchangeDataByDate(tt.args.date)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Rate.ExchangeDataByDate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Rate.ExchangeDataByDate() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRate_TrendDataByCurrency(t *testing.T) {
// 	type fields struct {
// 		RateRepository repositories.RateInterface
// 	}
// 	type args struct {
// 		base    string
// 		counter string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    []models.ExchangeData
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &Rate{
// 				RateRepository: tt.fields.RateRepository,
// 			}
// 			got, err := r.TrendDataByCurrency(tt.args.base, tt.args.counter)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Rate.TrendDataByCurrency() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Rate.TrendDataByCurrency() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
