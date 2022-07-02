package elon

import "strconv"

// Drive drive the car
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance display the distance driven
func (car *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// DisplayBattery display the battery percentage
func (car *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// CanFinish check if a remote control car can finish a race
func (car *Car) CanFinish(trackDistance int) bool {
	return float64(car.battery*car.speed/car.batteryDrain) >= float64(trackDistance)
}
