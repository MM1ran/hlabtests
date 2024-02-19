package main

import "fmt"

func main() {

	//Массив
	arr := [4]int{3, 2, 5, 4}

	//Разность типов
	a1 := [3]int{}
	b1 := [2]int{}
	// (a) [2]int и (b) [3]int - разные типы

	//Ленивое определение размера
	a2 := [...]int{1, 2, 3} // [3]int

	fmt.Println(arr, a1, b1, a2)

	//Клонирование переменных
	var initArray = [...]int{1, 2, 3}
	var copyArray = initArray

	fmt.Printf("Address of initArray: %p\n", &initArray)
	fmt.Printf("Address of copyArray: %p\n", &copyArray)

	/*
		Output:
		  Address of initArray: 0xc00001a018
		  Address of copyArray: 0xc00001a030
	*/

	//Слайсы
	var foo1 []byte
	// С помощью make
	s1 := make([]byte, 5, 5)

	// С помощью shorthand syntax
	bar1 := []byte{}

	fmt.Println(foo1, s1, bar1)

	//Заполнение нулевыми значениями
	var foo2 = make([]byte, 5)
	var bar2 = make([]int, 10)
	var fee = make([]string, 2)

	fmt.Println(foo2, bar2, fee)

	/*
		Output:
		  [0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [ ]
	*/

	//Вместимость не меняет ссылку при append
	var foo3 = make([]int, 5, 10)
	fmt.Printf("Address of foo array [before append]: %p\n", &foo3)

	foo3 = append(foo3, 222)
	fmt.Printf("Address of foo array [after append]: %p\n", &foo3)

	/*
		Output:
			Address of foo array [before append]: 0xc0000aa018
			Address of foo array [after append]: 0xc0000aa018
	*/

	//Увеличение cap вдвое
	var foo4 = make([]int, 10, 10) // Изначальная вместимость - 10
	foo4 = append(foo4, 2)         // Добавляем элемент
	fmt.Println("Length of the array:", len(foo4))
	fmt.Println("Capacity of the array:", cap(foo4))
	/*
		Output:
			Length of the array: 11
			Capacity of the array: 20
	*/

	//shorthand slice

	foo5 := []int{1, 2, 3}
	fmt.Println("Length of the array:", len(foo5))
	fmt.Println("Capacity of the array:", cap(foo5))

	/*
		Output:
			Length of the array: 3
			Capacity of the array: 3
	*/

	//Обрезка слайса
	name := []string{"D", "a", "n", "i", "i", "l"}
	firstThreeLetters := name[:3]
	fmt.Println(firstThreeLetters)

	/*
		Output:
			[D a n]
	*/

	//Динамически расширяемый массив
	nameArray := [6]string{"D", "a", "n", "i", "i", "l"}
	nameSlice := nameArray[:]
	nameSlice = append(nameSlice, "!")
	fmt.Println(nameSlice)

	/*
		Output:
			[D a n i i l !]
	*/

	//Слайсы под капотом

	//Срез на слайсе
	nameArray2 := [6]string{"D", "a", "n", "i", "i", "l"}
	nameSlice2 := nameArray2[:3]
	nameSlice[len(nameSlice2)-1] = "m"
	fmt.Println(nameSlice2) // [D a m]
	fmt.Println(nameArray2) // [D a m i i l]

	//Перераспределение cap в слайсе

	nameArray3 := [6]string{"D", "a", "n", "i", "i", "l"}
	nameSlice3 := nameArray3[:3]
	nameSlice[len(nameSlice3)-1] = "m"
	fmt.Println(nameSlice3) // [D a m]

	// Делаем новый срез
	nameSlice3 = nameSlice3[0:cap(nameSlice3)]
	fmt.Println(nameSlice3) // [D a m i i l]

	//Два слайса ссылаются на один массив
	nameSlice5 := []string{"D", "a", "n", "i", "i", "l"}
	secondNameSlice := nameSlice5
	secondNameSlice[0] = "T"
	fmt.Println(nameSlice5, secondNameSlice) // [T a n i i l] [T a n i i l]

	//Копирование
	nameSlice6 := []string{"D", "a", "n", "i", "i", "l"}
	secondNameSlice2 := make([]string, len(nameSlice6), cap(nameSlice2))
	copy(secondNameSlice2, nameSlice6)
	secondNameSlice2[0] = "T"

	fmt.Println(nameSlice6, secondNameSlice2) // [D a n i i l] [T a n i i l]

}

//Функции Go под капотом
// func append(slice []T, data ...T) []T {
// 	initialLength := len(slice)
// 	finalLength := m + len(data)
// 	if finalLength > cap(slice) { // if necessary, reallocate
// 		// allocate double what's needed, for future growth.
// 		newSlice := make([]byte, (n+1)*2)
// 		copy(newSlice, slice)
// 		slice = newSlice
// 	}
// 	slice = slice[0:finalLength]
// 	copy(slice[initialLength:finalLength], data)
// 	return slice
// }

// func copy(to [1,2,3,4,5]T, from [3,2,1]T) {
// 	for i := range from {

// 		to[i] = from[i]
// 	}
// }
