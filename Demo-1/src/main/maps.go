package main

import "fmt"
func main() {
	//To create an empty map,
	// use the builtin 
	//make: make(map[key-type]val-type).	
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map : ",m)

	v1 := m["k1"]

	fmt.Println("v1 : ",v1)
	fmt.Println("len : ",len(m))
}