package main

import (
	"fmt"
	)

func main()  {
	fmt.Println("Hello World!")
	fmt.Println()

	/*
	TYPES
	string
	bool
	int int8 int16 int 32 int64
	uint uint8 ... uint64 uint ptr (Unsigned integer, can't be negative)
	byte - alias for uint8
	rune - alias for int 32
	float32 float64
	complex64 complex128
	 */

	var x int = 3 //Static Typing
	y := 7.5 //Dynamic Typing

	const z = 3.14 //Constant, like final but can't be reassigned.
	sum := float64(x) * y * z  //Casting

	fmt.Println(sum)
	fmt.Printf("%T\n", sum) //Formats to show type

	sum2 := 0

	if sum2 == 0 {
		fmt.Println("test")
	}
}