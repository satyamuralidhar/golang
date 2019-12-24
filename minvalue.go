//minumum value caluclate

package main

import ("fmt"
	"math"
)

func main() {

	a ,b := 250.2 , 320.2
        c := math.Min(a,b)
        fmt.Println("minimum value of  a and b is:" ,c)
}


