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

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
}

func data_types() {
	// var x uint = 500
	// var y uint = 4500
	// fmt.Printf("Type: %T, value: %v", x, x)
	// fmt.Printf("Type: %T, value: %v", y, y)

	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

func goArray() {
	var arr1 = [5]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	arr2[4] = 50
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr3 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr4 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr3))
	fmt.Println(len(arr4))
}

func main() {
	goArray()
}
