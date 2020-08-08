package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ReneKroon/ttlcache"
	"github.com/rmukubvu/pumpdata/bus"
	"github.com/rmukubvu/pumpdata/cache"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/nosql"
	"github.com/rmukubvu/pumpdata/store"
	"time"
)

var (
	redisCache               *cache.Service
	alarm                    = make(map[int]model.SensorAlarm)
	contacts                 = make(map[int][]model.SensorAlarmContact)
	sensorCache              = make(map[string]model.SensorData)
	pumpCache                = make(map[string]model.Pump)
	eb                       *bus.EventBus
	lg                       *nosql.LogStore
	serviceCache             *ttlcache.Cache
	invalidSerialNumberError = errors.New("invalid serial number")
)

const (
	pumpTypesKey     = "pump.types"
	amakosiCompanyId = 1
	logsChannel      = "pump.services.logs"
)

func init() {
	//expiry cache
	serviceCache = ttlcache.NewCache()
	//set the expiry call back
	serviceCache.SetExpirationCallback(expirationCallback)
	//redis cache
	redisCache = cache.New()
	//cache the alarms
	alarms, _ := GetAllSensorAlarms()
	for _, e := range alarms {
		alarm[e.TypeId] = e
	}
	eb = bus.New()
	//cache pumps
	//cache the alarms
	allPumps, _ := GetAllPumps()
	for _, e := range allPumps {
		pumpCache[e.SerialNumber] = e
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
	//mongo db
	lg = nosql.NewConnection(store.MongoUrlConfig().Url)
	//for logging sms data
	go dequeue()
}

// if redis is down ,,, will be a problem
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
	//add annunciator data
	a := model.Annunciator{
		SiteName:     p.NickName,
		SerialNumber: p.SerialNumber,
		Msisdn:       "",
		CreatedDate:  time.Now(),
	}
	//run in background
	go AddAnnunciator(a)
	//continue
	_, err = GetPumpBySerialNumber(p.SerialNumber, true)
	return err
}

func AddCompany(p model.Company) error {
	err := store.AddCompany(p)
	if err == nil {
		redisCache.Set(p.Key(), p.ToJson())
	}
	return err
}

func AddAnnunciator(p model.Annunciator) error {
	return store.AddAnnunciator(p)
}

// do something if database is offline - save these messages somewhere .. i can
// see badger here
func AddSensor(p model.Sensor) error {
	err := store.AddSensor(p)
	if err == nil {
		//add sensor data
		pump := model.Pump{
			Id: p.PumpId,
		}
		serial, err := redisCache.Get(pump.IdKey())
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
		redisCache.Set(p.Key(), p.ToJson())
	}
	return err
}

func AddPumpService(p model.PumpService) error {
	err := store.AddPumpService(p)
	if err != nil {
		return err
	}
	d := p.NextServiceDate.Sub(p.LastServiceDate)
	msg := pumpServiceMessage(p.SerialNumber, p.NextServiceDate)
	serviceExpiryWithNotify(p.Key(), msg, d)
	return nil
}

func AddPumpServiceHistory(p model.PumpServiceHistory) error {
	return store.AddPumpServiceHistory(p)
}

func AddSensorAlarm(p model.SensorAlarm) error {
	err := store.AddSensorAlarm(p)
	if err == nil {
		redisCache.Set(p.Key(), p.ToJson())
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

func AddPumpTests(p model.PumpTest) error {
	return store.AddPumpTests(p)
}

func PumpTestBySerial(serial string) ([]model.PumpTest, error) {
	return store.PumpTestBySerial(serial)
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

func GetOrangeNotification(date time.Time) ([]model.SensorNotifications, error) {
	return store.GetOrangeNotification(date)
}

func GetRedNotification(date time.Time) ([]model.SensorNotifications, error) {
	return store.GetRedNotification(date)
}

func GetWaterTankLevelForSerial(serial string) (model.WaterTankLevels, error) {
	m := model.WaterTankLevels{}
	var p []model.SensorData
	arr := [4]int{10, 11, 12, 13}
	tanks := 0
	for _, id := range arr {
		key := fmt.Sprintf("sensor.%s.%d", serial, id)
		level := sensorCache[key]
		if level != (model.SensorData{}) {
			tanks++
			p = append(p, level)
		}
	}
	m.Data = p
	m.Tanks = tanks
	return m, nil
}

func GetAllPumps() ([]model.Pump, error) {
	return store.GetAllPumps()
}

func GetAllSensorAlarms() ([]model.SensorAlarm, error) {
	return store.GetAllAlarms()
}

func PumpService(serial string) (model.PumpService, error) {
	return store.PumpService(serial)
}

func GetAnnunciators() ([]model.Annunciator, error) {
	return store.AllAnnunciator()
}

func GetAnnunciatorBySerial(serial string) (model.Annunciator, error) {
	return store.AnnunciatorBySerial(serial)
}

func GetSensorDataBySerial(serial string) ([]model.SensorData, error) {
	numberOfSensors := store.GetNumberOfSensors()
	slice := make([]model.SensorData, 0, numberOfSensors)
	for i := 1; i <= numberOfSensors; i++ {
		key := fmt.Sprintf("sensor.%s.%d", serial, i)
		found := sensorCache[key]
		if found == (model.SensorData{}) {
			found = createMissingSensorValues(serial, i)
		}
		slice = append(slice, found)
	}
	return slice, nil
}

func GetCompanyById(id int) (model.Company, error) {
	p := model.Company{CompanyId: id}
	item, err := redisCache.Get(p.Key())
	if err != nil {
		return store.CompanyById(id)
	}
	err = p.FromJson([]byte(item))
	return p, err
}

func SensorByTypeAndId(typeId, pumpId int) (model.Sensor, error) {
	p := model.Sensor{TypeId: typeId, PumpId: pumpId}
	item, err := redisCache.Get(p.Key())
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
		pumpCache[serialNumber] = p
		redisCache.Set(p.Key(), p.ToJson())
		redisCache.Set(p.IdKey(), serialNumber)
		//done
		return p, err
	}

	item, err := redisCache.Get(p.Key())
	if err != nil {
		p, err = store.GetPumpBySerialNumber(serialNumber)
		//cache it
		pumpCache[serialNumber] = p
		redisCache.Set(p.Key(), p.ToJson())
		redisCache.Set(p.IdKey(), serialNumber)
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

func GetPumpsUnderCompany(id int) ([]model.Pump, error) {
	return store.PumpsUnderCompany(id)
}

func GetPumpsServices() ([]model.PumpService, error) {
	return store.AllPumpsDueForService()
}

func FetchPumpTypes() ([]model.PumpTypes, error) {
	items, err := redisCache.Get(pumpTypesKey)
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

func SensorDataBySerialAndType(id int, serial string) ([]model.Sensor, error) {
	return store.SensorDataBySerialAndType(id, serial)
}

func SensorViewModelBySerialAndType(typeId int, serial string) ([]model.SensorViewModel, error) {
	return store.SensorViewModelBySerialAndType(typeId, serial)
}

func GetServiceHistoryForPump(serialNumber string) ([]model.PumpServiceHistory, error) {
	return store.PumpServiceHistory(serialNumber)
}

func pumpServiceMessage(serial string, serviceDate time.Time) string {
	p := pumpCache[serial]
	//create the cache key
	due := serviceDate.Format("January 2, 2006")
	//cache it
	return fmt.Sprintf("%s|%s|%s", p.NickName, serial, due)
}

func validateSerialNumber(serial string) error {
	if pumpCache[serial] == (model.Pump{}) {
		return invalidSerialNumberError
	}
	return nil
}

func TearDown() {
	store.CloseDBConnection()
	serviceCache.Close()
}
