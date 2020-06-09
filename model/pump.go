package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Pump struct {
	Id           int     `json:"id" db:"id"`
	TypeId       int     `json:"type_id" db:"type_id"`
	SerialNumber string  `json:"serial_number" db:"serial_number"`
	NickName     string  `json:"nick_name" db:"nick_name"`
	Lat          float32 `json:"lat" db:"lat"`
	Lng          float32 `json:"lng" db:"lng"`
	CreatedDate  int64   `json:"created_date" db:"created_date"`
}

const (
	InsertPump = `insert into pump (type_id,serial_number,nick_name,lat,lng,created_date) 
				  values (:type_id,:serial_number,:nick_name,:lat,:lng,:created_date)`
	SelectPumpWithSerial = `select * from pump where serial_number = $1`
)

func (p *Pump) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *Pump) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *Pump) Key() string {
	return fmt.Sprintf("pump-%s", p.SerialNumber)
}

func (p *Pump) ToMap() map[string]interface{} {
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
