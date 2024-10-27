/*
	Константы и арифметика. Определи константу для Пи и вычисли площадь и периметр круга с заданным радиусом. Используй функции для этих вычислений.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.14
	var rad float64 = 3

	per := 2 * pi * rad
	fmt.Println("Perimetr=", per)

	sq := pi * math.Pow(rad, 2)
	fmt.Println("Square=", sq)
}
