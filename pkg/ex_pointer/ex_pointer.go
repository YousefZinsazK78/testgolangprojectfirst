package expointer

import "fmt"

func DeclarePointer() {
	var digit2 int = 2
	//POINTER
	//in pointer variable instead of getting copy of value with get copy of address of variable in memory that allocated
	//for store pointer to use asterisk symbol before our type
	//for get copy of variable address we use ampersand symbol before our target variable
	//get address = &YourVariableName
	// var myValue = 2
	// a *int := &myValue
	var newPointerToInt *int = &digit2
	fmt.Printf("newPointerToInt address of variable pointer : %v\n", newPointerToInt)
	//for getting value from pointer you must put asterisk before your pointer variable
	//that in golang called derefrence pointer *variablename
	fmt.Printf("newPointerToInt value pointer : %v\n", *newPointerToInt)
	fmt.Printf("newPointerToInt own address in memory : %p\n", &newPointerToInt)
	fmt.Println("change value of pointer with other function :")
	//second type : pointer
	fmt.Println("second type : pointer")
	fmt.Println("change value of pointer without pointer parameters function")
	changeValueWithoutPointer(*newPointerToInt)
	fmt.Printf("newPointerToInt value pointer : %v\n", *newPointerToInt)
	//third type : pointer
	fmt.Println("third type : pointer")
	fmt.Println("change value of pointer")
	changeValueWithPointer(newPointerToInt)
	fmt.Printf("newPointerToInt value pointer : %v\n", *newPointerToInt)

}

func changeValueWithPointer(n *int) {
	*n = *n + 2
}

func changeValueWithoutPointer(n int) {
	n = n + 2
}
