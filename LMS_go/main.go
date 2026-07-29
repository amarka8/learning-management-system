package main

import "fmt"

func main() {
	fmt.Println("Testing if Go works")
	// Variables with Initializers dont need a declared type
	var bool1, bool2, bool3 = true, false, true
	fmt.Printf("The values of the three booleans are %t %t %t\n", bool1, bool2, bool3)
	// declaring variables without initializers
	var bool1_1, bool2_1, bool3_1 bool
	fmt.Printf("The values of the three booleans are %t %t %t\n", bool1_1, bool2_1, bool3_1)
	// variables declared using assignment statement with implict type without var
	bool_1_2 := true
	fmt.Println("The final value is", bool_1_2)
}
