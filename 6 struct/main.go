package main

import (
	"fmt"
)

// 1. Создайте структуру Person с полями Name и Age. Добавьте метод SayHello, который выводит: Привет, меня зовут Name, мне Age лет.

type Person struct {
	name string
	age  int
}

func SayHello(p Person) {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", p.name, p.age)
}

// 2. Определите структуру Product с названием и ценой. Напишите метод ChangePrice(newPrice float64), который обновляет цену.

type Product struct {
	name  string
	price float64
}

func ChangePrice(p Product, newPrice float64) {
	p.price = newPrice

	fmt.Println(p.name, p.price)
}

type Student struct {
	name  string
	grade int
}

func StudentActivity(s Student) {
	if s.grade >= 60 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

type Account struct {
	name    string
	balance float64
}

func Deposit(a Account, newSum float64) {
	a.balance += newSum

	fmt.Println(a.balance)
}

func Withdraw(a Account, newSum float64) {
	a.balance -= newSum

	fmt.Println(a.balance)
}

type Movie struct {
	name   string
	rating float64
}

func IsHit(m Movie) {
	if m.rating >= 8.0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

type User struct {
	name     string
	password string
}

func ChangePassword(user User, oldPass, newPass string) {
	if user.password == oldPass {
		user.password = newPass
		fmt.Println(user, "Пароль успешно изменён")
	} else {
		fmt.Println("Неверный старый пароль")
	}
}

type Item struct {
	Name  string
	Price float64
}

type Store struct {
	Name     string
	Products []Item
}

func (s Store) TotalPrice() float64 {
	total := 0.0
	for _, product := range s.Products {
		total += product.Price
	}
	return total
}

type Client struct {
	Name    string
	Balance float64
}

func FindRichestClient(clients []Client) Client {
	if len(clients) == 0 {
		return Client{}
	}

	richest := clients[0]
	for _, client := range clients {
		if client.Balance > richest.Balance {
			richest = client
		}
	}
	return richest
}

type UserTwo struct {
	Log string
	pwd string
}

func Login(users []UserTwo, login, password string) bool {
	for _, user := range users {
		if user.Log == login && user.pwd == password {
			return true
		}
	}
	return false
}

func main() {

	person := Person{
		name: "Agata",
		age:  11,
	}
	SayHello(person)

	product := Product{
		name:  "TV",
		price: 50000.0,
	}
	ChangePrice(product, 45000.0)

	// 3. Создайте структуру Student, содержащую имя студента и его оценку (число от 0 до 100). Напишите функцию, которая принимает объект Student и возвращает true, если студент сдал экзамен (оценка 60 и выше).
	student := Student{
		name:  "Agata",
		grade: 77,
	}
	StudentActivity(student)
	// 4. Создайте структуру Account, которая содержит имя владельца и текущий баланс. Реализуйте методы Deposit(sum float64) — для пополнения счёта и Withdraw(sum float64) — для снятия со счёта.

	account := Account{
		name:    "Andrey",
		balance: 50000.0,
	}
	Deposit(account, 45000.0)
	Withdraw(account, 15000.0)

	// 5. Создайте структуру Movie с полями Название и Рейтинг. Добавьте метод IsHit(), который возвращает true, если рейтинг фильма 8.0 или выше.

	movie := Movie{
		name:   "Movie",
		rating: 5.0,
	}

	IsHit(movie)

	// 6. Создайте структуру User с полями Имя и Пароль. Добавьте метод ChangePassword(oldPass,	newPass string), который меняет пароль, только если введён правильный старый пароль.
	user := User{
		name:     "Movie",
		password: "oldPass",
	}

	ChangePassword(user, "oldPass", "newPass")

	// 	7. Создайте структуру Item с полями Название и Цена. Создайте структуру Store, которая содержит имя магазина и слайс из товаров []Item. Добавьте метод TotalPrice(), который	возвращает суммарную стоимость всех товаров в магазине.

	items := []Item{
		{Name: "Книга", Price: 500},
		{Name: "Ручка", Price: 50},
		{Name: "Блокнот", Price: 200},
	}

	store := Store{
		Name:     "Канцтовары",
		Products: items,
	}

	fmt.Printf("Общая стоимость товаров в магазине '%s': %.2f руб.\n",
		store.Name, store.TotalPrice())

	// 	8. Создайте структуру Client с полями Имя и Баланс. Создайте слайс клиентов и функцию	FindRichestClient(clients []Client), которая находит и возвращает клиента с самым большим	балансом.
	clients := []Client{
		{Name: "Иван", Balance: 5000},
		{Name: "Петр", Balance: 12000},
		{Name: "Мария", Balance: 8000},
	}

	richest := FindRichestClient(clients)
	fmt.Printf("Самый богатый клиент: %s (Баланс: %.2f)\n", richest.Name, richest.Balance)

	// 	9. Создайте структуру User с полями Логин и Пароль. Создайте функцию Login(users []User, login,	password string), которая проверяет, есть ли пользователь с таким логином и паролем. Если	пользователь найден — верните true, иначе — false
	users := []UserTwo{
		{Log: "admin", pwd: "12345"},
		{Log: "user1", pwd: "qwerty"},
	}

	fmt.Println(Login(users, "admin", "12345"))
	fmt.Println(Login(users, "wrong", "12345"))
}
