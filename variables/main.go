package main

import "fmt"

func main() {
	var firstName string = "venkata" //Explicit type

	// Implicit type and most used declaration
	secondName := "pavan"
	age := 26
	height := 5.2
	isMajor := true
	fmt.Printf("My full name is %s %s\nAge is %d\nheight is %f\nmajor : %t\n\n", firstName, secondName, age, height, isMajor)

	// Default Values 
	// we can declare or initialize the type here but we dob't assign its value because its goes down to 0 values
	var myString string
	var myBool bool
	var myInteger int
	var myfloat float64

	fmt.Printf("Default values:\n\nstring is empty%s\nmy Boo is %t\nmy Integer is %d\nmy float : %f", myString, myBool, myInteger, myfloat)
	
}