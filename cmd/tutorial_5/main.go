package main

import (
	"fmt"
	"strings"
)

func main() {
	// String
	var myString = "résumé" // é is a non asci character
	fmt.Println(myString)

	var indexed = myString[0] // the type is uint8
	fmt.Printf("%v %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of myString of string type is %v\n", len(myString))

	// Rune
	var myString2 = []rune("résumé")
	fmt.Println(myString2)

	var indexed2 = myString2[0]
	fmt.Printf("%v %T\n", indexed2, indexed2)

	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of myString2 of type rune is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	// String Methods

	// String Builder
	// -> String concatenation. But this is inefficient. In this method we are concatenating string into a new variable every time.
	var strSlice = []string{"j", "a", "h", "a", "n"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}

	fmt.Println(catStr)

	/*
	* String are immutable in go. Cannot modify them once created
	 */

	// -> String concatenation: The efficient way
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	fmt.Println(strBuilder.String())

}
