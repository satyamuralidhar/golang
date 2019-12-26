package main

import "fmt"

func main() {
	for i:=1 ;i<=15 ; i++ {
		if i>=6 {
			break
		}
		fmt.Printf("value of i is  %d \n",i)
	}
	fmt.Println("exit")
}
