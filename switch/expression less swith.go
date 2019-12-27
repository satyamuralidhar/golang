package main

import "fmt"

func main() {
	num := 70
	switch {
	case num >= 10 && num <= 50:
		fmt.Println("its inbetween 10 to 50 range")
	case num >= 50 && num <= 80:
		fmt.Println("its inbetween 50 to 80 range")
	case num >= 101:
		fmt.Println("its greaterthan 100")
	}
}
