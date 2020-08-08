package model

import (
	"encoding/json"
	"time"
)

type PumpServiceHistory struct {
	Id             int       `json:"id" db:"id"`
	SerialNumber   string    `json:"pump_id" db:"pump_id"`
	ServiceDate    time.Time `json:"service_date" db:"service_date"`
	Serviced       int       `json:"serviced" db:"serviced"`
	ServiceComment string    `json:"service_comment" db:"service_comment"`
	CreatedDate    time.Time `json:"created_date" db:"created_date"`
}

const (
	InsertPumpServiceHistory = `insert into pump_service_history (pump_id, service_date, serviced, service_comment , created_date) 
								values 
									(	
											:pump_id,
											:service_date,
											:serviced,
											:service_comment,
											current_date
									)`
	SelectPumpServiceHistoryWithId = `select * from pump_service_history where pump_id = $1`
	SelectAllPumpServiceHistory    = `select * from pump_service_history order by service_date`
)

func (p *PumpServiceHistory) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *PumpServiceHistory) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *PumpServiceHistory) ToMap() map[string]interface{} {
	//set the created date time stamp here
	b, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil
	}
	return res
}
