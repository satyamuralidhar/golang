package main

import "fmt"

func main() {
	for j := 1; j <= 15; j++ {
		if j >= 7 {
			continue
		}
		fmt.Printf("value of j is %d \n", j)
	}
	fmt.Println("Exit")
}
