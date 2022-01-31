package fibonacci

import (
	"github.com/stretchr/testify/require"
	fibonacci "go_level_1/go_level_1/lesson10/fibonacci"
	"testing"
)

// Use testify
func TestFibonacci(t *testing.T) {
	var x, _ = fibonacci.Fibonacci(5)
	var y int = 5
	require.Equal(t, x, y, "The two numbers should be the same.")
}
