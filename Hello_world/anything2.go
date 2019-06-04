package main

import "fmt"

var y=44

//declear there is a variable with an identifier "z"
// and that the variable with the identifier "z" is of TYPE int
//ASSIGNS the ZERO VALUE of TYPE into "z"
var z int

func main(){
	//short delaration operator
	//DECLARE a variable and ASSIGN a VALUE(of a certain TYPE)

	x:=42
	fmt.Println(x)
	//:= works fine inside the {}, but the var can

	foo()
	//var y = 43
	fmt.Println(y)

	fmt.Println(z)
}

func foo(){
	fmt.Println(y)
}