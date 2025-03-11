package strategy

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(left, right int) int {
	return o.Operator.Apply(left, right)
}

type Addition struct{}

func (Addition) Apply(left, right int) int {
	return left + right
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main() {
	mult := Operation{Multiplication{}}
	mult.Operate(3, 5)

	add := Operation{Addition{}}
	add.Operate(3, 5)

}
