package main

import(
	"fmt"
)

func charCase(r rune)string{
	switch{
	case r>='A' && r<='Z':
	return"Заглавая латинская буква"
    case r>='z' && r<='z':
	return"Не заглавная латинская буква"
	case r>='А' && r<='Я':
	return "Заглавная кирилическая буква"
	case r>='а' && r<='я':
	return "Заглавная кирилическая буква"
	default:
		return "Другое"
}

}

func main(){

	fmt.Println(charCase('A'))
	fmt.Println(charCase('я'))
	fmt.Println(charCase('Б'))
	fmt.Println(charCase('#'))
}