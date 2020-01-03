package main 
import (
	"fmt"
)
func fullname(firstname * string , lastname * string) {
	if firstname == nil {
		panic("firstname not nil")
		
	}
	if lastname == nil {
		panic("last name not nil")
	
	}	
	fmt.Printf("%s ,%s", *firstname , *lastname)
}
func main()  {
	firstname := "satya"
	lastname := "peddireddi"
	fullname(&firstname , &lastname)
	fmt.Println("\n exit")
}
