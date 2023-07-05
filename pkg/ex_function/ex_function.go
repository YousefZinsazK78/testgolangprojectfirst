package exfunction

import "fmt"

// ExampleFunction is to print sample text to standard output.
func ExampleFunction() {
	fmt.Println("ExampleFunction called successfully...")
}

// ExampleFunctionWithParameters is to double two integer parameters and return it.
func ExampleFunctionWithParameters(m, n int) int {
	return m + n
}

// ExampleFunctionWithMultipleReturn is to return two value.
func ExampleFunctionWithMultipleReturn(m, n int) (int, int) {
	return m, n
}
