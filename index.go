package main
import "fmt"
func main() {
	var a = []int{1,2,3,4,5,6} 
	for index,element := range a {
		fmt.Println(index,")",element,"\n")
	}
}	
