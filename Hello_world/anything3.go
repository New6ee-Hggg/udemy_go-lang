package main

import (
	"fmt"
)

var y = 42

//DECLARE the VAR with the IDENTIFIER "z"
//is of TYPE string
//and I ASSIGN the VALUE "shakend, note stireed"

var z string = "Shaken, not stirred"

var a string = `James said 
"Shakend,

not stirred`

//this is a STATIC programming language
//a VAR is DECLARED to hold a VALUE of cetain TYPE
//not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
