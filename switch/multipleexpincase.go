package main

import "fmt"

func main() {
	letter := "u"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not vowel")
	}
}
