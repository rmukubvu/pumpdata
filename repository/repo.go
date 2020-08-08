package repository

import (
	"encoding/json"
	"fmt"
	"github.com/rmukubvu/pumpdata/bus"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/notifications"
	"github.com/rmukubvu/pumpdata/sms"
	"github.com/rmukubvu/pumpdata/store"
	"strconv"
	"time"
)

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

	/*c, err := getCompanyByPumpId(p.PumpId)
	if err != nil {
		return err
	}*/

	var message = "ok"
	var smsMessage = ""
	siteName := pumpCache[serial].NickName
	sensorName := store.GetSensorName(p.TypeId)
	if i < res.MinValue {
		smsMessage = fmt.Sprintf("%s at %s\nvalue is below threshold of %d.", sensorName, siteName, res.MinValue)
		message = fmt.Sprintf("%s\nvalue is below threshold of %d.\n%s", siteName, res.MinValue, res.AlertMessage)
	}

	if i > res.MaxValue {
		smsMessage = fmt.Sprintf("%s at %s\nvalue is above threshold of %d.", sensorName, siteName, res.MaxValue)
		message = fmt.Sprintf("%s\nvalue is above threshold of %d.\n%s", siteName, res.MaxValue, res.AlertMessage)
	}

	//check if push to orange ...
	//x >= MIN && x <= MAX
	sn := model.SensorNotifications{
		SiteName:    siteName,
		SensorName:  sensorName,
		SensorValue: p.Value,
		Comment:     "",
		CreatedDate: time.Now(),
	}

	if i >= res.Orange && i <= res.Red {
		sn.Comment = "this is a warning notification"
		store.AddOrangeNotification(sn)
	}
	//check if push to red
	if i >= res.Red {
		sn.Comment = "this is a danger notification , attend to this"
		store.AddRedNotification(sn)
	}

	//dont send anything
	if message == "ok" {
		return nil
	}
	//trigger -- for dashboard
	go dailyAlarms(serial, p.TypeId)
	m := contacts[amakosiCompanyId] //only amakosi staff get the sms information
	/*rb.Trigger = rabbit.TriggerMessage{
		Message:      message,
		SerialNumber: serial,
		CompanyId:    c.CompanyId,
		TypeId:       p.TypeId,
		PumpId:       p.PumpId,
		Value:        p.Value,
		Contacts:     m,
		CreatedDate:  time.Now().String(),
	}*/
	//send to slack
	go SendSlack(smsMessage)
	//send an sms
	go sendSms(m, smsMessage)
	//do the trigger
	return nil
	//rb.TriggerAlarm()
}

func SendSlack(message string) error {
	var errMsg = "ok"
	err := notifications.SlackSend(message)
	if err != nil {
		errMsg = err.Error()
	}
	log := model.PumpLogs{
		Sender:  "amakosi-ops",
		Payload: message,
		Date:    time.Now().String(),
		Error:   errMsg,
	}
	go eb.Publish(logsChannel, "slack_logs", log)
	return err
}

func SendSms(messages []model.DigitalMessage) error {
	f := store.TwilioSmsConfig()
	c := sms.TwilioConfig{
		Sid:   f.Sid,
		Token: f.Token,
		From:  f.Number,
	}
	s := sms.New(c)
	logs := s.SendToMultiple(messages)
	for _, e := range logs {
		eb.Publish(logsChannel, "sms_logs", e)
	}
	return nil
}

func sendEmail(pump, serial, serviceDate string) {
	//emails are only sent to internal amakosi management
	m := contacts[amakosiCompanyId]
	for _, e := range m {
		res := notifications.SendEmail(e.EmailAddress, pump, serial, serviceDate)
		log := model.PumpLogs{
			Sender:  e.EmailAddress,
			Payload: serial + " due on " + serviceDate,
			Date:    time.Now().String(),
			Error:   res,
		}
		eb.Publish(logsChannel, "email_logs", log)
	}
}

func sendSms(contacts []model.SensorAlarmContact, msg string) {
	phones := make([]string, 0, cap(contacts))
	for _, contact := range contacts {
		phones = append(phones, contact.Cellphone)
	}
	f := store.TwilioSmsConfig()
	c := sms.TwilioConfig{
		Sid:   f.Sid,
		Token: f.Token,
		From:  f.Number,
	}
	s := sms.New(c)
	logs := s.SendToMultipleNumbers(phones, msg)
	for _, e := range logs {
		eb.Publish(logsChannel, "sms_logs", e)
	}
}

func getCompanyByPumpId(id int) (model.Company, error) {
	c := model.Company{PumpId: id}
	item, err := redisCache.Get(c.Key())
	if err != nil {
		c, err = store.CompanyById(id)
	} else {
		err = c.FromJson([]byte(item))
	}
	return c, err
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
	return redisCache.Set(pumpTypesKey, jsonItem)
}

func dequeue() {
	ch1 := make(chan bus.DataEvent)
	eb.Subscribe(logsChannel, ch1)
	for {
		select {
		case msg := <-ch1:
			lg.InsertRecord(msg)
		}
	}
}
