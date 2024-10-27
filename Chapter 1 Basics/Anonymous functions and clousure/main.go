/*
	Анонимные функции и замыкания.
	Реализуй анонимную функцию, которая принимает два числа и возвращает их сумму.
	Используй замыкание для сохранения результата предыдущего вызова функции.
*/

package main

import "fmt"

func sum() func(int, int) int {
	previousSum := 0

	return func(a, b int) int {
		currentSum := a + b + previousSum
		previousSum = currentSum
		return currentSum
	}
}

func main() {
	f := sum()

	fmt.Println(f(5, 6))
	fmt.Println(f(3, 4))
	fmt.Println(f(1, 2))
}
