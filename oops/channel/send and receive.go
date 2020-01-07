package main 
import "fmt"
func satya(ch chan int) {
	fmt.Println(230 + <-ch)
}
func main() {
	fmt.Println("starting main method")
	//var ch chan int
	ch := make(chan int)
	go satya(ch)	
	ch <- 23
	fmt.Println("ending main method")
	
}
