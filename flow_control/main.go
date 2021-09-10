package main

// run with go run main.go arg1 arg2 arg3 arg4
//Note: arg4 should be an interger
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, Flow controls in GO")

	//if,else if, esle Statements

	price, inStock := 100, true
	_ = inStock

	if price <= 100 && inStock {
		fmt.Println("Too Expensive!, Price is:", price, "in stock:", inStock)
	}

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Its Expensive")
	}

	age := 100

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years !\n", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please vote, it's important!!")
	} else {
		fmt.Println("Invalid age!")
	}

	//Command line arguments

	fmt.Println("os.Args:", os.Args) //run with go run main.go arguments
	fmt.Println("Path:", os.Args[0]) // go run main.go Ayomide flutter Physics
	//Path: C:\..\AppData\Local\Temp\go-build3884311586\b001\exe\main.exe
	fmt.Println("1st arguent:", os.Args[1])       //Ayomide
	fmt.Println("1st arguent:", os.Args[2])       // Physics
	fmt.Println("1st arguent:", os.Args[3])       //Flutter
	fmt.Println("No of arguments:", len(os.Args)) //4
	// note: arguments are always strings type(convert to wanted type)
	var result, err = strconv.ParseFloat(os.Args[4], 64)
	fmt.Println("Error:", err) // prints out any error generated
	fmt.Println("Result:", result)

	//Simple If Statement

	i, err := strconv.Atoi("50")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", i)
	}

	if k, err := strconv.Atoi("20"); err == nil {
		fmt.Println("Value:", k)
	} else {
		fmt.Println("Error:", err)
	}

	//For loops

	for i1 := 0; i1 <= 10; i1++ {
		fmt.Println(i1)
		//can put the i1++ here
	}

	//Constructing a loop with same effects as while loop in go

	j1 := 10
	for j1 >= 0 {
		fmt.Println(j1)
		j1--
	}

	// j2 := 0
	// for {
	// 	j2++
	// }
	// fmt.Println(j2) // ifinite loop

	//For and Continue statements
	fmt.Println("For and Continue")
	for i3 := 0; i3 < 10; i3++ {
		if i3%2 != 0 {
			continue
		}
		fmt.Println(i3) //prints even numbers between 0 and ten
	}

	//For and Break statements
	fmt.Println("For and break")
	count := 0
	for m := 0; true; m++ {
		if m%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", m)
			count++
		}

		if count == 10 {
			break
		}

	}
	fmt.Println("A message after the loop")

	//Label statements
	fmt.Println("Label statement")

	people := [5]string{"Helen", "Mark", "Ayomide", "Lola", "David"}
	friends := [2]string{"Mark", "Ayomide"}

outer: //labels used to terminate outer enclosing loops, does not conflicts with identifiers thhat are not lables
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}

	fmt.Println("Next Instruction after break")

	//Goto
	fmt.Println("Goto")

	g := 0
loop: //label
	if g < 5 {
		fmt.Println(g)
		g++
		goto loop
	}

	//Swtich statements
	fmt.Println("Swtich statements")
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning python")
	case "golang":
		fmt.Println("You are learning golang")
	case "python", "Golang":
		fmt.Println("Python and Golang")

	default:
		fmt.Println("Any other language")
	}
	// use switch true{} or just switch{} for
	//Scopes in Go
	fmt.Println("Scopes in Go")

	//scope means visibility
	// file, pacakage,block(local) scope

	//import "fmt" //file scope
	// const done = false //package scoped
	//func main(){
	//	const done = true //local scoped
	//	 var task = "Running" // local scoped
	//}

}
