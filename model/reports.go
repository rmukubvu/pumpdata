package model

import "time"

type ReportPaging struct {
	Page      int `json:"page"`
	PageLimit int `json:"limit"`
}

type SensorDataReportRequest struct {
	ReportPaging
	CompanyId int    `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type AlarmsDataReportRequest struct {
	ReportPaging
	SerialNumber string `json:"serial_number"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

type ReportResponse struct {
	Error        string      `json:"error"`
	Data         interface{} `json:"data,omitempty"`
	CurrentPage  int         `json:"current_page,omitempty"`
	TotalRecords int         `json:"total_records,omitempty"`
	TotalPages   int         `json:"total_pages,omitempty"`
}

type SensorDataReport struct {
	SerialNumber string    `json:"serial_number,omitempty"`
	PumpName     string    `json:"pump_name,omitempty"`
	PumpType     string    `json:"pump_type,omitempty"`
	SensorName   string    `json:"sensor_name,omitempty"`
	SensorValue  string    `json:"sensor_value,omitempty"`
	CreatedDate  time.Time `json:"created_date,omitempty"`
}

type ServiceReportRequest struct {
	ReportPaging
	SerialNumber string `json:"serial_number"`
}

type ServiceReportResponse struct {
	SerialNumber   string    `json:"serial_number,omitempty" db:"serial_number"`
	PumpName       string    `json:"pump_name,omitempty" db:"nick_name"`
	PumpType       string    `json:"pump_type,omitempty" db:"name"`
	ServiceComment string    `json:"service_comment,omitempty" db:"service_comment"`
	ServiceDate    time.Time `json:"service_date,omitempty" db:"service_date"`
}
