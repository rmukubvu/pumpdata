package rabbit

import (
	"encoding/json"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/streadway/amqp"
	"log"
)

type QueueService struct {
	Trigger TriggerMessage
	R       Rbmq
}

type Rbmq struct {
	ch   *amqp.Channel
	q    amqp.Queue
	conn *amqp.Connection
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

func New() *QueueService {
	qs := &QueueService{}
	var err error
	qs.R.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	qs.R.ch, err = qs.R.conn.Channel()
	failOnError(err, "Failed to open a channel")

	qs.R.q, err = qs.R.ch.QueueDeclare(
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

func (p *QueueService) TriggerAlarm() error {
	err := p.R.ch.Publish(
		"",         // exchange
		p.R.q.Name, // routing key
		false,      // mandatory
		false,      // immediate
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
