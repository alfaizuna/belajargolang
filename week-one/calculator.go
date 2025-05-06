package main

import (
	"fmt"
)

func main(){
	var num1, num2 int
	var operator string

	fmt.Println("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Println("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Println("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Operator tidak valid")
		return
	}

	fmt.Println("Hasil: ", result)
}