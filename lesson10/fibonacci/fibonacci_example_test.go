package fibonacci_test

import (
	"fmt"
	fibonacci "go_level_1/go_level_1/lesson10/fibonacci"
)

func ExampleFibonacci() {
	f, err := fibonacci.Fibonacci(5)
	fmt.Printf("%#v %#v ", f, err)
	// Output: 5 <nil>
}
