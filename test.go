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

const PI = 3.14

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func constant() {
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

func formatting_verbs() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}

func main() {
	formatting_verbs()
}
