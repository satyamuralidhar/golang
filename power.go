package main

import "fmt"

func main() {

	x :=8  //2^3

	fmt.Println(x >> 3) //2^3/2^3 = 2^0 = 1

	fmt.Println(x << 3) //2^3*2^3 = 2^6 = 64
}
