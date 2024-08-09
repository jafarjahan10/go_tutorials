// Structs and Interfaces
/*
* Struct is a way of defining our own type
* Struct also have concept of method which are functions directly tied to the struct and has access to the instance of the struct
 */

package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it.")
	} else {
		fmt.Println("Need to fuel up first.")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Printf("MPG: %v\n", myEngine.mpg)
	fmt.Printf("GALLONS: %v\n", myEngine.gallons)

	// Anonymous struct. This is not recommended
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{30, 20}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// Struct methods
	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 250)

	var myEvEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myEvEngine, 50)

}
