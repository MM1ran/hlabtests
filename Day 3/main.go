package main

import (
	"day3/funcs"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {

	log.Println("Переменные и типы данных")
	funcs.VariablesAndTypes()
	fmt.Println()

	log.Println("Функции")

	fmt.Println(funcs.Plus(1, 7))
	fmt.Println()

	funcs.Defers()
	fmt.Println()

	funcs.Anonnymous()
	fmt.Println()

	log.Println("Импортирование модуля")
	//Функция из нового пакета выводит приветствие.
	fmt.Println(quote.Hello())
}

//Загрузка стороннего пакета: "go get "название ссылки""
//Загруженный мной для примера пакет:
//rsc.io/quote
//После загрузки создастся файл go.sum, а в go.mod появится информация о пакете и его зависимостях
