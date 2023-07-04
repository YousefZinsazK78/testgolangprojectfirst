package main

import (
	"fmt"

	excontrolflow "github.com/yousefzinsazk78/testgolangprojectfirst/pkg/ex_control_flow"
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
}
