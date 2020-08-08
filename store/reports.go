package store

import "github.com/rmukubvu/pumpdata/model"

func GenerateSensorDataReport(id int, startDate, endDate string) ([]model.SensorDataReport, error) {
	var p []model.SensorDataReport
	err := db.Select(&p, model.PumpSensorDataByCompany, id, startDate, endDate)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GenerateAlarmSensorReport(serial, startDate, endDate string) ([]model.SensorDataReport, error) {
	var p []model.SensorDataReport
	err := db.Select(&p, model.AlarmPerPump, serial, startDate, endDate)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GenerateServiceReport(serial string) ([]model.ServiceReportResponse, error) {
	var p []model.ServiceReportResponse
	err := db.Select(&p, model.ServiceReport, serial)
	if err != nil {
		return p, err
	}
	return p, nil
}
