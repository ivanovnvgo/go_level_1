//Unit.test, запуск: go test fibonacci_test.go
package fibonacci

import (
	fibonacci "go_level_1/go_level_1/lesson10/fibonacci"
	"testing"
)

// Проверка на ввод отрицательного числа: тест на задымление
func TestFibonacciSmoke(t *testing.T) {
	result, err := fibonacci.Fibonacci(-10)
	if err == nil {
		t.Fatalf("Было введено отрицательное число, а ошибка не получена")
	}
	expected := 0
	if result != expected {
		t.Fatalf("При вводе отрицательного числа, Fibonacci должна вернуть %d, а вернула %d\n", expected, result)
	}
}

// Проверка на равенство с эталонной последовательностью: табличный тест
func TestFibOnMapTable(t *testing.T) {
	mapTestFibonacci := map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 8: 21, 9: 34, 10: 55, 11: 89, 12: 144, 13: 233, 14: 377, 15: 610, 16: 987, 17: 1597, 18: 2584, 19: 4181, 20: 6765, 21: 10946, 22: 17711}
	mapFibonacci := make(map[int]int)
	mapFibonacci = fibonacci.FibOnMap(22, 22, mapTestFibonacci)
	for i := 0; i < 22; i++ {
		//fmt.Println(mapFibonacci[i], mapTestFibonacci[i]) // Test print
		if mapFibonacci[i] != mapTestFibonacci[i] {
			t.Fatalf("Последовательность Фибоначчи рассчитана не верно")
		}
	}
}
