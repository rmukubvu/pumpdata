package store

import "github.com/rmukubvu/pumpdata/model"

func FetchPumpTypes() ([]model.PumpTypes,error){
	rows , err := db.Query("select * from pump_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rs := make([]model.PumpTypes, 0)
	for rows.Next() {
		var t model.PumpTypes
		if err := rows.StructScan(&t); err != nil {
			return nil, err
		}
		rs = append(rs, t)
	}
	return rs, nil
}
