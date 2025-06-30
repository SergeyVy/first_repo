package main

import (
	"fmt"
	"errors"
	"math"
)

func secondMax(numbers []int) (int, error) {
	if len(numbers) < 2 {
		return 0, errors.New("в срезе элементов меньше двух")
	}

	max1 := math.MinInt
	max2 := math.MinInt

	for _, num := range numbers {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 && num != max1 {
			max2 = num
		}
	}

	if max2 == math.MinInt {
		return 0, errors.New("в срезе нет второго уникального максимума")
	}

	return max2, nil
}

func main() {
result, err := secondMax([]int{2, 5, 6, 8, 9})
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Второе максимальное число:", result)
	}
}
	
