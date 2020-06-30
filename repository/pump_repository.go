package repository

import (
	"encoding/json"
	"fmt"
	"github.com/rmukubvu/pumpdata/cache"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/rabbit"
	"github.com/rmukubvu/pumpdata/sms"
	"github.com/rmukubvu/pumpdata/store"
	"strconv"
	"time"
)

var (
	cacheService *cache.Service
	alarm        = make(map[int]model.SensorAlarm)
	contacts     = make(map[int][]model.SensorAlarmContact)
	sensorCache  = make(map[string]model.SensorData)
	rb           *rabbit.QueueService
)

const (
	pumpTypesKey     = "pump.types"
	amakosiCompanyId = 1
)

func init() {
	cacheService = cache.New()
	//initiate the rabbit mq
	rb = rabbit.New()
	//cache the alarms
	alarms, _ := GetAllSensorAlarms()
	for _, i := range alarms {
		alarm[i.TypeId] = i
	}
	//cache contacts
	alarmContacts, _ := store.GetAlarmContacts()
	for _, e := range alarmContacts {
		//get and append
		slice := contacts[e.CompanyId]
		slice = append(slice, e)
		contacts[e.CompanyId] = slice
	}
	//cache sensor data
	data, _ := store.GetSensorData()
	for _, e := range data {
		cacheSensorData(e)
	}
}

//if redis is down ,,, will be a problem
// see how to handle this in near future ..
// easy way is on startup if redis is unavailable panic
// if already running have a keepalive tick interval @ 30 seconds
// to check if redis is up
// investigate if there is any effect of
func AddPump(p model.Pump) error {
	err := store.AddPump(p)
	if err != nil {
		return err
	}
	_, err = GetPumpBySerialNumber(p.SerialNumber, true)
	return err
}

func AddCompany(p model.Company) error {
	err := store.AddCompany(p)
	if err == nil {
		cacheService.Set(p.Key(), p.ToJson())
	}
	return err
}

//do something if database is offline - save these messages somewhere .. i can
//see badger here
func AddSensor(p model.Sensor) error {
	err := store.AddSensor(p)
	if err == nil {
		//add sensor data
		pump := model.Pump{
			Id: p.PumpId,
		}
		serial, err := cacheService.Get(pump.IdKey())
		if err != nil {
			return err
		}
		//do trigger
		go triggerAlarm(p, serial)
		//do the sensor data
		p := model.SensorData{
			SerialNumber: serial,
			TypeText:     store.GetSensorName(p.TypeId),
			TypeId:       p.TypeId,
			Value:        p.Value,
			UpdateDate:   p.CreatedDate,
		}
		//cache it
		cacheSensorData(p)
		//go and save it
		go store.AddSensorData(p)
		//cache the rest
		cacheService.Set(p.Key(), p.ToJson())
	}
	return err
}

func AddSensorAlarm(p model.SensorAlarm) error {
	err := store.AddSensorAlarm(p)
	if err == nil {
		cacheService.Set(p.Key(), p.ToJson())
	}
	return err
}

func AddSensorAlarmContact(p model.SensorAlarmContact) error {
	err := store.AddSensorAlarmContact(p)
	if err == nil {
		//get and append
		slice := contacts[p.CompanyId]
		slice = append(slice, p)
		contacts[p.CompanyId] = slice
	}
	return err
}

func DashboardAlarms() ([]model.SensorWithAlarms, error) {
	todayAlarm, err := store.DailyAlarms()
	if err != nil {
		return nil, err
	}
	var p []model.SensorWithAlarms
	for _, e := range todayAlarm {
		s, _ := store.SensorDataBySerialNumberAndId(e.SerialNumber, e.TypeId)
		a := alarm[s.TypeId]
		m := model.SensorWithAlarms{
			S: s,
			A: a,
		}
		p = append(p, m)
	}
	return p, nil
}

func GetWaterTankLevelForSerial(serial string) (model.SensorData, error) {
	const waterTankTypeId = 9
	//return store.SensorDataBySerialNumberAndId(serial, waterTankTypeId)
	key := fmt.Sprintf("sensor.%s.%d", serial, waterTankTypeId)
	return sensorCache[key], nil
}

func GetAllPumps() ([]model.Pump, error) {
	return store.GetAllPumps()
}

func GetAllSensorAlarms() ([]model.SensorAlarm, error) {
	return store.GetAllAlarms()
}

func GetSensorDataBySerial(serial string) ([]model.SensorData, error) {
	last := store.GetNumberOfSensors()
	slice := make([]model.SensorData, 0, last)
	for i := 1; i <= last; i++ {
		key := fmt.Sprintf("sensor.%s.%d", serial, i)
		found := sensorCache[key]
		if found == (model.SensorData{}) {
			found = createMissingSensorValues(serial, i)
		}
		slice = append(slice, found)
	}
	return slice, nil
	/*return store.SensorDataBySerialNumber(model.SensorData{
		SerialNumber: serial,
	})*/
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

func GetPumpBySerialNumber(serialNumber string, refresh bool) (model.Pump, error) {
	p := model.Pump{SerialNumber: serialNumber}
	var err error
	if refresh {
		//do sensor creation
		go createDefaultSensorValues(serialNumber)
		//continue with other business
		p, err = store.GetPumpBySerialNumber(serialNumber)
		//cache it
		cacheService.Set(p.Key(), p.ToJson())
		cacheService.Set(p.IdKey(), serialNumber)
		//done
		return p, err
	}

	item, err := cacheService.Get(p.Key())
	if err != nil {
		p, err = store.GetPumpBySerialNumber(serialNumber)
		//cache it
		cacheService.Set(p.Key(), p.ToJson())
		cacheService.Set(p.IdKey(), serialNumber)
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

func GetAlarmContacts(id int) []model.SensorAlarmContact {
	return contacts[id]
}

func triggerAlarm(p model.Sensor, serial string) error {
	//convert value to int
	i, err := strconv.Atoi(p.Value)
	if err != nil {
		return err
	}
	//success
	res := alarm[p.TypeId]
	if res == (model.SensorAlarm{}) {
		return nil
	}

	c, err := getCompanyByPumpId(p.PumpId)
	if err != nil {
		return err
	}

	var message = "ok"
	if i < res.MinValue {
		message = fmt.Sprintf("value is below threshold of %d.\n%s", res.MinValue, res.AlertMessage)
	}

	if i > res.MaxValue {
		message = fmt.Sprintf("value is above threshold of %d.\n%s", res.MaxValue, res.AlertMessage)
	}

	//dont send anything
	if message == "ok" {
		return nil
	}

	//trigger -- for dashboard
	go dailyAlarms(serial, p.TypeId)
	m := contacts[amakosiCompanyId] //only amakosi staff get the sms information
	rb.Trigger = rabbit.TriggerMessage{
		Message:      message,
		SerialNumber: serial,
		CompanyId:    c.CompanyId,
		TypeId:       p.TypeId,
		PumpId:       p.PumpId,
		Value:        p.Value,
		Contacts:     m,
		CreatedDate:  time.Now().String(),
	}
	//send an sms
	go sendSms(m, message)
	//do the trigger
	return rb.TriggerAlarm()
}

func sendSms(contacts []model.SensorAlarmContact, msg string) {
	phones := make([]string, 0, cap(contacts))
	for _, contact := range contacts {
		phones = append(phones, contact.Cellphone)
	}
	sms.Send(phones, msg)
}

func getCompanyByPumpId(id int) (model.Company, error) {
	c := model.Company{PumpId: id}
	item, err := cacheService.Get(c.Key())
	if err != nil {
		c, err = store.CompanyById(id)
	} else {
		err = c.FromJson([]byte(item))
	}
	return c, err
}

func GetPumpsUnderCompany(id int) ([]model.Pump, error) {
	return store.PumpsUnderCompany(id)
}

func dailyAlarms(serialNumber string, typeId int) {
	store.AddDailyAlarm(model.DailyAlarms{
		SerialNumber: serialNumber,
		TypeId:       typeId,
		CreatedDate:  store.GetCreatedDate(),
	})
}

func createDefaultSensorValues(serial string) {
	s, err := store.GetSensorTypes()
	if err != nil {
		return
	}

	for _, e := range s {
		p := model.SensorData{
			SerialNumber: serial,
			TypeId:       e.Id,
			TypeText:     e.Name,
			Value:        e.DefaultValue,
			UpdateDate:   time.Now().Unix(),
		}
		_ = store.AddSensorDataRaw(p)
		//cache it
		cacheSensorData(p)
	}
}

func createMissingSensorValues(serial string, typeId int) model.SensorData {
	e := store.GetSensorTypeByTypeId(typeId)
	p := model.SensorData{
		SerialNumber: serial,
		TypeId:       e.Id,
		TypeText:     e.Name,
		Value:        e.DefaultValue,
		UpdateDate:   time.Now().Unix(),
	}
	_ = store.AddSensorDataRaw(p)
	//cache it
	cacheSensorData(p)
	//done
	return p
}

func cacheSensorData(m model.SensorData) {
	key := fmt.Sprintf("sensor.%s.%d", m.SerialNumber, m.TypeId)
	sensorCache[key] = m
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
