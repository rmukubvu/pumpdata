package model

type SensorViewModel struct {
	PumpName     string `json:"pump_name,omitempty" db:"pump_name"`
	SensorName   string `json:"sensor_name,omitempty" db:"sensor_name"`
	SerialNumber string `json:"serial_number,omitempty" db:"serial_number"`
	Value        string `json:"s_value,omitempty" db:"s_value"`
	Date         string `json:"created_date,omitempty" db:"created_date"`
}
