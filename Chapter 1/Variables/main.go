/*
	Простая программа и переменные.
	Напиши программу, которая выводит на экран информацию о пользователе (имя, возраст), используя переменные разных типов (string, int, bool).
	Используй несколько переменных в одной строке объявления.
*/

package main

import "fmt"

func main() {

	//var name string = "Slava"
	//var age int = 25
	//var isOpenToWork = false

	//name := "Slava";
	//age := 25;
	//isOpenToWork := false;

	var (
		name         string = "Slava"
		age          int    = 25
		isOpenToWork bool   = false
	)

	//fmt.Println(name, age, isOpenToWork)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is open to work:", isOpenToWork)
}
