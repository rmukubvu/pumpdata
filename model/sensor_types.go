package model

import (
	"encoding/json"
)

type SensorTypes struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name"  db:"name"`
	DisplayName  string `json:"display_name"  db:"display_name"`
	DefaultValue string `json:"default_value" db:"default_value"`
}

//https://stackoverflow.com/questions/41774046/enabling-intellijs-fancy-%E2%89%A0-not-equal-to-operator
func (p *SensorTypes) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *SensorTypes) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}
