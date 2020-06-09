package store

import "github.com/rmukubvu/pumpdata/model"

func AddPumpType(p model.PumpTypes) error {
	return insert("insert into pump_type (name) values (:name)", p.ToMap())
}

func FetchPumpTypes() ([]model.PumpTypes, error) {
	var p []model.PumpTypes
	err := db.Select(&p, "select * from pump_type")
	if err != nil {
		return nil, err
	}
	return p, nil
}
