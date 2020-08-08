package sensors

import (
	"context"
	"github.com/rmukubvu/pumpdata/model"
	"github.com/rmukubvu/pumpdata/repository"
	"github.com/rmukubvu/pumpdata/store"
	"log"
	"strconv"
	"time"
)

type Server struct {
}

func (s *Server) SendSensorInformation(ctx context.Context, req *SensorRequest) (*SensorResponse, error) {
	log.Printf("new message from pump with serial - %s", req.SerialNumber)
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

func (s *Server) CreatePump(ctx context.Context, req *PumpRequest) (*PumpResponse, error) {
	log.Printf("creating pump with serial number - %s", req.SerialNumber)
	var msg string
	p, err := buildPump(req)
	if err != nil {
		msg = err.Error()
	} else {
		err = repository.AddPump(p)
		if err == nil {
			msg = "pump data has been added to pipeline"
		} else {
			msg = err.Error()
		}
	}
	return &PumpResponse{Message: msg}, nil
}

func (s *Server) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	log.Printf("test ping message - %s", req.Data)
	return &PingResponse{
		Data: "PONG",
	}, nil
}

func buildPumpSensor(req *SensorRequest) (model.Sensor, error) {
	pump, err := repository.GetPumpBySerialNumber(req.SerialNumber, false)
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

func buildPump(req *PumpRequest) (model.Pump, error) {
	var lat float32 = 0.0
	var lon float32 = 0.0
	if lt, err := strconv.ParseFloat(req.Lat, 32); err == nil {
		lat = float32(lt)
	}

	if lg, err := strconv.ParseFloat(req.Lon, 32); err == nil {
		lon = float32(lg)
	}

	return model.Pump{
		TypeId:       1,
		SerialNumber: req.SerialNumber,
		Lat:          lat,
		Lng:          lon,
	}, nil
}
