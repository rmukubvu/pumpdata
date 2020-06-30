package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type SensorData struct {
	Id           int    `json:"id,omitempty" db:"id"`
	SerialNumber string `json:"serial_number,omitempty" db:"serial_number"`
	TypeId       int    `json:"type_id,omitempty" db:"type_id"`
	TypeText     string `json:"type_text" db:"type_text"`
	Value        string `json:"s_value" db:"s_value"`
	UpdateDate   int64  `json:"update_date,omitempty" db:"update_date"`
}

type SensorWithAlarms struct {
	S SensorData  `json:"sensor"`
	A SensorAlarm `json:"alarm"`
}

const (
	InsertSensorData = `INSERT INTO sensor_data (serial_number, type_id , s_value , type_text , update_date)
						VALUES
						(
							:serial_number,
							:type_id,
							:s_value , 
							:type_text,
							:update_date
						)
						ON CONFLICT (serial_number, type_id )
							DO
								UPDATE
						SET s_value = :s_value , update_date = :update_date , type_text = :type_text`
	SelectSensorDataBySerialNumber      = `select serial_number,type_id,type_text,s_value,update_date from sensor_data where serial_number = $1`
	SelectSensorDataBySerialNumberAndId = `select serial_number,type_id,type_text,s_value,update_date from sensor_data where serial_number = $1 and type_id = $2`
	SelectAllSensorData                 = `select * from sensor_data`
)

func (p *SensorData) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SensorData) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *SensorData) Key() string {
	return fmt.Sprintf("pump.sensor.data-%s", p.SerialNumber)
}

func (p *SensorData) ToMap() map[string]interface{} {
	//set the created date time stamp here
	p.UpdateDate = time.Now().Unix()
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
