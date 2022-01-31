//Benchmark test, запуск: go test -bench=. -benchmem -benchtime=100x bench_test.go
package fibonacci

import (
	fibonacci "go_level_1/go_level_1/lesson10/fibonacci"
	"testing"
)

//fibi - тестовая функция, использующая метод итераций, нужна для сравнения бенчмарков
func fibi(n int) int {
	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

//BenchmarkFibi - тест функции Fibonacci (метод итераций)
func BenchmarkFibi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibi(22)
	}
	//b.Fatalf("End Fibi: %d\n", x)
}

//BenchmarkFibonacci - тест функции Fibonacci (рекурсивная функция)
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.Fibonacci(22)
	}
	//b.Fatalf("End Fibonacci: %d\n", y)
}
