package store

import (
	"github.com/rmukubvu/pumpdata/model"
)

func AddPump(p model.Pump) (int64, error) {
	return insertWithLastId(model.InsertPump, p.ToMap())
}

func AddSensorAlarm(p model.SensorAlarm) error {
	return insert(model.InsertSensorAlarm, p.ToMap())
}

func GetAllAlarms() ([]model.SensorAlarm, error) {
	var p []model.SensorAlarm
	err := db.Select(&p, model.SelectAllAlarms)
	if err != nil {
		return p, err
	}
	return p, nil
}

func AddCompany(p model.Company) error {
	return insert(model.InsertPumpCompany, p.ToMap())
}

func AddSensor(p model.Sensor) error {
	return insert(model.InsertSensor, p.ToMap())
}

func AddSensorAlarmContact(p model.SensorAlarmContact) error {
	return insert(model.InsertAlarmContact, p.ToMap())
}

func AddSensorData(p model.Sensor, serial, typeText string) error {
	sd := model.SensorData{
		SerialNumber: serial,
		TypeText:     typeText,
		TypeId:       p.TypeId,
		Value:        p.Value,
		UpdateDate:   p.CreatedDate,
	}
	return insert(model.InsertSensorData, sd.ToMap())
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

func GetPumpBySerialNumber(serialNumber string) (model.Pump, error) {
	p := model.Pump{}
	err := db.Get(&p, model.SelectPumpWithSerial, serialNumber)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetAlarmContactsByCompanyId(v model.SensorAlarmContact) ([]model.SensorAlarmContact, error) {
	var p []model.SensorAlarmContact
	err := db.Select(&p, model.SelectAllAlarmsContacts, v.CompanyId)
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
