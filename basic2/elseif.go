package main 

import "fmt"

func main() {
	num :=51
	if num <=51 {
		fmt.Printf("num is less than %d \n" , num)
	}else if num >= 51 && num <=100 {
		fmt.Printf("num is greater than %d  \n" , num) 
	}else {
		fmt.Println("num is greater than 100")
	}
}
