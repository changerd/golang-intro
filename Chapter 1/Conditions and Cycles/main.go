/*
	Условные выражения и циклы.
	Напиши программу, которая проверяет, является ли число простым. Входное число вводится с клавиатуры.
	Реализуй проверку с помощью циклов и условных операторов.
*/

package main

import "fmt"

func main() {
	var (
		num    int  = 13
		result bool = true
	)

	for i := 2; i < num; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}

	if result {
		fmt.Println("Input number is prime!")
	} else {
		fmt.Println("Input number is not prime")
	}
}
