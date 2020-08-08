package store

import (
	"github.com/rmukubvu/pumpdata/model"
	"time"
)

func AddPump(p model.Pump) error {
	return insert(model.InsertPump, p.ToMap())
}

func AddPumpService(p model.PumpService) error {
	return insert(model.InsertPumpService, p.ToMap())
}

func AddPumpServiceHistory(p model.PumpServiceHistory) error {
	return insert(model.InsertPumpServiceHistory, p.ToMap())
}

func AddOrangeNotification(p model.SensorNotifications) error {
	return insert(model.InsertOrange, p.ToMap())
}

func AddRedNotification(p model.SensorNotifications) error {
	return insert(model.InsertRed, p.ToMap())
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

func AddPumpTests(p model.PumpTest) error {
	return insert(model.InsertPumpTest, p.ToMap())
}

func AddAnnunciator(p model.Annunciator) error {
	return insert(model.InsertAnnunciator, p.ToMap())
}

func PumpTestBySerial(serial string) ([]model.PumpTest, error) {
	var p []model.PumpTest
	err := db.Select(&p, model.SelectPumpTestWithSerial, serial)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SensorByTypeAndId(v model.Sensor) (model.Sensor, error) {
	p := model.Sensor{}
	err := db.Get(&p, model.SelectSensorByTypeAndPumpId, v.TypeId, v.PumpId)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SensorDataBySerialAndType(typeId int, serial string) ([]model.Sensor, error) {
	var p []model.Sensor
	err := db.Select(&p, model.SelectSensorByTypeAndSerial, typeId, serial)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SensorViewModelBySerialAndType(typeId int, serial string) ([]model.SensorViewModel, error) {
	var p []model.SensorViewModel
	err := db.Select(&p, model.SelectSensorViewModel, serial, typeId)
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

func PumpService(id string) (model.PumpService, error) {
	var p model.PumpService
	err := db.Get(&p, model.SelectPumpServiceWithId, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func PumpServiceHistory(id string) ([]model.PumpServiceHistory, error) {
	var p []model.PumpServiceHistory
	var err error
	if id == "-1" {
		err = db.Select(&p, model.SelectAllPumpServiceHistory)
	} else {
		err = db.Select(&p, model.SelectPumpServiceHistoryWithId, id)
	}
	if err != nil {
		return p, err
	}
	return p, nil
}

func AllPumpsDueForService() ([]model.PumpService, error) {
	var p []model.PumpService
	err := db.Select(&p, model.SelectAllPumpServices)
	if err != nil {
		return p, err
	}
	return p, nil
}

func AllAnnunciator() ([]model.Annunciator, error) {
	var p []model.Annunciator
	err := db.Select(&p, model.SelectAnnunciator)
	if err != nil {
		return p, err
	}
	return p, nil
}

func AnnunciatorBySerial(serial string) (model.Annunciator, error) {
	var p model.Annunciator
	err := db.Get(&p, model.SelectAnnunciatorBySerial, serial)
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

func GetOrangeNotification(date time.Time) ([]model.SensorNotifications, error) {
	var p []model.SensorNotifications
	err := db.Select(&p, model.SelectOrangeByDate, date)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetRedNotification(date time.Time) ([]model.SensorNotifications, error) {
	var p []model.SensorNotifications
	err := db.Select(&p, model.SelectRedByDate, date)
	if err != nil {
		return p, err
	}
	return p, nil
}
