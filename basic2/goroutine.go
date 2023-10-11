package main 
import ("fmt"
		"time")
func satya(){
	fmt.Println("hi satya")
}
func main() {
	go satya()
	time.Sleep(1 * time.Second)
	fmt.Println("how are you")
}