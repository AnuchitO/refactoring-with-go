package decorator

import "log"

type Operator func(int) int

func LogDecorate(fn Operator) Operator {
	return func(n int) int {
		log.Println("start decorate:", n)

		result := fn(n)

		log.Println("completed decorate:", result)

		return result
	}
}

func Triple(n int) int {
	return n * 3
}

func main() {
	f := LogDecorate(Triple)

	f(5)
}
