//modifying slice
package main
import "fmt"
func main() {
	satya := [...]int{10,20,30,40,50,60,70,80,90,100}
	dslice := satya[2:5]
	fmt.Println("array before", satya)
	for i := range dslice{
		dslice[i]++
	}
	fmt.Println("array after",satya)
}