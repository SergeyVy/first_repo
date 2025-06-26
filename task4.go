package main

import (
	"fmt"
	// "math"
)


func maxOfThree(a,b,c int)int{

var max int
if a>b{
	max=a
}else{
	max=b
}

if c>max{
	return c
}
return max 

// maxAB:=math.Max(float64(a),float64(b))   //math.Max(float64()) эту библиотеку можно использовать только с типом float64
// maxABC:=math.Max(maxAB,float64(c))
// return int(maxABC)   //это приведение типа float64 в int, так как main ждет int

}


func main(){

	fmt.Println(maxOfThree(3,7,5))
	fmt.Println(maxOfThree(10,2,8))
}