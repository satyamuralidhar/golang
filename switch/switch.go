package main

import "fmt"

func main() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("thumb")
	case 2:
		fmt.Println("index")
	case 3:
		fmt.Println("middle")
	case 4:
		fmt.Println("ring")
	case 5:
		fmt.Println("pinky")
	}
}
