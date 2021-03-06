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
				  values 
					(	
							:type_id,
                          	:serial_number,
 							:nick_name,
							:lat,
							:lng,
							:created_date
					)
					ON CONFLICT ( serial_number )
									DO
									UPDATE
									SET type_id = :type_id , nick_name = :nick_name , lat = :lat , lng = :lng`
	SelectPumpWithSerial = `select * from pump where serial_number = $1`
	SelectAllPumps       = `select * from pump`
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

func (p *Pump) IdKey() string {
	return fmt.Sprintf("pump.%d", p.Id)
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
