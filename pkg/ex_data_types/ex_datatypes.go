package exdatatypes

import (
	"fmt"
)

func DeclareArray() {
	fmt.Println("==============array============")
	fmt.Println("first type : array")
	//first type : array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
	//second type : array
	fmt.Println("second type : array")
	var arr2 [3]int
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	fmt.Println(arr2)
	//third type : array
	fmt.Println("third type : array")
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
}

func DeclareSlice() {
	fmt.Println("==============slice============")

	//first type : slice
	fmt.Println("first type : slice")
	var a []int = []int{2, 4, 6}
	fmt.Println(a)
	//second type : slice
	fmt.Println("second type : slice")
	var b = []int{2, 4, 6}
	fmt.Println(b)
	//third type : slice
	fmt.Println("third type : slice")
	var c = make([]int, 3, 4)
	fmt.Println(c)
	fmt.Println("slice length : ", len(c))
	fmt.Println("slice capacity :", cap(c))
	//fourth type : slice
	fmt.Println("append new value to slice")
	c = append(c, 4)
	fmt.Println(c)
}

func DeclareMap() {
	fmt.Println("==============map============")
	//first type : map
	fmt.Println("first type : map")
	var newMap map[string]int = map[string]int{}
	newMap["test"] = 1
	newMap["test2"] = 2
	fmt.Println(newMap)
	fmt.Println(newMap["test"])
	//first type : map
	fmt.Println("second type : map")
	// var newMap2 map[string]int
	// newMap2["test"] = 1
	// newMap2["test2"] = 2
	// fmt.Println(newMap)
	// fmt.Println(newMap["test"])
	fmt.Println("in nil map we can't add new value to it, if we wanna do that we get panic!:{")
	//first type : map
	fmt.Println("third type : map")
	cMap := make(map[string]int, 3)
	cMap["test"] = 2
	fmt.Println(cMap)
	fmt.Println("length of third map :", len(cMap))
	//fourth type : map
	//comma, ok statement in map
	fmt.Println("fourth type : map")
	value, ok := cMap["test2"]
	if !ok {
		fmt.Println("i can't find this key in map", " => test2")
		// return
	}
	fmt.Println("i found it your key requested : ", value)
	//fifth type : map
	//delete map key and value
	fmt.Println("fifth type : map")
	fmt.Println(cMap)
	delete(cMap, "test")
	fmt.Println(cMap)
}
