package funcs

import "fmt"

func Defers() {

	fmt.Println("Hello1")
	fmt.Println("Hello2")
	fmt.Println("Hello3")
	fmt.Println()

	//Запустятся в конце функции Defers и в обратном порядке
	defer fmt.Println("Bye1")
	var bye = "bye2"
	//Значение переменной bye в defer сохранится то, каким оно было при объявлении defer функции
	//Поменяв снизу "bye2" на "bye3" в переменной, значение во втором defer останется тем же
	defer fmt.Println(bye)
	bye = "bye3"
	defer fmt.Println(bye)
}
