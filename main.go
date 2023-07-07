package main

import (
	"fmt"

	excontrolflow "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_control_flow"
	exdatatypes "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_data_types"
	exfunction "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_function"
	expointer "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_pointer"
	exstruct "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_struct"
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

	//call DeclareArray to define and show array and value of that.
	exdatatypes.DeclareArray()

	//call DeclareSlice to define and show slice and value of that.
	exdatatypes.DeclareSlice()

	//call DeclareMap to define and show map and value of that.
	exdatatypes.DeclareMap()

	//divide output
	fmt.Println("====================")

	//call declare pointer to define pointer variable
	expointer.DeclarePointer()

	//divide output
	fmt.Println("====================")

	//call exstruct to define and show values of struct
	exstruct.ShowStruct()

}
