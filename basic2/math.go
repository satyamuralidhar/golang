package main 
import (
	"fmt"
	"math"
)
func main() {
	var num float64 = 12
	var result = math.Sqrt(num)
	fmt.Println(result)
	fmt.Printf("the values is %.2f", result)
}