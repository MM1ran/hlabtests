package funcs

import "fmt"

func Anonnymous() {

	//Запись анонимной функции в переменную и последующее использование
	f1 := func(x, y int) int { return x + y }
	fmt.Println(f1(3, 4)) // 7
	fmt.Println(f1(6, 7)) // 13

	//Фактически функция выше объявлена как:
	//var f func(int, int) int = func(x, y int) int{ return x + y}

	//Анонимная функция как аргумент функции
	//Объявляем функцию в аргументе и сразу же вызываем
	action(10, 25, func(x int, y int) int { return x + y }) // 35
	action(5, 6, func(x int, y int) int { return x * y })   // 30

	//Анонимная функция как результат функции
	f2 := selectFn(1)
	fmt.Println(f2(2, 3)) // 5
	fmt.Println(f2(4, 5)) // 9
	fmt.Println(f2(7, 6)) // 13
}

//Анонимная функция как аргумент функции
func action(n1 int, n2 int, operation func(int, int) int) {
	//Код внутри функции operation будет производиться над n1 и n2. Содержимое функции может быть каким угодно, пока их аргументы и возвращаемое значение будут равны
	result := operation(n1, n2)
	fmt.Println(result)
}

//Анонимная функция как результат функции
func selectFn(n int) func(int, int) int {
	//Функция срабатывает на моменте возврата ответа. В ответ приходит уже результат числа. Таким образом мы можем в разных случаях возвращать результат разных операций, как на примере ниже
	if n == 1 {
		//Результат всё так же остаётся int
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}
