// Generics

package main

import "fmt"

type gasEngine struct {
	gallon float32
	mpg    float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallon: 12.4,
			mpg:    25.6,
		},
	}

	var ev = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model Y",
		engine: electricEngine{
			kwh:   42.6,
			mpkwh: 12.5,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(ev)
}
