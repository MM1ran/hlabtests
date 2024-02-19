package main

import (
	"fmt"
	"log"
)

// Структура для работы с мапой. Вымышленная база данных пользователей
type user struct {
	name    string
	surname string
	age     int
	gender  string
}

func main() {

	log.Println("Хеш-таблицы/Мапы")
	//Мапа с ключом в виде строки, и значением - структурой
	peoples := map[string]user{}
	//Создадим парочку тестовых пользователей
	Jason := user{
		name:    "Jason",
		surname: "Statham",
		age:     56,
		gender:  "Male",
	}

	Batman := user{
		name:    "Kevin",
		surname: "Conroy",
		age:     37,
		gender:  "Male",
	}

	//Запишем пользователей в базу (мапу/хеш-таблицу) под ключом "Statham" и "Batman"
	peoples["Statham"] = Jason
	peoples["Batman"] = Batman

	//Проверим, есть ли в нашей базе пользователь под ключом "Spider-Man"
	_, ok := peoples["Spider-Man"]
	//Если его нет, то создадим и заполним его поля сразу же
	if !ok {
		peoples["Spider-Man"] = user{name: "Tobey", surname: "Maguire", age: 48}
	}

	//Получим структуру из ключа "Spider-Man". Т.е получим пользователя
	human1 := peoples["Spider-Man"]
	fmt.Printf("%+v", human1)
	fmt.Printf("%+v", peoples)

	//Новая мапа для итерации по ней
	var m map[int]string = map[int]string{1: "test1", 2: "test2", 3: "test3", 4: "test4", 5: "test5"}
	//Вместо индекса, при итерации над мапой будет ключ
	for k, v := range m {
		//Результат может идти не по порядку
		fmt.Println(k, v)
	}

	log.Println("Panic и Recover")
	//Запустим функцию с потенциальной паникой
	mnozh(-1, 0)

	//Показатель того, что рекавер позволяет продолжить работу программы даже с паникой
	fmt.Println("Выжившее сообщение после паники!")
}

func mnozh(a, b int) (result int, err error) {
	//Рекавер для спасения от паники
	//Обязательно через defer, потому что defer срабатывает ПОСЛЕ паники
	defer func(a int) {
		if rec := recover(); rec != nil {
			log.Println("Panic handled!", a, rec)
		}
		//Анонимная функция. Потому в конце стоят скобки с передаваемым параметром
	}(a)

	//Вызываем специально панику при верном условии, не давая закончить результат
	if b <= 0 || a <= 0 {
		panic("\n[ERROR!!!]\n\n Number B or A is Below Zero!")
	}

	//Если сработает паника, данная часть кода не сработает
	result = a / b
	return result, err
}
