package store

import "github.com/rmukubvu/pumpdata/model"

func AddPump(p model.Pump) error {
	return insert(model.InsertPump, p.ToMap())
}

func GetPumpBySerialNumber(serialNumber string) (model.Pump, error) {
	p := model.Pump{}
	err := db.Get(&p, model.SelectPumpWithSerial, serialNumber)
	if err != nil {
		return p, err
	}
	return p, nil
}

func AddCompany(p model.Company) error {
	return insert(model.InsertPumpCompany, p.ToMap())
}

func CompanyById(id int) (model.Company, error) {
	p := model.Company{}
	err := db.Get(&p, model.SelectPumpCompanyWithId, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func AddSensor(p model.Sensor) error {
	return insert(model.InsertSensor, p.ToMap())
}

func SensorByTypeAndId(v model.Sensor) (model.Sensor, error) {
	p := model.Sensor{}
	err := db.Get(&p, model.SelectSensorByTypeAndPumpId, v.TypeId, v.PumpId)
	if err != nil {
		return p, err
	}
	return p, nil
}

func SensorByPumpId(v model.Sensor) (model.Sensor, error) {
	p := model.Sensor{}
	err := db.Get(&p, model.SelectSensorByTypeAndPumpId, v.TypeId, v.PumpId)
	if err != nil {
		return p, err
	}
	return p, nil
}
