package io

import (
	"SensorReceiveUploadGtk/models"
	"fmt"
	"math/rand"
	"time"
)

/*
Make sure that all functions from the receivedata.go
file are wrapped in goroutines.
*/

func ReceiveFromSensors(c chan models.Sensors) {
	_ = models.Sensors{}
	fmt.Println(len(c))
}

func ReceiveFromFakeSensors(c chan models.Sensors) {
	for {
		c <- models.Sensors{
			Pool:     randomValue(80.0, 83.0),
			Air:      randomValue(85.0, 87.0),
			Humidity: randomValue(12.0, 20.0),
		}

		time.Sleep(3 * time.Second)
	}
}

func randomValue(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
