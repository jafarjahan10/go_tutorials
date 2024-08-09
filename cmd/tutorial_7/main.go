// Pointers

/*
* -> Pointers are special type.
* -> These are variables with stored memory location rather than the values like int, string etc.
 */

package main

import "fmt"

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is %p\n", &thing1)

	var result [5]float64 = square(&thing1)
	fmt.Printf("The result: %v\n", result)
	fmt.Printf("The value of thing1: %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
