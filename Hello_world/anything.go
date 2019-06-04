package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the coolest line of words")

	foo()

	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

		// assign and declear
		x := 42
		fmt.Println(x)

		// change the value
		x = 99
		fmt.Println(x)

		y := 100 + 27
		fmt.Println(y)

	}
}

func foo() {
	fmt.Println("I am in foo now")
}

//control flow
//1. seqyebce
//2. loop; iterative
//3. conditional
