package store

import (
	"github.com/rmukubvu/pumpdata/model"
)

func AddPump(p model.Pump) error {
	return insert(model.InsertPump, p.ToMap())
}

func AddSensorAlarm(p model.SensorAlarm) error {
	return insert(model.InsertSensorAlarm, p.ToMap())
}

func AddCompany(p model.Company) error {
	return insert(model.InsertPumpCompany, p.ToMap())
}

func AddSensor(p model.Sensor) error {
	return insert(model.InsertSensor, p.ToMap())
}

func AddDailyAlarm(p model.DailyAlarms) error {
	return insert(model.InsertDailyAlarms, p.ToMap())
}

func AddSensorAlarmContact(p model.SensorAlarmContact) error {
	return insert(model.InsertAlarmContact, p.ToMap())
}

func AddSensorData(p model.SensorData) error {
	return insert(model.InsertSensorData, p.ToMap())
}

func AddSensorDataRaw(p model.SensorData) error {
	return insert(model.InsertSensorData, p.ToMap())
}

func SensorByTypeAndId(v model.Sensor) (model.Sensor, error) {
	p := model.Sensor{}
	err := db.Get(&p, model.SelectSensorByTypeAndPumpId, v.TypeId, v.PumpId)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SensorDataBySerialNumber(v model.SensorData) ([]model.SensorData, error) {
	var p []model.SensorData
	err := db.Select(&p, model.SelectSensorDataBySerialNumber, v.SerialNumber)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SensorDataBySerialNumberAndId(serial string, id int) (model.SensorData, error) {
	p := model.SensorData{}
	err := db.Get(&p, model.SelectSensorDataBySerialNumberAndId, serial, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetPumpBySerialNumber(serialNumber string) (model.Pump, error) {
	p := model.Pump{}
	err := db.Get(&p, model.SelectPumpWithSerial, serialNumber)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetAlarmContacts() ([]model.SensorAlarmContact, error) {
	var p []model.SensorAlarmContact
	err := db.Select(&p, model.SelectAllAlarmsContacts)
	if err != nil {
		return p, err
	}
	return p, nil
}

func CompanyById(id int) (model.Company, error) {
	p := model.Company{}
	err := db.Get(&p, model.SelectPumpCompanyWithPumpId, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetAllPumps() ([]model.Pump, error) {
	var p []model.Pump
	err := db.Select(&p, model.SelectAllPumps)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetSensorData() ([]model.SensorData, error) {
	var p []model.SensorData
	err := db.Select(&p, model.SelectAllSensorData)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetAllAlarms() ([]model.SensorAlarm, error) {
	var p []model.SensorAlarm
	err := db.Select(&p, model.SelectAllAlarms)
	if err != nil {
		return p, err
	}
	return p, nil
}

func PumpsUnderCompany(id int) ([]model.Pump, error) {
	var p []model.Pump
	err := db.Select(&p, model.SelectPumpsByCompanyId, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func DailyAlarms() ([]model.DailyAlarms, error) {
	var p []model.DailyAlarms
	date := GetCreatedDate()
	err := db.Select(&p, model.SelectDailyAlarmsByDate, date)
	if err != nil {
		return p, err
	}
	return p, nil
}
