package main

import (
	"log"

	"github.com/JoelTeoGom/iot-kafka-system/pkg/kafka"
)

type SensorServiceInterface interface {
	sendData(any float64) error
}

type Sensor struct {
	DeviceID      string `json:"device_id"`
	SensorService SensorServiceInterface
}

type SensorDataHumidity struct {
	Humidity  float64 `json:"humidity"`
	Timestamp int64   `json:"timestamp"`
}

type SensorDataTemperature struct {
	Temperature float64 `json:"temperature"`
	Timestamp   int64   `json:"timestamp"`
}

type SensorDataPressure struct {
	Pressure  float64 `json:"temperature"`
	Timestamp int64   `json:"timestamp"`
}

func main() {

	//special port, faster conections
	brokers := []string{"broker1:6969"}

	sensors := []Sensor{
		Sensor{DeviceID: "s1", SensorService: &SensorDataHumidity{}},
		Sensor{DeviceID: "s2"},
		Sensor{DeviceID: "s3"}}

	producer, err := kafka.NewProducer(brokers)

	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	for s, _ := range sensors {
		go s.SensorService.sendData()
	}

	producer.Close()

}

func (s *SensorDataHumidity) sendData() error {

}

func (s *SensorDataPressure) sendData() error {

}

func (s *SensorDataHumidity) sendData() error {

}
