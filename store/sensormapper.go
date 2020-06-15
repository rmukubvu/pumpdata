package store

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/rmukubvu/pumpdata/model"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	mu               sync.Mutex //guards lastId
	lastId           = 0
	sensorMap        = make(map[int]string)
	sensorMapFlipped = make(map[string]int)
	sensorTypes      = make([]model.SensorTypes, 0)
	sensorConfigured = errors.New("this sensor has be added to this server")
)

const (
	sensorsFile = "sensors.txt"
)

func init() {
	readFileLineAndMap()
}

// return key if exists else zero
func GetSensorId(s string) (int, error) {
	return getSensor(s)
}

// return key if exists else zero
func GetSensorName(id int) string {
	return sensorMap[id]
}

// get sensors types
func GetSensorTypes() ([]model.SensorTypes, error) {
	return sensorTypes, nil
}

func getSensor(s string) (int, error) {
	ss := strings.ReplaceAll(s, " ", "_")
	id := sensorMapFlipped[ss]
	if id == 0 {
		//its new add to collection and return new id
		return addNewSensor(ss), sensorConfigured
	}
	return id, nil
}

func readFileLineAndMap() {
	file, err := os.Open(sensorsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		ss := split[1]
		update(ss)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addNewSensor(s string) int {
	update(s)
	//flush them to file
	line := fmt.Sprintf("\n%d:%s", lastId, s)
	appendToFile(line)
	return lastId
}

func appendToFile(line string) {
	f, err := os.OpenFile(sensorsFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(line); err != nil {
		log.Println(err)
	}
}

func update(s string) {
	mu.Lock()
	id := lastId + 1
	sensorMap[id] = s
	sensorMapFlipped[s] = id
	sensorTypes = append(sensorTypes, model.SensorTypes{
		Id:   id,
		Name: s,
	})
	lastId = id //no thread safe
	mu.Unlock()
}
