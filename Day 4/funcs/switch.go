package funcs

import "fmt"

func Foo(n int) int {
	fmt.Println(n)
	return n
}

//Наглядный пример проверки switch'а. Число в условии равно двум, потому case Foo(3) котрый был через запятую в условии не будет проверяться
func Switchcs() {
	fmt.Println()
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")

		//fallthrough перебрасывает на следующий case внизу
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
}
