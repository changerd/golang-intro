/*
	Реализуй программу, которая выполняет деление двух чисел, введённых пользователем.
	Используй defer для вывода сообщения о завершении операции, а panic для обработки ошибок, если происходит деление на ноль.
*/

package main

import (
	"fmt"
	"os"
)

func divide(a, b float64) float64 {
	if b == 0 {
		panic("Divide by zero error!")
	}

	return a / b
}

func main() {
	operationSuccessful := true

	//Handle panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured:", r)
			operationSuccessful = false
		}
		if operationSuccessful {
			fmt.Println("Operation successful")
		}
	}()

	var a, b float64
	fmt.Print("First number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Enter error:", err)
		os.Exit(1)
	}
	fmt.Print("Second error: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Enter error:", err)
		os.Exit(1)
	}

	result := divide(a, b)
	fmt.Printf("Result: %.2f\n", result)
}
