package main

import (
	"fmt"

	excontrolflow "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_control_flow"
	exfunction "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_function"
	exvariable "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_variable"
)

func main() {
	fmt.Println("hello, world!")

	//divide output
	fmt.Println("====================")

	//call assign variable from exvariable package
	exvariable.AssignVariable()

	//divide output
	fmt.Println("====================")

	//call control flow from excontrolflow package
	excontrolflow.ControlFlow()

	//divide output
	fmt.Println("====================")

	//call Example function to run simple fmt.Println.
	exfunction.ExampleFunction()

	//call ExampleFunctionWithParameters to sum this two parameters and print it.
	fmt.Println(exfunction.ExampleFunctionWithParameters(2, 3))

	//call ExampleFunctionWithParameters to double this two parameters and print it.
	a, b := exfunction.ExampleFunctionWithMultipleReturn(2, 3)
	fmt.Println(a, b)

	//divide output
	fmt.Println("====================")

}
