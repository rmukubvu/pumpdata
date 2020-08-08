package model

import "encoding/json"

type Remote struct {
	SerialNumber string `json:"serial_number"`
	Message      string `json:"message"`
}

type RemoteResponse struct {
	SerialNumber string `json:"serial_number"`
	Message      string `json:"message"`
}

func (p *Remote) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}
