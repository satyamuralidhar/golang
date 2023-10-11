package main

import "fmt"

func main() {
	a := 10 //1010
	b := 3  //0011

	fmt.Printf("%v\n", a&b)  //0010 (2) AND operator
	fmt.Printf("%v\n", a|b)  //0011 (11) OR operator
	fmt.Printf("%v\n", a^b)  // 1001(9) exclusive AND operator
	fmt.Printf("%v\n", a&^b) // 0110 (8) exclusive OR operator

	//shift operands
	//left shift
	x := 8     // 2^3
	x = x << 3 // 2^3 * 2^3 = 64
	fmt.Printf("%v\n", x)
	//right shift
	y := 8
	y = y >> 3 // 2^3 / 2^3 = 2^0 --> 1
	fmt.Printf("%v\n", y)
}
