package main

import(
	"fmt"
	"errors"
)

func colculator(a,b int,op string)(int,error){

	switch op{
	case "+":return a+b,nil
	case "-":return a-b,nil
	case "*":return a*b,nil
	case"/":
		if b==0{
		return 0,errors.New("Число не делится на ноль")
	}
	return a/b,nil
default:
	return 0,errors.New("не известный оператор")

}
}
func main(){

	res,err:=colculator(10,5,"/")
	if err!=nil{
		fmt.Println("Ошибка",err)
	}else{
		fmt.Println("результат:",res)
	}
}