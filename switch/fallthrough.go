package main

import "fmt"

func number() int {
	num := 15 * 10
	return num
}
func main() {
	switch num := number(); {
	case num <= 50:
		fmt.Printf("%d less than 50 \n", num)
		fallthrough
	case num >= 100:
		fmt.Printf("%d greater than 100 \n", num)
		fallthrough
	case num <= 200:
		fmt.Printf("%d less tahn 200 \n", num)
		fallthrough
	default:
		fmt.Println("Exit")
	}
}
