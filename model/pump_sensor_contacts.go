package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type SensorAlarmContact struct {
	Id           int    `json:"id" db:"id"`
	CompanyId    int    `json:"company_id" db:"company_id"`
	Cellphone    string `json:"cell_phone" db:"cell_phone"`
	EmailAddress string `json:"email_address" db:"email_address"`
	CreatedDate  int64  `json:"created_date" db:"created_date"`
}

const (
	InsertAlarmContact = ` INSERT INTO sensor_alarms_contacts (company_id , cell_phone , email_address , created_date)
									VALUES
									(
										:company_id,
										:cell_phone,
										:email_address , 
										:created_date
									)`
	SelectAllAlarmsContacts = `select * from sensor_alarms_contacts where company_id = 1`
)

func (p *SensorAlarmContact) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SensorAlarmContact) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *SensorAlarmContact) Key() string {
	return fmt.Sprintf("pump.sensor.contacts-%d", p.CompanyId)
}

func (p *SensorAlarmContact) ToMap() map[string]interface{} {
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
