package main

import (
	// "bufio"
	"fmt"
	// "strings"
)

func main() {

	// 1. Вывести каждый символ строки “Привет” 2 способами:, используя for и for range.

	// str := "Привет"

	// // Способ 1: Обычный for
	// fmt.Println("Способ 1:")
	// runes := []rune(str)
	// for i := 0; i < len(runes); i++ {
	// 	fmt.Printf("%c ", runes[i])
	// }
	// fmt.Println()

	// // Способ 2: for range
	// fmt.Println("Способ 2:")
	// for _, char := range str {
	// 	fmt.Printf("%c ", char)
	// }
	// fmt.Println()

	// 2. Вывести строку GoLang в обратном порядке

	str := "GoLang"

	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println()

	// 3. Проверить, содержит ли строка hello world подстроку world.
	// 4. Подсчитать, сколько раз подстрока go встречается в gogogopher.
	// 5. Заменить все cat на dog в строке cat-cat-dog.
	// 6. Заменить первое вхождение go на GO в строке go go go.
	// 7. Сделать все буквы строки hello заглавными.
	// 8. Убрать пробелы из строки “ hello world “.
	// 9. Разбить строку a,b,c,d на срез строк.
	// 10. Объединить []string{"go", "lang"} в строку с пробелом между ними.
	// 11. Вывести ha 5 раз подряд.
	// 12. Проверить, начинается ли строка golang с go.
	// 13. Проверить, заканчивается ли строка index.html на .html.
	// 14. Из строки h e l l o удалить все пробелы.
	// 15. Для строки GoLang напечатать каждый символ и его код, например: G - 71, o - 111, L - 76 и т.д.
	// 16. Подсчитать количество слов в строке go is awesome.
	// 17. Подсчитать количество заглавных букв в строке GoLang.
	// 18. Если строка заканчивается на ., удалить её.
	// 19. В предложении I love apples, заменить apples на oranges.
	// 20. Учитывать только латинские гласные a, e, i, o, u.
	// 21. Инверсия регистра (если буква — заглавная, сделать строчной и наоборот) GoLang → gOlANG
	// 22. Проверка, является ли строка палиндромом. казак, шалаш → true
	// 23. Поиск самого длинного слова в строке go is an expressive and concise language → expressive
	// 24. Найти все уникальные символы в строке aabccdee → a b c d e
	// 25. Подсчитать количество цифр в строке Пример: в abc123def456 — 6 цифр.
	// 26. Сделать первую букву каждого слова заглавной Пример: go is fun → Go Is Fun.
	// 27. Подсчитать количество символов пунктуации Для строки Hello, world! How are you? должно
	// быть 3 (,, !, ?).
	// 28. Найти подстроки, начинающиеся с заглавной буквы Пример: Welcome To Go Language →
	// Welcome, To, Go, Language
	// 29. Проверить, является ли строка числом Примеры: 123 → true, 12a3 → false
	// Метод: strconv.Atoi
	// 30. Удалить из строки все согласные Пример: banana → aaa
}
