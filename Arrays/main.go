package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello!, Arrays in Golang")

	// Arrays in Go
	fmt.Println("Arrays")
	//arrays has a fixed length
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	//Declaraing arrays

	var a1 = [4]float64{} // assigns 0000 - [0 0 0 0]
	fmt.Println(a1)
	var a2 = [3]int{-10, 1, 100}
	fmt.Println(a2)
	a3 := [4]string{"Ola", "Ayo", "Debola", "Victor"}
	fmt.Println(a3)

	a4 := [4]string{"x", "y"}
	fmt.Println(a4)

	//ellipsis operator
	a5 := [...]int{1, 2, 14, -10, 30, 6}
	fmt.Println(a5)
	fmt.Println("length:", len(a5))

	//Array Operations

	nums := [3]int{10, 20, 30}
	fmt.Println("Numbers:", nums) // [10 20 30]
	nums[0] = 7
	fmt.Println("Modified to:", nums) // [7 20 30]
	//interation in arrays with range keyword
	for i, v := range nums { // i is index v is the values
		fmt.Println("Index is:", i, "Value is:", v)
	}

	// interation with for loop
	fmt.Println(strings.Repeat("*", 25))
	for k := 0; k < len(nums); k++ {
		fmt.Println("Index:", k, "Value:", nums[k])
	}
	fmt.Println(strings.Repeat("*", 25))
	//multi directional arrays
	bals := [2][3]int{
		{5, 6, 7}, {8, 9, 3},
	}
	fmt.Println(bals)
	fmt.Println(strings.Repeat("*", 25))

	//Arrays with Keyed Elements
	grades := [3]int{1: 10, 0: 5, 2: 7}
	fmt.Println(grades) // [5 10 7]
	accounts := [3]int{2: 50}
	fmt.Println(accounts) // [0 0 50]
	nams := [...]string{5: "Daniel"}
	fmt.Println(nams) //[     Daniel]
	//an unkeyed elements gets it index from the last key index

}
