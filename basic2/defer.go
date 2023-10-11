package main
import "fmt"
func main() {
	for i:=0 ; i<100 ; i++ {
		if i == 101 {
			break
		}else {
			fmt.Println(i)
		}
		fmt.Println(i)
	}
}