package model

import (
	"encoding/json"
)

type DigitalMessage struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

//https://stackoverflow.com/questions/41774046/enabling-intellijs-fancy-%E2%89%A0-not-equal-to-operator
func (p *DigitalMessage) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *DigitalMessage) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}
