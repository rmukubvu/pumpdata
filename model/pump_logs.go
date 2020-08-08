package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PumpLogs struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Payload string             `json:"payload" bson:"payload,omitempty"`
	Sender  string             `json:"sender,omitempty" bson:"sender,omitempty"`
	Error   string             `json:"error,omitempty" bson:"error,omitempty"`
	Date    string             `json:"created_date" bson:"created_date,omitempty"`
}

func (p *PumpLogs) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (p *PumpLogs) FromJson(body []byte) error {
	err := json.Unmarshal(body, &p)
	return err
}
