package main

import (
	// "bufio"
	// "fmt"
	// "strings"
)

func main() {

	// 1. Напишите программу, которая принимает два числа и выводит наибольшее из них.
	// var a, b int
	// fmt.Print("Введите первое число (а): ")
	// fmt.Scanln(&a)

	// fmt.Print("Введите первое число (b): ")
	// fmt.Scanln(&b)

	// max := a
	// if b > a {
	// 	max = b
	// }
	// fmt.Printf("Максимальное число: %d\n", max)

	// 2. Напишите программу, которая выводит сумму чисел от 1 до N, где N вводится пользователем.

	// var n int
	// fmt.Print("введите число n: ")
	// fmt.Scanln(&n)

	// for i := 0; i <= n; i++ {
	// 	fmt.Println(i)
	// }

	// 3. Напишите программу, которая по номеру дня недели (от 1 до 7) выводит название дня
	// (например, для 1 — "Понедельник").
	// Ввод: 2 → Вывод: Вторник
	// Ввод: 5 → Вывод: Пятница

	// 	var dayn int
	// 	fmt.Print("Введите номер дня недели (1-7): ")
	// 	fmt.Scanln(&dayn)
	// 	dayName := getDayName(dayn)
	// 	fmt.Println(dayName)
	// }

	// func getDayName(day int) string {
	// 	switch day {
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
	// 	case 7:
	// 		return "Воскресенье"
	// 	default:
	// 		return "Неверный номер дня недели. Введите число от 1 до 7."
	// 	}

	// 4. Напишите программу, которая принимает число N и выводит все числа от 1 до N, делящиеся
	// на 3
	// Ввод: 10 → Вывод: 3 6 9
	// Ввод: 20 → Вывод: 3 6 9 12 15 18

	// var n int
	// fmt.Print("введите число n: ")
	// fmt.Scanln(&n)

	// for i := 0; i <= n; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// 5. Напишите программу, которая проверяет, делится ли введенное число на 5 и 10.
	// Ввод: 25 → Вывод: Делится на 5, но не на 10
	// Ввод: 30 → Вывод: Делится на 5 и на 10

	// var n int
	// fmt.Print("введите число n: ")
	// fmt.Scanln(&n)

	// divisibleBy5 := n%5 == 0
	// divisibleBy10 := n%10 == 0

	// switch {
	// case divisibleBy5 && divisibleBy10:
	// 	fmt.Printf("%d делится на 5 и на 10\n", n)
	// case divisibleBy5:
	// 	fmt.Printf("%d делится на 5, но не на 10\n", n)
	// default:
	// 	fmt.Printf("%d не делится ни на 5, ни на 10\n", n)
	// }

	// 6. Напишите программу, которая принимает 3 числа и выводит максимальное из них.
	// Ввод: 3, 5, 2 → Вывод: 5
	// Ввод: 9, 7, 8 → Вывод: 9

	// var x, y, z int
	// fmt.Print("Введите три числа через пробел: ")
	// fmt.Scanln(&x, &y, &z)

	// max := x
	// if y > max {
	// 	max = y
	// }
	// if z > max {
	// 	max = z
	// }

	// fmt.Printf("максимальное число: %d\n", max)

	// 7. Напишите программу, которая вычисляет факториал числа.
	// Ввод: 5 → Вывод: 120
	// Ввод: 7 → Вывод: 5040

	// 	var n int
	// 	fmt.Print("Введите число для вычисления факториала: ")
	// 	fmt.Scan(&n)

	// 	if n < 0 {
	// 		fmt.Println("Факториал отрицательного числа не определен")
	// 		return
	// 	}

	// 	result := factorial(n)
	// 	fmt.Printf("Факториал %d = %d\n", n, result)
	// }
	// func factorial(n int) int {
	// 	if n == 0 {
	// 		return 1
	// 	}
	// 	return n * factorial(n-1)

	// 8. Напишите программу, которая выводит все нечетные числа от 1 до N.
	// Ввод: 10 → Вывод: 1 3 5 7 9
	// Ввод: 6 → Вывод: 1 3 5

	// var n int
	// fmt.Print("введите число n: ")
	// fmt.Scanln(&n)

	// for i := 0; i <= n; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// 9. Напишите программу, которая принимает строку и подсчитывает количество гласных в ней.
	// Ввод: Hello World → Вывод: 3
	// Ввод: Go Programming → Вывод: 4

	// fmt.Print("Введите строку: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()

	// count := 0
	// for _, c := range text {
	// 		switch c {
	// 		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
	// 				count++
	// 		}
	// }

	// fmt.Println("Гласных букв:", count)

	// 10. Напишите программу, которая проверяет, является ли строка палиндромом.
	// Ввод: madam → Вывод: Да, это палиндром
	// Ввод: hello → Вывод: Нет, это не палиндром

	// 	var input string
	// 	fmt.Print("Введите строку: ")
	// 	fmt.Scanln(&input)

	// 	if isPalindrome(input) {
	// 		fmt.Println("Да, это палиндром")
	// 	} else {
	// 		fmt.Println("Нет, это не палиндром")
	// 	}
	// }

	// func isPalindrome(s string) bool {
	// 	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	// 	for i := 0; i < len(s)/2; i++ {
	// 		if s[i] != s[len(s)-1-i] {
	// 			return false
	// 		}
	// 	}
	// 	return true

	// 11. Напишите программу, которая находит сумму чисел от X до Y.
	// Ввод: 3, 5 → Вывод: 12
	// Ввод: 1, 4 → Вывод: 10

	// var x, y, sum int
	// fmt.Print("Введите два числа через пробел: ")
	// fmt.Scanln(&x, &y)
	// sum = 0
	// for i := x; i <= y; i++ {
	// 	sum += i
	// }
	// fmt.Printf("сумма: %d\n", sum)

	// 12. Напишите программу, которая выводит все простые числа до числа N.
	// Ввод: 10 → Вывод: 2 3 5 7
	// Ввод: 20 → Вывод: 2 3 5 7 11 13 17 19

	// var n int
	// fmt.Print("Введите число N: ")
	// fmt.Scanln(&n)

	// fmt.Printf("Простые числа до %d: ", n)

	// for num := 2; num <= n; num++ {
	// 	isPrime := true
	// 	for i := 2; i*i <= num; i++ {
	// 		if num%i == 0 {
	// 			isPrime = false
	// 			break
	// 		}
	// 	}
	// 	if isPrime {
	// 		fmt.Print(num, " ")
	// 	}
	// }
	// fmt.Println()

	// 13. Напишите программу, которая находит индекс первого вхождения символа в строку.
	// Ввод: hello, e → Вывод: 1
	// Ввод: world, o → Вывод: 1

	// var str string
	// var char rune

	// fmt.Print("Введите строку: ")
	// fmt.Scanln(&str)

	// fmt.Print("Введите символ для поиска: ")
	// fmt.Scanf("%c\n", &char)

	// index := strings.IndexRune(str, char)

	// if index == -1 {
	// 	fmt.Printf("Символ '%c' не найден в строке\n", char)
	// } else {
	// 	fmt.Printf("Индекс первого вхождения '%c': %d\n", char, index)
	// }

	// 14. Напишите программу, которая выводит все числа от 1 до N с шагом 2.
	// Ввод: 10 → Вывод: 1 3 5 7 9
	// Ввод: 6 → Вывод: 1 3 5

	// var N int
	// fmt.Print("Введите N: ")
	// fmt.Scanln(&N)

	// for i := 1; i <= N; i++ {
	// 	fmt.Printf("%d\n", i)
	// }

	// 15. Напишите программу, которая проверяет, является ли число отрицательным.
	// Ввод: -5 → Вывод: Отрицательное число
	// Ввод: 3 → Вывод: Не отрицательное число

	// var N int
	// fmt.Print("Введите число: ")
	// fmt.Scanln(&N)

	// if N >= 0 {
	// 	fmt.Println("Не отрицательное число")
	// } else {
	// 	fmt.Println("Отрицательное число")
	// }

	// 16. Напишите программу, которая принимает 2 строки и выводит их длину.
	// Ввод: Hello, World → Вывод: 5 5
	// Ввод: Go, Programming → Вывод: 2 11

	// var str1, str2 string

	// fmt.Print("Введите две строки через пробел: ")
	// fmt.Scanf("%s %s", &str1, &str2)

	// str1 = strings.TrimSpace(str1)
	// str2 = strings.TrimSpace(str2)

	// fmt.Printf("%d %d\n", len(str1), len(str2))

}
