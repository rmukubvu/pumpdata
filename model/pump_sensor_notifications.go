package model

import (
	"encoding/json"
	"time"
)

type SensorNotifications struct {
	Id          int       `json:"id" db:"id"`
	SiteName    string    `json:"site_name" db:"site_name"`
	SensorName  string    `json:"sensor_name" db:"sensor_name"`
	SensorValue string    `json:"s_value" db:"s_value"`
	Comment     string    `json:"s_comment" db:"s_comment"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}

const (
	InsertOrange = `insert into orange_notification (site_name,sensor_name,s_value,s_comment,created_date) 
				    values (:site_name,:sensor_name,:s_value,:s_comment,:created_date)`
	InsertRed = `insert into red_notification (site_name,sensor_name,s_value,s_comment,created_date) 
				    values (:site_name,:sensor_name,:s_value,:s_comment,:created_date)`

	SelectOrangeByDate = `select * from orange_notification where created_date = $1`

	SelectRedByDate = `select * from red_notification where created_date = $1`
)

func (p *SensorNotifications) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SensorNotifications) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *SensorNotifications) ToMap() map[string]interface{} {
	p.CreatedDate = time.Now()
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
