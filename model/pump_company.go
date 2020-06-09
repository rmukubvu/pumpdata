package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Company struct {
	Id          int   `json:"id" db:"id"`
	PumpId      int   `json:"pump_id" db:"pump_id"`
	CompanyId   int   `json:"company_id" db:"company_id"`
	CreatedDate int64 `json:"created_date" db:"created_date"`
}

const (
	InsertPumpCompany = `insert into pump_company (pump_id,company_id,created_date) 
				  		 values (:pump_id,:company_id,:created_date)`
	SelectPumpCompanyWithId = `select * from pump_company where company_id = $1`
)

func (p *Company) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *Company) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *Company) Key() string {
	return fmt.Sprintf("pump.company-%d", p.CompanyId)
}

func (p *Company) ToMap() map[string]interface{} {
	//set the created date time stamp here
	p.CreatedDate = time.Now().Unix()
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
