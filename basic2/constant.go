package main 

import "fmt"

const (

	a = iota
	
	b = iota

	c = iota
	
	d = iota
)

const (

	k = iota

	l = iota
)

func main() { 

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d,"\n")
 
	fmt.Println(k)

	fmt.Println(l)

}

