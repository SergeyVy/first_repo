package main

import(
	"fmt"
)

func sumToN(n int)int{

sum:=0

for i := 0; i <= n; i++ {
sum+=i
}
return sum

}

func main(){

fmt.Println(sumToN(5))
fmt.Println(sumToN(10))

}