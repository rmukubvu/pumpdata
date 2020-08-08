package model

import (
	"encoding/json"
	"time"
)

type PumpTest struct {
	Id           int       `json:"id" db:"id"`
	PumpName     string    `json:"site_name" db:"site_name"`
	SerialNumber string    `json:"pump_serial" db:"pump_serial"`
	TestMessage  string    `json:"test_msg" db:"test_msg"`
	TestDate     time.Time `json:"test_date" db:"test_date"`
	CreatedDate  time.Time `json:"created_date" db:"created_date"`
}

const (
	InsertPumpTest = `insert into pump_test_data (site_name , pump_serial, test_msg, test_date, created_date) 
					  values 
						(
							:site_name,
							:pump_serial,
							:test_msg,
							:test_date,
							current_date
						) `
	SelectPumpTestWithSerial = `select * from pump_test_data where pump_serial = $1`
	SelectAllPumpTests       = `select * from pump_test_data`
)

func (p *PumpTest) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *PumpTest) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *PumpTest) ToMap() map[string]interface{} {
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
