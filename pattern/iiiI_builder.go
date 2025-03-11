package car

import "fmt"

type Color string

const (
	_          Color = ""
	RedColor         = "red"
	GreenColor       = "green"
	BlueColor        = "blue"
)

type Wheels string

const (
	_            Wheels = ""
	SportsWheels        = "sports"
	SteelWheels         = "steel"
)

type Speed float64

const (
	_   Speed = 0
	MPH       = 1
	KPH       = 1.60934
)

type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c car) Paint(color Color) Builder {
	c.color = color
	return c
}

func (c car) Wheels(wheels Wheels) Builder {
	c.wheels = wheels
	return c
}

func (c car) TopSpeed(speed Speed) Builder {
	c.speed = speed
	return c
}

func (c car) Drive() error {
	fmt.Println("Driving...")
	return nil
}
func (c car) Stop() error {
	fmt.Println("Breaking..")
	return nil
}

func (c car) Build() Interface {
	return c
}

func main() {
	assembly := CarBuilder().Paint(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}

func CarBuilder() Builder {
	return car{}
}
