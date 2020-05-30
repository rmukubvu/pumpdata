package model

import (
	"encoding/json"
)

type PumpTypes struct {
	Id int 		`json:"id" db:"id"`
	Name string	`json:"name" db:"name"`
}

//https://stackoverflow.com/questions/41774046/enabling-intellijs-fancy-%E2%89%A0-not-equal-to-operator
func (p PumpTypes) toJson() string {
	b , err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func fromJson(body string) (PumpTypes,error) {
	p := PumpTypes{}
	err := json.Unmarshal([]byte(body), &p)
	return p, err
}