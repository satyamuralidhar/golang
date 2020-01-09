package main 
import "fmt" 
func main() {
	channel := make(chan int, 4)
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	close(channel)
	for element := range channel {
		fmt.Println(element)
	}
}