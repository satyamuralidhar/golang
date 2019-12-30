package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("bye")

}
