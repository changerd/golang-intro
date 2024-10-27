/*
	Рекурсия. Напиши рекурсивную функцию для вычисления факториала числа. Обработай случай, когда введённое число отрицательное, с помощью условного оператора.
*/

package main

import "fmt"

func fact(x int) int {
	if x == 1 {
		return 1
	}

	return x * fact(x-1)
}

func main() {
	num := 5
	fmt.Println(fact(num))
}
