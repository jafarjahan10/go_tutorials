package main

import "fmt"

func main() {
	// Array
	var intArr [3]int32 = [3]int32{110, 112, 123}

	fmt.Println(intArr)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slice
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)

	var intSlice2 []int32 = make([]int32, 3, 8)
	fmt.Println(len(intSlice2), cap(intSlice2))

	// Map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "John": 24, "Jane": 21}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["John"]
	if ok {
		fmt.Printf("Age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	delete(myMap2, "Jane")
	fmt.Println(myMap2)

	// For loop
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While loop workaround
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
}
