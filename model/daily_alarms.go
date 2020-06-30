package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type DailyAlarms struct {
	Id           int    `json:"id,omitempty" db:"id"`
	SerialNumber string `json:"serial_number,omitempty" db:"serial_number"`
	TypeId       int    `json:"type_id,omitempty" db:"type_id"`
	CreatedDate  int    `json:"created_date,omitempty" db:"created_date"`
}

const (
	InsertDailyAlarms = `INSERT INTO daily_alarms (serial_number, type_id , created_date)
						VALUES
						(
							:serial_number,
							:type_id,
							:created_date
						)`
	SelectDailyAlarmsByDate = `select * from daily_alarms where created_date = $1`
)

func (p *DailyAlarms) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *DailyAlarms) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *DailyAlarms) Key() string {
	p.CreatedDate = getDate()
	return fmt.Sprintf("pump.daily.alarms-%d", p.CreatedDate)
}

func (p *DailyAlarms) ToMap() map[string]interface{} {
	//set the created date time stamp here
	p.CreatedDate = getDate()
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

func getDate() int {
	now := time.Now()
	year, month, day := now.Date()
	res := fmt.Sprintf("%d%d%d", year, int(month), day)
	date, _ := strconv.Atoi(res)
	return date
}
