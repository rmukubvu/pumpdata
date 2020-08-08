package rabbit

import (
	"encoding/json"
	"fmt"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/streadway/amqp"
	"log"
)

type QueueService struct {
	session
	Trigger TriggerMessage
}

type session struct {
	*amqp.Channel
	amqp.Queue
	*amqp.Connection
}

type TriggerMessage struct {
	Message      string                     `json:"message"`
	SerialNumber string                     `json:"serial_number"`
	CompanyId    int                        `json:"company_id"`
	TypeId       int                        `json:"type_id"`
	PumpId       int                        `json:"pump_id"`
	Value        string                     `json:"s_value"`
	Contacts     []model.SensorAlarmContact `json:"contacts"`
	CreatedDate  string                     `json:"created_date"`
}

const queueName = "amakhosi.alarms"

//xoxb-1215817235969-1200780626695-Me0HuIvp450n52dn1UcFdwEk
func New(url string) *QueueService {
	qs := &QueueService{}
	var err error
	qs.Connection, err = amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	qs.Channel, err = qs.Connection.Channel()
	failOnError(err, "Failed to open a channel")

	qs.Queue, err = qs.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return qs
}

func (p *QueueService) SystemIs(message string) error {
	err := p.Channel.Publish(
		"",                       // exchange
		"pump.service.available", // routing key
		false,                    // mandatory
		false,                    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}

func (p *QueueService) RemoteCalls(r model.Remote) (model.RemoteResponse, error) {
	message := fmt.Sprintf("bagder|%s|%s\r\n", r.SerialNumber, r.Message)
	return model.RemoteResponse{
		SerialNumber: r.SerialNumber,
		Message:      "queued for processing",
	}, p.SystemIs(message)
}

func (p *QueueService) TriggerAlarm() error {
	err := p.Channel.Publish(
		"",     // exchange
		p.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(p.ToJson()),
		})
	return err
}

func (p *QueueService) ToJson() string {
	b, err := json.Marshal(p.Trigger)
	if err != nil {
		return ""
	}
	return string(b)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Close tears the connection down, taking the channel with it.
func (p *QueueService) Close() error {
	if p.Connection == nil {
		return nil
	}
	return p.Connection.Close()
}
