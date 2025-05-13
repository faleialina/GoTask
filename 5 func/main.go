package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

// func sumNumbers(slice []int) (int, []int, error) {
// 	if len(slice) == 0 {
// 		return 0, nil, errors.New("передан пустой слайс")
// 	}

// 	sum := 0
// 	indices := []int{}

// 	for i, num := range slice {
// 		if num%2 != 0 {
// 			sum += num
// 			indices = append(indices, i)
// 		}
// 	}

// 	if len(indices) == 0 {
// 		return 0, nil, errors.New("в слайсе нет нечётных чисел")
// 	}

// 	return sum, indices, nil
// }

// func isPalindromeSlice(slice []int) bool {
// 	if len(slice) <= 1 {
// 		return true
// 	}

// 	for i := 0; i < len(slice)/2; i++ {
// 		if slice[i] != slice[len(slice)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func sumAboveAverage(slice []int) (int, error) {

// 	if len(slice) == 0 {
// 		return 0, errors.New("передан пустой слайс")
// 	}

// 	sum := 0
// 	for _, num := range slice {
// 		sum += num
// 	}
// 	average := float64(sum) / float64(len(slice))

// 	result := 0
// 	for _, num := range slice {
// 		if float64(num) > average {
// 			result += num
// 		}
// 	}

// 	return result, nil
// }

func generateRandomSlice(size int) ([]int, error) {
	if size <= 0 {
		return nil, errors.New("размер слайса должен быть положительным числом")
	}

	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) + 1
	}
	return slice, nil
}

func findIntersection(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("один из слайсов пуст")
	}

	elements := make(map[int]bool)
	for _, num := range slice1 {
		elements[num] = true
	}

	var intersection []int
	for _, num := range slice2 {
		if elements[num] {
			intersection = append(intersection, num)
			elements[num] = false
		}
	}

	if len(intersection) == 0 {
		return nil, errors.New("нет общих элементов")
	}

	return intersection, nil
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
	// testCases := [][]int{
	// 	{1, 2, 3, 4, 5},  // нормальный случай
	// 	{2, 4, 6, 8},     // нет нечётных чисел
	// 	{},               // пустой слайс
	// 	{11, 22, 33, 44}, // другой нормальный случай
	// }

	// for _, tc := range testCases {
	// 	sum, indices, err := sumNumbers(tc)
	// 	if err != nil {
	// 		fmt.Printf("Ошибка для слайса %v: %v\n", tc, err)
	// 		continue
	// 	}

	// 	fmt.Printf("Слайс: %v\n", tc)
	// 	fmt.Printf("Сумма нечётных чисел: %d\n", sum)
	// 	fmt.Printf("Индексы нечётных чисел: %v\n", indices)
	// 	fmt.Println("-----")
	// }

	// 12. Напишите программу, которая проверяет, является ли слайс чисел палиндромом, то есть
	// читается ли слайс одинаково в обоих направлениях. В случае палиндрома программа должна
	// вывести true, иначе false.
	// testCases := [][]int{
	// 	{1, 2, 3, 2, 1},    // true (палиндром)
	// 	{1, 2, 3, 4, 5},    // false (не палиндром)
	// 	{},                 // true (пустой слайс)
	// 	{5},                // true (один элемент)
	// 	{1, 2, 2, 1},       // true (четное количество элементов)
	// 	{1, 2, 3, 3, 2, 1}, // true (палиндром)
	// 	{1, 2, 3, 4, 2, 1}, // false (не палиндром)
	// }

	// for _, tc := range testCases {
	// 	result := isPalindromeSlice(tc)
	// 	fmt.Printf("%v -> %t\n", tc, result)
	// }

	// 13. Напишите программу, которая находит сумму всех чисел в слайсе, которые больше среднего значения

	// testCases := [][]int{
	// 	{1, 2, 3, 4, 5},
	// 	{10, 20, 30},
	// 	{},
	// 	{5, 5, 5, 5},
	// 	{2, 4, 6, 8, 10},
	// }

	// for _, tc := range testCases {
	// 	sum, err := sumAboveAverage(tc)
	// 	if err != nil {
	// 		fmt.Printf("Ошибка для слайса %v: %v\n", tc, err)
	// 		continue
	// 	}
	// 	fmt.Printf("Слайс: %v\n", tc)
	// 	fmt.Printf("Сумма чисел больше среднего: %d\n", sum)
	// 	fmt.Println("-----")
	// }
	// 14. Напишите программу, которая генерирует два случайных слайса чисел от 1 до 100 и находит пересечение этих слайсов (элементы, которые встречаются в обоих слайсах).

	slice1, err := generateRandomSlice(10)
	if err != nil {
		fmt.Println("Ошибка генерации первого слайса:", err)
		return
	}

	slice2, err := generateRandomSlice(8)
	if err != nil {
		fmt.Println("Ошибка генерации второго слайса:", err)
		return
	}

	fmt.Println("Первый слайс:", slice1)
	fmt.Println("Второй слайс:", slice2)

	intersection, err := findIntersection(slice1, slice2)
	if err != nil {
		fmt.Println("Ошибка поиска пересечения:", err)
		return
	}

	fmt.Println("Пересечение слайсов:", intersection)

	// 15. Напишите программу, которая генерирует слайс из N случайных чисел от 1 до 100, затем
	// создает два слайса: один с чисел, делящихся на 5, а другой — на 7.

}
