package main

import (
	"errors"
	"fmt"
)

func printMe(printValue string) {
	fmt.Println("Hello", printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder = numerator % denominator

	return result, remainder, err
}

func main() {
	printMe("Jahan")

	var result, remainder, err = intDivision(11, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	}
}
