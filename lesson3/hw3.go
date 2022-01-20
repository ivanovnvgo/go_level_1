//Третье домашнее задание по курсу "Go. Уровень 1"
//доработать калькулятор: больше операций и валидации данных.
package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string
	fmt.Println("Введите первое число, оператор (*, /, +, -, %(Процент), второе число. После ввода каждого значения нажимайте клавишу ENTER: ")
	fmt.Scan(&a, &op, &b)
	switch op {
	case "+":
		fmt.Printf("%0.2f + %0.2f = %0.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("%0.2f - %0.2f = %0.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("%0.2f * %0.2f = %0.2f\n", a, b, a*b)
	case "%":
		fmt.Printf("%%%0.2f\n", a*b/100)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка! Второе число = 0, на 0 делить нельзя")
		} else {
			fmt.Printf("%0.2f / %0.2f = %0.2f\n", a, b, a/b)
		}
	default:
		fmt.Printf("Ошибка! Неизвестный оператор: %s\n", op)
	}
}
