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
	sensorModel, err := buildPumpSensor(req)
	if err == nil {
		err = store.AddSensor(*sensorModel)
	}
	return &SensorResponse{Message: err.Error()}, nil
}

func buildPumpSensor(req *SensorRequest) (*model.Sensor, error) {
	pump, err := repository.GetPumpBySerialNumber(req.SerialNumber)
	if err != nil {
		return nil, err
	}
	sensorId, err := store.GetSensorId(req.SensorName)
	if err != nil {
		return nil, err
	}
	return &model.Sensor{
		PumpId:      pump.Id,
		TypeId:      sensorId,
		Value:       req.Value,
		CreatedDate: time.Now().Unix(),
	}, nil
}
