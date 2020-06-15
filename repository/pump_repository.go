package repository

import (
	"encoding/json"
	"github.com/rmukubvu/pumpdata/cache"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/store"
)

var cacheService *cache.Service

const (
	pumpTypesKey = "pump.types"
)

func init() {
	cacheService = cache.New()
}

func AddPump(p model.Pump) error {
	return store.AddPump(p)
}

func AddCompany(p model.Company) error {
	err := store.AddCompany(p)
	if err == nil {
		cacheService.Set(p.Key(), p.ToJson())
	}
	return err
}

func AddSensor(p model.Sensor) error {
	err := store.AddSensor(p)
	if err == nil {
		cacheService.Set(p.Key(), p.ToJson())
	}
	return err
}

func GetCompanyById(id int) (model.Company, error) {
	p := model.Company{CompanyId: id}
	item, err := cacheService.Get(p.Key())
	if err != nil {
		return store.CompanyById(id)
	}
	err = p.FromJson([]byte(item))
	return p, err
}

func SensorByTypeAndId(typeId, pumpId int) (model.Sensor, error) {
	p := model.Sensor{TypeId: typeId, PumpId: pumpId}
	item, err := cacheService.Get(p.Key())
	if err != nil {
		return store.SensorByTypeAndId(p)
	}
	err = p.FromJson([]byte(item))
	return p, err
}

func GetPumpBySerialNumber(serialNumber string) (model.Pump, error) {
	p := model.Pump{SerialNumber: serialNumber}
	item, err := cacheService.Get(p.Key())
	if err != nil {
		p, err = store.GetPumpBySerialNumber(serialNumber)
		//cache it
		cacheService.Set(p.Key(), p.ToJson())
		//done
		return p, err
	}
	err = p.FromJson([]byte(item))
	return p, err
}

func AddPumpType(p model.PumpTypes) error {
	err := store.AddPumpType(p)
	if err == nil {
		go func() {
			cacheTypes()
		}()
	}
	return err
}

func FetchPumpTypes() ([]model.PumpTypes, error) {
	items, err := cacheService.Get(pumpTypesKey)
	if err != nil {
		return store.FetchPumpTypes()
	}
	var p []model.PumpTypes
	err = json.Unmarshal([]byte(items), &p)
	return p, err
}

func FetchSensorTypes() ([]model.SensorTypes, error) {
	return store.GetSensorTypes()
}

func cacheTypes() error {
	items, err := store.FetchPumpTypes()
	if err != nil {
		return err
	}
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	jsonItem := string(b)
	return cacheService.Set(pumpTypesKey, jsonItem)
}
