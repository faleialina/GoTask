package main

import (
	"fmt"
)

// "strings"

// "bufio"

// func verification(a int) string {
// 	if a%2 == 0 {
// 		return "Четное"
// 	} else {
// 		return "Нечетное"
// 	}
// }
// func verification(a []int) int {
// 	sum := 0
// 	for _, el := range a {
// 		sum = sum + el
// 	}
// 	return sum
// }

// func verification(a []float64) (float64, float64) {
// 	max := math.Inf(-1)
// 	min := math.Inf(1)
// 	for _, el := range a {
// 		if el > max {
// 			max = el
// 		}
// 		if el < min {
// 			min = el
// 		}
// 	}
// 	return max, min
// }

// func verification(num int) string {
// 	switch num {
// 	case 1:
// 		return "Понедельник"

// 	case 2:
// 		return "Вторник"

// 	case 3:
// 		return "Среда"

// 	case 4:
// 		return "Четверг"

// 	case 5:
// 		return "Пятница"

// 	case 6:
// 		return "Суббота"

// 	default:
// 		return "Воскресенье"
// 	}

// }

// func verification(sl []int) ([]int, []int) {
// 	var odd []int
// 	var even []int
// 	for _, el := range sl {
// 		if el%2 == 0 {
// 			odd = append(odd, el)
// 		} else {
// 			even = append(even, el)
// 		}
// 	}
// 	return odd, even
// }

// func verification(sl []int) []int {
// 	var newSl []int

// 	for _, el := range sl {
// 		if el != 0 {
// 			newSl = append(newSl, el)
// 		}
// 	}
// 	return newSl
// }

func verification(sl []int, num int) bool {
	var res bool

	for _, el := range sl {
		if el == num {
			res = true
		}
	}
	return res
}

func main() {

	// Напишите функцию, которая принимает целое число и возвращает строку "Четное" или "Нечетное", в зависимости от того, является ли число чётным или нечётным.
	// fmt.Println(verification(5))

	// Напишите функцию, которая принимает слайс целых чисел и возвращает сумму всех его элементов.
	// fmt.Println(verification([]int{1, 2, 3, 4}))

	// Напишите функцию, которая принимает слайс целых чисел и возвращает максимальное и минимальное число в этом слайсе.
	// fmt.Println(verification([]float64{1, 2, 3, 4}))

	// Напишите функцию, которая принимает число от 1 до 7 (представляющее день недели) и возвращает строку, например "Понедельник", "Вторник" и т. д. Используйте оператор switch.
	// fmt.Println(verification(5))

	// Напишите функцию, которая принимает слайс целых чисел и разделяет его на два слайса: один с чётными числами, другой с нечётными.
	// fmt.Println(verification([]int{1, 2, 3, 4}))

	// Напишите функцию, которая принимает слайс целых чисел и удаляет все нули из него.
	// fmt.Println(verification([]int{1, 2, 3, 4, 0}))
	// Напишите функцию, которая принимает слайс целых чисел и число, и проверяет, содержится ли это число в слайсе.
	fmt.Println(verification([]int{1, 2, 3, 4, 0}, 0))

}
