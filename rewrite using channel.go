package main 
import ("fmt"
		"time"
)
func sample() {
	fmt.Println("hi")
}
func main() {
	go sample()
	time.Sleep(2 * time.Second)
	fmt.Println("exit")

}

//rewrite above program using channnel

package main
import "fmt"
func sample(done chan bool) {
	fmt.Println("hi")
	done <- true

}
func main() {
	done := make(chan bool)
	go sample(done)
	<- done
	fmt.Println(exit)

}