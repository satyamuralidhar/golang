//type and size

package main

import ("fmt"
	"unsafe"
)

func main() {

	var a , b int  = 63 , 64

	fmt.Println("value of a is:" ,a, "value of b is:" ,b)

	fmt.Printf("type of a is %T , size of a is %d" ,a, unsafe.Sizeof(a))
        
	fmt.Printf("type of b is %T , size of b is %d" , b, unsafe.Sizeof(b))


}
