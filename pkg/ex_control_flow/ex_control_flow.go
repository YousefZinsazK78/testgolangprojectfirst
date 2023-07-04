package excontrolflow

import "fmt"

func ControlFlow() {
	//first type => if else statement
	n := 1
	if n == 2 {
		fmt.Println("n is 2")
	} else {
		fmt.Println("oh no! n is not 2")
	}

	//second type => switch statement
	//in switch statement we don't have break this automatically handle by golang
	//another cool feature in golang is fallthrough keyword for using one case computing for multiple cases....
	switch n {
	case 1:
		fmt.Println("hooray!:) n is 1 ...")
	case 2:
		fmt.Println("oh wow! n is 2 ...")
	default:
		fmt.Println("oh my god! i don't know what to say!!")
	}
}
