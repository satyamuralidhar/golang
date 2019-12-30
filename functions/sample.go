package main 
import ("fmt"
)
//func calc(x int , y int) (int ,int) {
	//var out1 = x + y
	//var out2 = x - y
	//return out1 , out2
//}
func calc(x , y int) (out1, out2 int) {
	out1 =  x + y
	out2 = x -y
	return 
}
func main() {
	num1 := 10
	num2 := 20
	result1,result2 := calc(num1 , num2)
	fmt.Println(result1 , result2)

}