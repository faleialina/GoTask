package main

import (
	"errors"
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

// func verification(sl []int, num int) bool {
// 	var res bool

// 	for _, el := range sl {
// 		if el == num {
// 			res = true
// 		}
// 	}
// 	return res
// }

// func calculate(a, b float64, operation string) float64 {
// 	switch operation {
// 	case "+":
// 		return a + b
// 	case "-":
// 		return a - b
// 	case "*":
// 		return a * b
// 	case "/":
// 		return a / b
// 	default:
// 		return 0
// 	}
// }

// func average_value(sl []int) int {
// 	var sum int

// 	for _, item := range sl {
// 		sum += item
// 	}
// 	return sum / len(sl)
// }

// func removeDuplicates(nums []int) []int {

// 	seen := make(map[int]bool)
// 	result := []int{}

// 	for _, num := range nums {
// 		if !seen[num] {
// 			seen[num] = true
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }
// func processArray(numbers []int) []int {
// 	uniqueNumbers := removeDuplicates(numbers)
// 	if len(uniqueNumbers) == 0 {
// 		fmt.Println("Массив пуст после удаления дубликатов")
// 	}
// 	return uniqueNumbers
// }

func sumNumbers(slice []int) (int, []int, error) {
	if len(slice) == 0 {
		return 0, nil, errors.New("передан пустой слайс")
	}

	sum := 0
	indices := []int{}

	for i, num := range slice {
		if num%2 != 0 {
			sum += num
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return 0, nil, errors.New("в слайсе нет нечётных чисел")
	}

	return sum, indices, nil
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
	// fmt.Println(verification([]int{1, 2, 3, 4, 0}, 0))
	// 	8. Напишите функцию, которая принимает два числа и операцию (например, +, -, *, /) и выполняет
	// соответствующую арифметическую операцию.

	// fmt.Println("10 + 5 =", calculate(10, 5, "+"))
	// fmt.Println("10 - 5 =", calculate(10, 5, "-"))
	// fmt.Println("10 * 5 =", calculate(10, 5, "*"))
	// fmt.Println("10 / 5 =", calculate(10, 5, "/"))
	// fmt.Println("10 / 0 =", calculate(10, 0, "/"))

	// 9. Напишите функцию, которая принимает слайс целых чисел и возвращает среднее арифметическое этих чисел.

	// fmt.Println(average_value([]int{1, 2, 3, 4, 5}))
	// 10. Создайте функцию, которая принимает массив чисел, находит все повторяющиеся элементы и
	// удаляет их. После этого функция должна проверить, не является ли массив пустым, и вывести
	// соответствующее сообщение.
	// fmt.Println("Итоговый результат:", processArray([]int{1, 2, 2, 3, 4, 4, 5}))

	// 11. Разработайте программу, которая находит сумму всех нечётных чисел в слайсе и выводит их индексы.
	testCases := [][]int{
		{1, 2, 3, 4, 5},  // нормальный случай
		{2, 4, 6, 8},     // нет нечётных чисел
		{},               // пустой слайс
		{11, 22, 33, 44}, // другой нормальный случай
	}

	for _, tc := range testCases {
		sum, indices, err := sumNumbers(tc)
		if err != nil {
			fmt.Printf("Ошибка для слайса %v: %v\n", tc, err)
			continue
		}

		fmt.Printf("Слайс: %v\n", tc)
		fmt.Printf("Сумма нечётных чисел: %d\n", sum)
		fmt.Printf("Индексы нечётных чисел: %v\n", indices)
		fmt.Println("-----")
	}

	// 12. Напишите программу, которая проверяет, является ли слайс чисел палиндромом, то есть
	// читается ли слайс одинаково в обоих направлениях. В случае палиндрома программа должна
	// вывести true, иначе false.
	// 13. Напишите программу, которая находит сумму всех чисел в слайсе, которые больше среднего
	// значения
	// 14. Напишите программу, которая генерирует два случайных слайса чисел от 1 до 100 и находит
	// пересечение этих слайсов (элементы, которые встречаются в обоих слайсах).
	// 15. Напишите программу, которая генерирует слайс из N случайных чисел от 1 до 100, затем
	// создает два слайса: один с чисел, делящихся на 5, а другой — на 7.

}
