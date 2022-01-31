package fibonacci

import (
	"flag"
	"fmt"
)

// Fibonacci рекурсивно вычисляет n-ое число ряда Фибоначчи.
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n < 0")
	}
	if 0 < n && n < 2 {
		return n, nil
	} else {
		f1, _ := Fibonacci(n - 2)
		f2, _ := Fibonacci(n - 1)
		return f1 + f2, nil
	}
}

// FibOnMap заполняет карту nFibonacci значениями последовательности чисел Фибоначчи.
func FibOnMap(s, n int, nFibonacci map[int]int) map[int]int {
	for i := s; i <= n; i++ {
		nFibonacci[i], _ = Fibonacci(i)
		fmt.Println(i, nFibonacci[i]) // Контроль вычислений последовательности чисел, после отладки можно убрать
	}
	return nFibonacci
}

// IOFibonacci запрашивает n-ое число последовательности ряда Фибоначчи и выводит его значение
func IOFibonacci() (int, int, error) {
	fmt.Println("Введите первое и второе число для расчета последовательности ряда Фибоначчи")
	var nFlag = flag.Int("nFlag", 0, "n-ое число последовательности ряда Фибоначчи")
	var sFlag = flag.Int("sFlag", 0, "n-ое число последовательности ряда Фибоначчи")
	flag.Parse()
	var n int
	var s int
	n = *nFlag
	s = *sFlag
	var firstFibonacci, secondFibonacci int //Рассчитанные числа Фибоначчи
	//Создал карту (map) для хранения предварительно рассчитанной последовательности чисел Фибоначчи для функции FibOnMap
	mapFibonacci := make(map[int]int)
	if n < 0 {
		return 0, 0, fmt.Errorf("С отрицательными числами не работаем! Вы ввели: %d\n", n)
	}
	//Заполняем карту числами Фибоначчи от 0 до первого введенного числа n
	FibOnMap(0, n, mapFibonacci)
	firstFibonacci = mapFibonacci[n-2] + mapFibonacci[n-1] //Получили первое число Фибоначчи
	if s <= n {
		secondFibonacci = mapFibonacci[s-2] + mapFibonacci[s-1] //Получили второе число Фибоначчи,
		// если его значение уже рассчитано
	} else {
		FibOnMap(n, s, mapFibonacci)
		secondFibonacci = mapFibonacci[s-2] + mapFibonacci[s-1] //Получили второе число Фибоначчи,
		// если его значение больше, чем уже расчитанные и его нужно досчитать
	}
	return firstFibonacci, secondFibonacci, nil
}
