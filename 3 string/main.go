package main

// "strings"

// "bufio"

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

	// str := "GoLang"

	// runes := []rune(str)

	// for i := len(runes) - 1; i >= 0; i-- {
	// 	fmt.Printf("%c", runes[i])
	// }
	// fmt.Println()

	// 3. Проверить, содержит ли строка hello world подстроку world.

	// str := "hello world"
	// fmt.Println(strings.Contains(str, "world"))

	// 4. Подсчитать, сколько раз подстрока go встречается в gogogopher.

	// str := "gogogopher"
	// 1
	// count := strings.Count(str, "go")
	// fmt.Printf("подстрока '%s' встречается %d раз(а)\n", "go", count)

	// 2
	// count := 0
	// n := len("go")
	// for i := 0; i <= len(str)-n; i++ {
	// 	if str[i:i+n] == "go" {
	// 		count++
	// 	}
	// }
	// fmt.Printf("подстрока '%s' встречается %d раз(а)\n", "go", count)

	// 5. Заменить все cat на dog в строке cat-cat-dog.

	// str := "cat-cat-dog"
	// fmt.Println(strings.Replace(str, "cat", "dog", -1))

	// 6. Заменить первое вхождение go на GO в строке go go go.

	// str := "go go go"
	// fmt.Println(strings.Replace(str, "go", "GO", 1))

	// 7. Сделать все буквы строки hello заглавными.

	// str := "hello"
	// fmt.Println(strings.ToUpper(str))

	// 8. Убрать пробелы из строки “ hello world “.

	// str := " hello world "
	// fmt.Println(strings.TrimSpace(str))

	// 9. Разбить строку a,b,c,d на срез строк.

	// str := "a,b,c,d"
	// fmt.Println(strings.Split(str, ","))

	// 10. Объединить []string{"go", "lang"} в строку с пробелом между ними.

	// str := []string{"go", "lang"}
	// fmt.Println(strings.Join(str, ","))

	// 11. Вывести ha 5 раз подряд.

	// str := "ha"
	// fmt.Println(strings.Repeat(str, 5))

	// 12. Проверить, начинается ли строка golang с go.

	// str := "golang"
	// fmt.Println(strings.HasPrefix(str, "go"))

	// 13. Проверить, заканчивается ли строка index.html на .html.

	// str := "index.html"
	// fmt.Println(strings.HasSuffix(str, ".html"))

	// 14. Из строки h e l l o удалить все пробелы.

	// str := "h e l l o"
	// fmt.Println(strings.Replace(str, " ", "", -1))

	// 15. Для строки GoLang напечатать каждый символ и его код, например: G - 71, o - 111, L - 76 и т.д.

	// str := "GoLang"
	// for _, char := range str {
	// 	fmt.Printf("%c - %d, ", char, char)
	// }

	// 16. Подсчитать количество слов в строке go is awesome.
	// // 1
	// str := "go is awesome"

	// count := strings.Count(str, " ") + 1
	// fmt.Printf("количество слов: %d\n", count)
	// // 2
	// str := "go is awesome"
	// words := strings.Split(str, " ")
	// count := len(words)

	// fmt.Printf("количество слов: %d\n", count)
	// // 3
	// str := "go is awesome"
	// words := strings.Fields(str)
	// count := len(words)

	// fmt.Printf("количество слов: %d\n", count)

	// 17. Подсчитать количество заглавных букв в строке GoLang.

	// str := "GoLang"
	// count := 0

	// for _, char := range str {
	// 	if unicode.IsUpper(char) {
	// 		count++
	// 	}
	// }

	// fmt.Printf("Строка: \"%s\"\n", str)
	// fmt.Printf("Количество заглавных букв: %d\n", count)

	// 18. Если строка заканчивается на ., удалить её.

	// str := "Hello World."
	// if strings.HasSuffix(str, ".") {
	// 	fmt.Println(str[:len(str)-1])
	// }

	// 19. В предложении I love apples, заменить apples на oranges.

	// str := "I love apples"
	// fmt.Println(strings.Replace(str, "apples", "oranges", -1))

	// 20. Учитывать только латинские гласные a, e, i, o, u.

	// 21. Инверсия регистра (если буква — заглавная, сделать строчной и наоборот) GoLang → gOlANG

	// str := "GoLang"
	// result := []rune(str)
	// for i, char := range result {
	// 	if unicode.IsUpper(char) {
	// 		result[i] = unicode.ToLower(char)
	// 	} else if unicode.IsLower(char) {
	// 		result[i] = unicode.ToUpper(char)
	// 	}
	// }
	// fmt.Printf("Исходная строка: %s\n", str)
	// fmt.Printf("Результат: %s\n", string(result))

	// 22. Проверка, является ли строка палиндромом. казак, шалаш → true

	// words := []string{"казак", "шалаш"}

	// for _, word := range words {

	// 	clean := strings.ToLower(strings.Replace(word, " ", "", -1))

	// 	isPalindrome := true
	// 	for i := 0; i < len(clean)/2; i++ {
	// 		if clean[i] != clean[len(clean)-1] {
	// 			isPalindrome = false
	// 			break
	// 		}
	// 	}

	// 	fmt.Printf("'%s' → %t\n", word, isPalindrome)
	// }

	// 23. Поиск самого длинного слова в строке go is an expressive and concise language → expressive

	// str := "go is an expressive and concise language"

	// words := strings.Fields(str)

	// longest := ""
	// for _, word := range words {
	// 	if len(word) > len(longest) {
	// 		longest = word
	// 	}
	// }

	// fmt.Printf("Самое длинное слово: \"%s\" (длина: %d)\n", longest, len(longest))

	// 24. Найти все уникальные символы в строке aabccdee → a b c d e

	// str := "aabccdee"
	// unique := make(map[rune]bool)

	// for _, char := range str {
	// 	unique[char] = true
	// }

	// // Выводим результат
	// fmt.Print("Уникальные символы: ")
	// for char := range unique {
	// 	fmt.Printf("%c ", char)
	// }

	// 25. Подсчитать количество цифр в строке Пример: в abc123def456 — 6 цифр.

	// str := "abc123def456"
	// count := 0
	// for i := 0; i < len(str); i++ {
	// 	c := str[i]
	// 	if c >= '0' && c <= '9' {
	// 		count++
	// 	}
	// }
	// fmt.Printf("В строке \"%s\" содержится %d цифр\n", str, count)
	// 26. Сделать первую букву каждого слова заглавной Пример: go is fun → Go Is Fun.

	// str := "go is fun"

	// words := strings.Fields(str)

	// for i, word := range words {
	// 	if len(word) > 0 {

	// 		first := []rune(word)[0]
	// 		if unicode.IsLower(first) {
	// 			words[i] = string(unicode.ToUpper(first)) + word[1:]
	// 		}
	// 	}
	// }

	// result := strings.Join(words, " ")
	// fmt.Println(result)

	// 27. Подсчитать количество символов пунктуации Для строки Hello, world! How are you? должно
	// быть 3 (,, !, ?).

	// str := "Hello, world! How are you?"
	// count := 0

	// for _, char := range str {
	// 	if unicode.IsPunct(char) {
	// 		count++
	// 	}
	// }

	// fmt.Printf("Количество символов пунктуации: %d\n", count)

	// 28. Найти подстроки, начинающиеся с заглавной буквы Пример: Welcome To Go Language →
	// Welcome, To, Go, Language

	// str := "Welcome To Go Language"
	// var result []string

	// words := strings.Fields(str)

	// for _, word := range words {
	// 	if len(word) > 0 && unicode.IsUpper([]rune(word)[0]) {
	// 		result = append(result, word)
	// 	}
	// }

	// fmt.Println(strings.Join(result, ", "))

	// 29. Проверить, является ли строка числом Примеры: 123 → true, 12a3 → false
	// Метод: strconv.Atoi

	// inputs := []string{"123", "12a3"}

	// for _, s := range inputs {
	// 	_, err := strconv.Atoi(s)
	// 	fmt.Printf("%q → %v\n", s, err == nil)
	// }

	// 30. Удалить из строки все согласные Пример: banana → aaa

	// s := "bananaE"
	// for _, c := range s {
	// 	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
	// 		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
	// 		fmt.Print(string(c))
	// 	}
	// }
}
