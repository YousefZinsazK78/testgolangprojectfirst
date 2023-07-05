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

	fmt.Println("===========use method of employeer and person struct =========")
	fmt.Println(person.GetSalary())
	fmt.Println(person.GetName())
	fmt.Println(person.GetAge())
	fmt.Println("===========use method of employeer and person struct =========")
}

func (p Person) GetSalary() string {
	return fmt.Sprintf("%s has 2000$ salary!", p.Name)
}

func (p Person) GetName() string {
	return fmt.Sprintf("%s is best employee", p.Name)
}

func (p Person) GetAge() string {
	return fmt.Sprintf("%s is %v years old.", p.Name, p.Age)
}
