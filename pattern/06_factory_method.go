package main

import "fmt"

type iVehicle interface {
	setName(string)
	getName() string
	setSpeed(int)
	getSpeed() int
}

type vehicle struct {
	name string
	speed int
}

func (v *vehicle) setName(name string) {
	v.name = name
}

func (v *vehicle) getName() string {
	return v.name
}

func (v *vehicle) setSpeed(speed int) {
	v.speed = speed
}

func (v *vehicle) getSpeed() int {
	return v.speed
}

type car struct {
	vehicle
}

func newCar(name string, speed int) iVehicle {
	return &car{
		vehicle {
			name: name + " car",	
			speed: speed,
		},
	}
}

type ship struct {
	vehicle
}

func newShip(name string, speed int) iVehicle {
	return &car{
		vehicle {
			name: name + " ship",	
			speed: speed,
		},
	}
}

func createVehicle(vecType, name string, speed int) iVehicle {
	if vecType == "car" {
		return newCar(name, speed)
	}

	if vecType == "ship" {
		return newShip(name, speed)
	}

	return nil
}

func main() {
	volvo := createVehicle("car", "Volvo", 90)
	titanic := createVehicle("ship", "Titanic", 30)
	
	fmt.Printf("Given the transport %s with speed %d\n", volvo.getName(), volvo.getSpeed())
	fmt.Printf("Given the transport %s with speed %d\n", titanic.getName(), titanic.getSpeed())
}
