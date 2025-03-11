package car_test

import (
	"github.com/anuchito/testing/car"
	"testing"
)

func TestBuildCar(t *testing.T) {
	assembly := car.Builder().RedColor()

	familyCar := assembly.SportsWheels().SpeedKPH(50).Build()
	familyCar.Drive()

	sportsCar := assembly.SteelWheels().SpeedMPH(150).Build()
	sportsCar.Drive()
}