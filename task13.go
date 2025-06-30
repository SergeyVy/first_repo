package main

import(
	"fmt"
)

func increment(ptr *int){

	*ptr=*ptr+1
	
	
}

func main(){

	n:=5
	increment(&n)
	increment(&n)
	fmt.Println(n)
	

}