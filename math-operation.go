package main

import "fmt"

func mathOperation() {
	var num1 int8 = 1
	var num2 int32 = 2

	fmt.Println(num1 + int8(num2))

	// augmented assigment
	var i = 10
	i += 5
	fmt.Println(i)

	// unary operator
	// ++
	// --
	// -
	// +
	// !

	var a = 2
	a++
	fmt.Println(a)
}
