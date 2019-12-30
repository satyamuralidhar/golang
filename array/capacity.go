package main 
import "fmt"
func main() {
	x := [...]string{"satya","mohan","pavan","subbu","vijay"}
	xslice := x[2:4]
	fmt.Printf("x length is %d and capacity is %d",len(xslice),cap(xslice))

}