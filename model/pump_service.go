package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type PumpService struct {
	Id              int       `json:"id" db:"id"`
	SerialNumber    string    `json:"pump_id" db:"pump_id"`
	LastServiceDate time.Time `json:"last_service_date" db:"last_service_date"`
	NextServiceDate time.Time `json:"next_service_date" db:"next_service_date"`
	CreatedDate     time.Time `json:"updated_date" db:"updated_date"`
}

const (
	InsertPumpService = `insert into pump_service (pump_id, last_service_date, next_service_date, updated_date) 
				  values 
					(	
							:pump_id,
                          	:last_service_date,
 							:next_service_date,
							current_date
					)
					ON CONFLICT ( pump_id )
									DO
									UPDATE
									SET last_service_date = :last_service_date , next_service_date = :next_service_date , updated_date = current_date`
	SelectPumpServiceWithId = `select * from pump_service where pump_id = $1`
	SelectAllPumpServices   = `select * from pump_service`
)

func (p *PumpService) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *PumpService) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *PumpService) Key() string {
	return fmt.Sprintf("pump.service.%s", p.SerialNumber)
}

func (p *PumpService) ToMap() map[string]interface{} {
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
