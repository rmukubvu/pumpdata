package model

import (
	"encoding/json"
	"time"
)

type Annunciator struct {
	Id           int       `json:"id,omitempty" db:"id"`
	SiteName     string    `json:"site_name,omitempty" db:"site_name"`
	SerialNumber string    `json:"serial_number,omitempty" db:"serial_number"`
	Msisdn       string    `json:"msisdn" db:"msisdn"`
	CreatedDate  time.Time `json:"created_date,omitempty" db:"created_date"`
}

const (
	InsertAnnunciator = `INSERT INTO annunciator (site_name , serial_number, msisdn , created_date)
						 VALUES
						 (
							:site_name,
							:serial_number,
							:msisdn,
							 current_date
						 )
						ON CONFLICT ( serial_number )
									DO
									UPDATE
									SET msisdn = :msisdn , site_name = :site_name , created_date = current_date`
	SelectAnnunciator         = `select * from annunciator `
	SelectAnnunciatorBySerial = `select * from annunciator where serial_number = $1`
)

func (p *Annunciator) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *Annunciator) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *Annunciator) ToMap() map[string]interface{} {
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
