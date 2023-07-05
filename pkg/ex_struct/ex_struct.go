package exstruct

import "fmt"

//Person is to store person information.
type Person struct {
	Name string
	Age  int
}

func ShowStruct() {
	//first type : struct
	//declare person struct .
	person := Person{
		Name: "yousef",
		Age:  23,
	}
	fmt.Println("Person value : ", person)
	fmt.Println("Person name : ", person.Name)
	fmt.Println("Person age : ", person.Age)
	//second type : struct .
	//declare person struct .
	per := Person{"sina", 18}
	fmt.Println("Person value : ", per)
	fmt.Println("Person name : ", per.Name)
	fmt.Println("Person age : ", per.Age)
}
