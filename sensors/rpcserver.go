package sensors

import (
	"context"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/repository"
	"github.com/rmukubvu/pumpdata/store"
	"log"
	"time"
)

type Server struct {
}

func (s *Server) SendSensorInformation(ctx context.Context, req *SensorRequest) (*SensorResponse, error) {
	log.Printf("Receive message from pump with serial - %s", req.SerialNumber)
	//do the mappings here
	var msg string
	sensor, err := buildPumpSensor(req)
	if err == nil {
		err = repository.AddSensor(sensor)
		if err == nil {
			msg = "sensor data has been added to pipeline"
		} else {
			msg = err.Error()
		}
	} else {
		msg = err.Error()
	}
	return &SensorResponse{Message: msg}, nil
}

func (s *Server) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	log.Printf("Receive message from ping message - %s", req.Data)
	return &PingResponse{
		Data: "PONG",
	}, nil
}

func buildPumpSensor(req *SensorRequest) (model.Sensor, error) {
	pump, err := repository.GetPumpBySerialNumber(req.SerialNumber)
	if err != nil {
		return model.Sensor{}, err
	}
	sensorId, err := store.GetSensorId(req.SensorName)
	if err != nil {
		return model.Sensor{}, err
	}
	return model.Sensor{
		PumpId:      pump.Id,
		TypeId:      sensorId,
		Value:       req.Value,
		CreatedDate: time.Now().Unix(),
	}, nil
}
