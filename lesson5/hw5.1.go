//Пятое домашнее задание по курсу "Go. Уровень 1"
//1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
//2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
//4. Посмотрите задачи из предыдущих уроков. Как можно улучшить дизайн задач?
// Что бы вы разбили на отдельные функции или даже пакеты? Часть кода, выполняющую сортировку в домашнем задании
// к уроку 4, я вынес в отдельную функцию.
// Остальные задания состоят из небольшого количества строк, там совсем не чего упрощать.
// Добавил функциональности:
// "Написать функцию, которая будет это делать "на лету".
// Если у неё запросили значение и его нет в кеше - оно будет вычислено и положено в кеш.
// Если же запросили значение которое уже есть в кеше - то будет возвращаться сразу из кеша"
package main

import (
	"fmt"
	"time"
)

//func fib рекурсивно вычисляет n-ое число ряда Фибоначчи.
//2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
func fib(n int) int {
	if 0 <= n && n < 2 { //Добавил условие 0 <= n &&
		return n // Изменил 1 на n
	} else {
		return fib(n-2) + fib(n-1)
	}
}

// Функция func fibToMap заполняет карту значениями последовательности чисел Фибоначии
func fibToMap(s, n int, nFibonacci map[int]int) map[int]int {
	for i := s; i < n; i++ {
		nFibonacci[i] = fib(i)
		fmt.Println(i, fib(i)) // Контроль вычислений последовательности чисел, после отладки можно убрать
	}
	return nFibonacci
}
func main() {
	var n int
	fmt.Print("Для вычисления числа ряда Фибоначчи введите целое положительное число и нажмите клавишу ENTER: ")
	fmt.Scanf("%d\n", &n)
	//Создал карту (map) для хранения предварительно расчитанной последовательности чисел Фибоначчи
	nFibonacci := make(map[int]int)
	var s int              // Первый элемент для расчета последовательности числе Фибоначии
	start := time.Now()    // start time
	fmt.Println("n = ", n) //отладка
	fibToMap(s, n, nFibonacci)
	duration := time.Since(start) // full time
	fmt.Println(duration)
	fmt.Println("Вычисленное значение из последовательности ряда чисел Фибоначчи: ", nFibonacci[n-2]+nFibonacci[n-1])
	var x int
	fmt.Print("Для вычисления еще одного числа ряда Фибоначчи введите целое положительное число и нажмите клавишу ENTER: ")
	fmt.Scanf("%d\n", &x)
	if 0 <= x && x < n { // Добавил условие  0 <= x &&
		fmt.Println("Print from if") //отладка
		fmt.Println("Вычисленное значение из последовательности ряда чисел Фибоначчи: ", nFibonacci[x-2]+nFibonacci[x-1])
	} else if x > n {
		fmt.Println("Print from else, n, x", n, x) //отладка
		fibToMap(n, x, nFibonacci)
		fmt.Println("Вычисленное значение из последовательности ряда чисел Фибоначчи: ", nFibonacci[x-2]+nFibonacci[x-1])
	} else {
		fmt.Println("Вы ввели отрицательное число")
	}
}
