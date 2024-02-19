package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println('😊')

	// fmt.Println(`C:\Users\mm1ran\go`)
	// var text string = "hello world!"

	// var a byte = 'a'
	// fmt.Printf("%c\n", a)

	// var pi rune = 960
	// var alpha rune = 940
	// var omega rune = 969
	// var bang byte = 33

	// var stroka = "text"
	// fmt.Println(stroka[1])
	// s := stroka[1] + 1
	// fmt.Printf("%c", s)

	question := "      hello boy   "

	// for i, c := range question {
	// 	fmt.Printf("%v %c\n", i, c)
	// }

	// message := "uv vagreangvbany fcnpr fgngvba"

	// for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII
	// 	c := message[i]
	// 	if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы
	// 		c = c + 13
	// 		if c > 'z' {
	// 			c = c - 26
	// 		}
	// 	}
	// 	fmt.Printf("%c", c)
	// }
	// fmt.Println(len(question), "bytes")                    // Выводит: 15 bytes
	// fmt.Println(utf8.RuneCountInString(question), "runes") // Выводит: 12 runes
	fmt.Println(strings.Contains(question, "а"))
	a := strings.Replace(question, " ", "", -1)
	fmt.Println(a)
	// c, size := utf8.DecodeRuneInString(question)
	// fmt.Printf("First rune: %c %v bytes", c, size) // Выводит: First rune: ¿ 2 bytes
}

// func main() {

// }
