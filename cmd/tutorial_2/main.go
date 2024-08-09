package main

import "fmt"

func main() {
	// Number types
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNumb32 int32 = 2
	var result32 float32 = floatNum32 + float32(intNumb32)
	fmt.Println(result32)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// String type
	var myString string = `Hello world`
	fmt.Println(myString)
	fmt.Println(len(`test`))
	var myRune rune = 'a'
	fmt.Println(myRune)

	// Boolean type
	var myBool bool = true
	fmt.Println(myBool)

	/*
	* The default value for int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64, rune 	types is 0
	* The default value for string type is ''
	* The default value for bool is false
	 */

	myVar := "hello"
	fmt.Println(myVar)

	var1, var2 := 1, 2

	fmt.Println(var1, var2)
}
