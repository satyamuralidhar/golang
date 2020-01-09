package main 
import(
	"fmt"
	"time"
)
func task() {
	for range time.Tick(time.Second * 1) {
		fmt.Println(" satya ")
	}
}	
func main() {
	go task()
	time.Sleep(time.Second * 6)
}	
