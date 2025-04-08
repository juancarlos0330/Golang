package main

import (
	"fmt"
)

func declaration() {
	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       // type is inferred

	var a string
	var b int
	var c bool

	a = "James"

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func multi_declaration() {
	var (
		a int
		b int    = 1
		c string = "hello"
		d string
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	multi_declaration()
}
