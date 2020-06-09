package model

import (
	"encoding/json"
)

type PumpTypes struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

//https://stackoverflow.com/questions/41774046/enabling-intellijs-fancy-%E2%89%A0-not-equal-to-operator
func (p *PumpTypes) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *PumpTypes) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}

func (p *PumpTypes) ToMap() map[string]interface{} {
	//set the created date time stamp here
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
