package funcs

import (
	"fmt"
	"log"
)

func Pointers() {
	log.Println("Указатели")
	//'*' Ставится перед ТИПОМ. Значит тип переменной pointer - указатель, который ссылается на значение типа int
	var pointer *int

	var x int = 4        //Переменная зарезервировала место в памяти размером типа Int (32-bit)
	pointer = &x         //Знак амперсанта '&' - ссылка/адрес в памяти к значению
	fmt.Println(pointer) //0xc00000abc123 - Адрес, на который ссылается переменная типа указатель "pointer"

	log.Println("Оператор разыменования")
	//Переменная типа указатель на Int
	var p *int = &x
	*p += 3 //'*' перед ПЕРЕМЕННОЙ означает, что это оператор разыменования. Т.е получение значения, которое прячется под адресом к значению (&)

	fmt.Println(*&x) //Обращайте внимание на ошибки и предупреждения, которые выдаёт программа. Они могут помочь в улучшении (или починке) кода. Наведите курсор на жёлтое подчёркивание

	fmt.Println(*p) //Выведем число, которое хранится в "x" после сложения его с числом выше
	fmt.Println(x)  //Выведем саму переменную, чтобы убедиться, что всё сходится

	log.Println("nil и пустой указатель")
	//Указатель по умолчанию не привязан к какому-либо адресу, поэтому его значение по умолчанию равно 'nil', т.е "ничего"
	var nothing *int
	fmt.Println(nothing) // Выводит: <nil>

	//Объявленная переменная с типом функции так же имеет по умолчанию nil
	var fn func(a, b int) int
	fmt.Println(fn == nil) // Выводит: true, так как функция разыменовывает nil (ничего) и нет функции присвоенной fn

	var pf *float64
	//Проверка на nil во избежание критических ошибок
	if pf != nil {
		fmt.Println("Value:", *pf) //Без этой проверки сработает паника и программа экстренно завершит работу из-за ошибки
	}

	log.Println("Функция new()")
	//Создали безымянный объект в памяти типа Int8(8-bit). Доступ к объекту в памяти доступен теперь только через переменную "ap", которая автоматом стала указателем
	ap := new(int8)

	fmt.Println("Value:", *ap) // Value: 0 - значение по умолчанию вместо 'nil'. Особенность функции new().
	*ap = 8                    // изменяем значение

	fmt.Println("Value:", *ap) // Value: 8

	log.Println("Указатель на указатель")
	//Что будет если указатель сделать ссылкой и привязать к нему ещё один указатель?
	newAp := &ap //Наведя курсор на newAp можно увидеть, что он стал "указателем на тип указатель, который ссылается на тип int8"

	//Если вывести второй указатель станет ясно, что он так же занимает отдельное место в памяти. Получилась цепь из ссылок и указателей.
	fmt.Println("Address of the second pointer:", newAp) // Выдаст иную ссылку, отличимую от первого указателя.

	//Разыменование выдаст результат указателя "ap". А учитывая, что он хранит в себе адрес, он его и вернтёт
	fmt.Printf("Address of the second pointer: %v is similar to the first : %v\n", ap, *newAp)

	//Для вывода значения переменной из первого указателя, нужно сделать разыменование ещё раз, тоесть разыменовать уже полученную ссылку после первого разыменования (после первой звёздочки - *)
	fmt.Println("Value from second pointer plus one:", **newAp+1) // Value: 8. Первая '*' покажет то, что хранит в себе "ap", тоесть ссылку на new(int8). Вторая '*' уже разыменует вторую ссылку, которую мы получили из указателя "ap"
}
