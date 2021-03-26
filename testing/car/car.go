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

type builder interface {
	Paint(Color) builder
	RedColor() builder
	Wheels(Wheels) builder
	SportsWheels() builder
	SteelWheels() builder
	TopSpeed(Speed) builder
	SpeedMPH(top Speed) builder
	SpeedKPH(top Speed) builder
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

func (c car) RedColor() builder {
	c.color = RedColor
	return c
}

func (c car) SportsWheels() builder {
	c.wheels = SportsWheels
	return c
}

func (c car) SteelWheels() builder {
	c.wheels = SteelWheels
	return c
}

func (c car) SpeedMPH(top Speed) builder {
	c.speed = top * MPH
	return c
}

func (c car) SpeedKPH(top Speed) builder {
	c.speed = top * KPH
	return c
}

func (c car) Paint(color Color) builder {
	c.color = color
	return c
}

func (c car) Wheels(wheels Wheels) builder {
	c.wheels = wheels
	return c
}

func (c car) TopSpeed(speed Speed) builder {
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
	assembly := Builder().Paint(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}

func Builder() builder {
	return car{}
}
