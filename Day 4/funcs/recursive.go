package funcs

import (
	"fmt"
	"log"
)

//_____________________________________________Рекурсивные функции_____________________________________________

// Функция вычисления факториала числа
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Factorial() {
	fmt.Println()
	log.Println("Рекурсия")

	//Функция вызывает саму себя. Подобно бесконечному циклу.
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720
}
