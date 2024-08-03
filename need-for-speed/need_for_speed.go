package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	newCar := Car{
		speed:        car.speed,
		distance:     car.distance + car.speed,
		batteryDrain: car.batteryDrain,
		battery:      car.battery - car.batteryDrain,
	}
	if newCar.battery < 0 {
		return car
	}
	return newCar
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	travelDistance := track.distance - car.distance
	for i := 0; i <= travelDistance; i += car.speed {
		if car.battery < 0 {
			return false
		}
		car.battery -= car.batteryDrain
	}

	return true
}
