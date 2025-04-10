package main

// "fmt"
// "math"
// "strconv"

func main()  {
// 1. Проверьте возраст у входных данных :
// Если < 18 — "Вы несовершеннолетний".
// Если 18-60 — "Вы взрослый".
// Если > 60 — "Вы пожилой человек".

// const age int = 25
// if age<18 {
// 	fmt.Println("Вы несовершеннолетний")
// }else if age >= 18 && age <= 60 {
// 	fmt.Println("Вы взрослый")
// }else{
// 	fmt.Println("Вы пожилой человек")
// }
// 2. Оцените успеваемость входных данных:
// 90-100 — "Отлично".
// 75-89 — "Хорошо".
// 50-74 — "Удовлетворительно".
// < 50 — "Неудовлетворительно".

// const progress int = 95
// if progress >=90 && progress <=100{
// 	fmt.Println("Отлично")
// } else if progress >=75 && progress <=89 {
// 	fmt.Println("Хорошо")
// }else if progress >=50 && progress <=74 {
// 	fmt.Println("Удовлетворительно")
// }else if progress <=50 {
// 	fmt.Println("Неудовлетворительно")
// }

// 3. Выведите день недели по номеру (1-7) с использованием switch

// var dayOfTheWeek int
// fmt.Print("Введите номер дня недели (1-7): ")
// fmt.Scanln(&dayOfTheWeek)
// switch dayOfTheWeek {
// case 1:
// 	fmt.Println("Monday")
// case 2:
// 	fmt.Println("Tuesday")
// case 3:
// 	fmt.Println("Wednesday")
// case 4:
// 	fmt.Println("Thursday")
// case 5:
// 	fmt.Println("Friday")
// case 6:
// 	fmt.Println("Saturday")
// case 7:
// 	fmt.Println("Sunday")
// }

// 4. Проверьте, является ли число четным или нечетным.

// var number int
// fmt.Print("введите число для проверки четности: ")
// fmt.Scanln(&number)
// if number%2==0 {
// 	fmt.Println("четное")
// } else {
// 	fmt.Println("нечетное")
// }

// 5. Проверьте, имеет ли человек право на льготу (меньше 18 или старше 65 лет).
// Вход: 15 → "Льгота есть"
// Вход: 40 → "Льгота нет"
// Вход: 70 → "Льгота есть"

// var age int
// fmt.Print("введите возраст: ")
// fmt.Scanln(&age)
// if age>18 && age<65 {
// 	fmt.Println("Льгота нет")
// } else {
// 	fmt.Println("Льгота есть")
// }

// 6. Объявите переменные разных типов (int, float64, string, bool) и присвойте значения.
// Выведите их.

// var age int 
// age = 25
// fmt.Println(age)

// var number float64 
// number = 65.5
// fmt.Println(number)

// var value string 
// value = "String"
// fmt.Println(value)

// var cred bool 
// cred = true
// fmt.Println(cred)

// // Короткая инициализация
// age := 25
// number := 65.5
// text := "String"
// isValid := true

// fmt.Println("Значения:")
// fmt.Println(age, number, text, isValid)

// fmt.Println("\nАдреса в памяти:")
// fmt.Println(&age, &number, &text, &isValid)


// 7. Преобразуйте целое число в строку и строку в число.
// Вход: 123 → "123" (преобразовать в строку)
// Вход: "456" → 456 (преобразовать в число)

// // Число → строка
// num := 123
// str1 := strconv.Itoa(num) 
// str2 := fmt.Sprint(num) 

// fmt.Printf("%d → %q (через Itoa)\n", num, str1)
// fmt.Printf("%d → %q (через Sprint)\n", num, str2)

// // Cтрока → число
// str := "456"
// strAsNumber, err := strconv.Atoi(str)
// if err != nil {
// 	fmt.Println("Ошибка:", err)
// 	return
// }
// fmt.Printf("%q → %d\n", str, strAsNumber)

// 8. Введите число с плавающей точкой и преобразуйте его в целое число (округлить).
// Ввод: 3.14 → 3
// Ввод: 7.99 → 8

// var floatNum float64

// fmt.Print("Введите число с плавающей точкой: ")
// fmt.Scanln(&floatNum)

// rounded := math.Round(floatNum)
// fmt.Printf("%.2f → %d\n", floatNum, int(rounded))

// 9. Выполните операции с входными числами: Сложение Вычитание Умножение Деление
// Остаток от деления

// var a, b int
    
//    fmt.Print("Введите первое число: ")
//     fmt.Scanln(&a)
    
//     fmt.Print("Введите второе число: ")
//     fmt.Scanln(&b)

   
//     sum := a + b 
//     diff := a - b
// 		product := a * b 
//     quotient := a / b 
//     remainder := a % b

//     fmt.Printf("%d + %d = %d\n", a, b, sum)
//     fmt.Printf("%d - %d = %d\n", a, b, diff)
//     fmt.Printf("%d * %d = %d\n", a, b, product)
//     fmt.Printf("%d / %d = %d\n", a, b, quotient)
//     fmt.Printf("%d %% %d = %d\n", a, b, remainder)

// 10. Вычислите периметр прямоугольника по длине и ширине.
// Вход: длина = 5, ширина = 3 → Периметр = 16
// var length, width int

//    fmt.Print("Введите длину прямоугольника: ")
//     fmt.Scanln(&length)
    
//     fmt.Print("Введите ширину прямоугольника: ")
//     fmt.Scanln(&width)
//     perimeter := 2 * (length + width)

//     fmt.Printf("Периметр прямоугольника (длина=%d, ширина=%d) = %d\n", length, width, perimeter)

// 11. Вычислите среднее арифметическое двух чисел.
// Вход: 10 и 20 → Среднее = 15

// a, b := 10, 20
// average := (a + b) / 2
// fmt.Printf("Среднее целых %d и %d равно %d\n", a, b, average)

// 12. Сравнение двух входных чисел
// Проверьте: Если a > b — "a больше" Если a == b — "a равно b" Если a < b — "b больше"

// var a, b int
    

// fmt.Print("Введите первое число (a): ")
// fmt.Scanln(&a)

// fmt.Print("Введите второе число (b): ")
// fmt.Scanln(&b)

// switch {
// case a > b:
//     fmt.Println("a больше")
// case a == b:
//     fmt.Println("a равно b")
// default:
//     fmt.Println("b больше")
// }

// 13. Найдите максимальное из двух входных чисел.

// var x, y int
//     fmt.Print("Введите два числа через пробел: ")
//     fmt.Scanln(&x, &y)
    
//     max := x
// 		if y > x {
// 			max = y
// 	}
//     fmt.Printf("Максимальное число: %d\n", max)

// 14. Найдите минимальное из трех входных чисел.

// var x, y, z int
//     fmt.Print("Введите три числа через пробел: ")
//     fmt.Scanln(&x, &y, &z)
    
//     min := x
//     if y < min {
//         min = y
//     }
//     if z < min {
//         min = z
//     }
   
//     fmt.Printf("Минимальное число: %d\n", min)

// 15. Проверьте, равны ли два входных числа.
// var x, y int
//     fmt.Print("Введите два числа через пробел: ")
//     fmt.Scanln(&x, &y)
    
//     if x==y {
//         fmt.Println("Числа равны")
//     } else {
//         fmt.Println("Числа не равны")
//     }
}

