package main

import (
	"fmt"
	"strings"
)

func printHello() {
	fmt.Println("Hello World")
}
func Add(a int, b int) {
	fmt.Println(a + b)
}
func isLogin(a bool) bool {
	return a
}

//Variadic function

func f1(a ...int) {
	fmt.Println(a)
}

func personInfromation(age int, names ...string) string {
	fullName := strings.Join(names, " ")

	return fullName
}

//Defer Statement

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}

func main() {
	fmt.Println("Functions in Golang")
	printHello()
	Add(1, 2)
	islogin := isLogin(true)
	fmt.Println(islogin)
	f1(1, 3, 4, 5, 5, 5)
	f1()

	person := personInfromation(10, "Ayomide", "Joshua", "Lade", "obi")
	fmt.Println(person)

	defer bar()
	foo()
	//foo runs before bar

	//Anonymous Functions
	//it does not contain any name used to form closures
	func(message string) {
		fmt.Println(message)

	}("I am an anonymous function")

	a := inCrease(10)
	//a()              //increment happens here
	fmt.Println(a()) //prints the increment which is 12 becuase of another increment
}

func inCrease(x int) func() int {
	return func() int {
		x++
		return x
	}
}
