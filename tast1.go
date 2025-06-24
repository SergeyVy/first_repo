package main

import (
	"fmt"
	"errors"
)

func SecondMax(numbers []int) (int, error) {
	if len(numbers) < 2 {
		return 0, errors.New("в срезе меньше двух элементов")
	}

	var max, second int

	// Инициализируем max и second
	if numbers[0] > numbers[1] {
		max = numbers[0]
		second = numbers[1]
	} else {
		max = numbers[1]
		second = numbers[0]
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > max {
			second = max
			max = numbers[i]
		} else if numbers[i] > second && numbers[i] != max {
			second = numbers[i]
		}
	}

	return second, nil
}

func main() {
	numbers := []int{10, 20, 4, 45, 99}

	res, err := SecondMax(numbers)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Второе максимальное число в срезе:", res)
	}
}