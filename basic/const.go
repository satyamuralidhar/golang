package main

import "fmt"

const x int = 10 // package level constant

/*
note: when ever we using same const in local and pkg level it will pickup local only.
adding : int16+int64 wont work
*/

func main() {
	const x int = 20 // local level constant
	//const() is enumerated constants iota will give o/p like 0,1,2,3 series
	const (
		a = iota
		b = iota
		c = iota
		d = iota
	)

	fmt.Println(x)
	fmt.Println("Enumerated constant: ", a)
	fmt.Println("Enumerated constant: ", b)
	fmt.Println("Enumerated constant: ", c)
	fmt.Println("Enumerated constant: ", d)

	/*
		o/p:
		=====
			Enumerated constant:  0
			Enumerated constant:  1
			Enumerated constant:  2
			Enumerated constant:  3
	*/

	//task 2:

	const (
		_ = iota
		series1
		series2
		series3
	)

	fmt.Printf("%v\n", series2) //o/p is 2
	fmt.Printf("%v\n", series1) //o/p is 1

	//task 3:
	/*
		_ will remove the 0th  of name iota kaniko1 = 1 value iota +5 ==> 1+5
	*/
	const (
		_ = iota + 5
		kaniko1
		kaniko2
		kaniko3
	)

	fmt.Printf("%v\n", kaniko1)

}
