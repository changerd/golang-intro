/*
	Функции и возвращение значений. Напиши функцию, которая принимает три числа и возвращает их сумму, среднее значение и максимальное число.
*/

package main

import (
	"fmt"
)

func main() {
	var sum, max, avg = myMethod(54, 79, 35)
	fmt.Printf("Sum=%d\nMax=%d\nAvg=%f", sum, max, avg)
}

func myMethod(x, y, z int) (sum, max int, avg float64) {
	sum = x + y + z
	avg = float64(sum) / 3

	max = x
	if y > max {
		max = y
	}

	if z > max {
		max = z
	}

	return sum, max, avg
}
