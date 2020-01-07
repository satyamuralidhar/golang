package main
import "fmt" 
func mychan(ch chan int) {
	for v:=0 ; v<5 ; v++ {
		ch <- 100
	}
	close(ch)
}
func main() {
	c := make(chan int)
	//calling goroutine
	go mychan(c)
	for {
		res,ok := <-c 
		if ok == false {
		fmt.Println("closing channel",ok)
			break 
		}
		fmt.Println("open channel",res,ok)
	}
}