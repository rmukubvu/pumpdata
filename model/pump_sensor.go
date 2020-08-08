package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Sensor struct {
	Id           int       `json:"id" db:"id"`
	TypeId       int       `json:"type_id" db:"type_id"`
	PumpId       int       `json:"pump_id" db:"pump_id"`
	SerialNumber string    `json:"serial_number" db:"serial_number"`
	Value        string    `json:"s_value" db:"s_value"`
	CreatedDate  int64     `json:"created_date" db:"created_date"`
	UpdatedDate  time.Time `json:"updated_date" db:"updated_date"`
}

const (
	InsertSensor = `insert into pump_sensor (type_id,pump_id,serial_number,s_value,created_date) 
				    values (:type_id,:pump_id,:serial_number,:s_value,:created_date)`
	SelectSensorByTypeAndPumpId = `select * from pump_sensor where type_id = $1 and pump_id = $2`
	SelectSensorByTypeAndSerial = `select * from pump_sensor where type_id = $1 and serial_number = $2`
	SelectSensorViewModel       = `	select
								p.nick_name as pump_name,
								REPLACE(st.name,'_',' ') as sensor_name,
								ps.serial_number,
								ps.s_value,
								TO_CHAR(ps.updated_date :: DATE, 'dd/mm/yyyy') as created_date
								from pump p
								inner join pump_sensor ps on p.id = ps.pump_id
								inner join sensor_type st on ps.type_id = st.id
								where ps.serial_number = $1
								and ps.type_id = $2
								`
)

func (p *Sensor) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *Sensor) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *Sensor) Key() string {
	return fmt.Sprintf("pump.sensor-%d-%d", p.TypeId, p.PumpId)
}

func (p *Sensor) ToMap() map[string]interface{} {
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
