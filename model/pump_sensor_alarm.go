package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type SensorAlarm struct {
	Id           int    `json:"id,omitempty" db:"id"`
	TypeId       int    `json:"type_id,omitempty" db:"type_id"`
	MinValue     int    `json:"min_value,omitempty" db:"min_value"`
	MaxValue     int    `json:"max_value" db:"max_value"`
	AlertMessage string `json:"alert_message" db:"alert_message"`
	CreatedDate  int64  `json:"created_date,omitempty" db:"created_date"`
}

const (
	InsertSensorAlarm = `INSERT INTO sensor_alarms (type_id , min_value , max_value , alert_message , created_date)
								VALUES
								(
									:type_id,
									:min_value,
									:max_value , 
									:alert_message,
									:created_date
								)
								ON CONFLICT ( type_id )
									DO
									UPDATE
									SET min_value = :min_value , max_value = :max_value , alert_message = :alert_message`
	SelectAllAlarms = `select * from sensor_alarms order by type_id asc`
)

func (p *SensorAlarm) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SensorAlarm) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *SensorAlarm) Key() string {
	return fmt.Sprintf("pump.sensor.alarm-%d", p.TypeId)
}

func (p *SensorAlarm) ToMap() map[string]interface{} {
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
