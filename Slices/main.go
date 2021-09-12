package main

import "fmt"

func main() {
	fmt.Println("Hello!, Slices World")

	//Differences between Arrays and Slices
	// 	● Has a fixed length defined at compile
	// time;
	// ● The length of an array is part of its type,
	// defined at compile time and cannot be
	// changed;
	// ● By default an uninitialized array has all
	// elements equal to zero;
	// ● Has a dynamic length (it can shrink or
	// grow);
	// ● The length of a slice is not part of its
	// type and it belongs to runtime;
	// ● An uninitialized slice is equal to nil (its
	// zero value is nil).

	//Declaring Slices and Basic Slice operations

	var cities []string
	//cities[0] = "Abeokuta"
	cities = append(cities, "Ayo", "Lola")
	fmt.Println(cities)               // [Ayo Lola]
	fmt.Println(cities == nil)        //true
	fmt.Printf("Type: %#v\n", cities) // Type: []string(nil)

	numbers := []int{2, 3, 4, 5} //Slices literals
	_ = numbers
	nums := make([]int, 2) // {0,0}
	_ = nums

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)
	fmt.Println("Friend is:", friends[0])
	friends[0] = "David"
	fmt.Println("Friend is:", friends[0])

	for index, value := range numbers {

		fmt.Printf("index: %v, value %v\n", index, value)
	}

	//Comparing Slices
	fmt.Println("Comparing Slices")

	var n []int
	fmt.Println(n == nil) //true

	m := []int{}
	fmt.Println(m == nil) //false

	//Appending to a Slice and Copying Slices
	fmt.Println("Appending to a Slice and Copying Slices")

	newSlice := []int{}
	newSlice = append(newSlice, 10, 20, 30)
	fmt.Println(newSlice) //[10 20 30]

	anotherslice := []int{100, 200}
	newSlice = append(newSlice, anotherslice...)
	fmt.Println(newSlice) //[10 20 30 100 200]

	//Copying slice
	src := []int{10, 20, 30, 40}
	dst := make([]int, 2)
	nn := copy(dst, src)                             //The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	fmt.Println("src:", src, "dst:", dst, "nn:", nn) //src: [10 20 30 40] dst: [10 20] nn: 2

	//Slice Expressions
	fmt.Println("Slice Expressions")

	s1 := []int{1, 2, 3, 4}
	s2 := s1[1:4]
	fmt.Println(s2)     // [2 3 4]
	fmt.Println(s1[1:]) // [2 3 4] same as s1[1:len(S1)]
	fmt.Println(s1[:3]) // [1 2 3] same as s1[0:3]
	fmt.Println(s1[:])  // [1 2 3 4] same as s1[0:len(s1)]

	//Slice's Internal Backing Array and Slice Header
	fmt.Println("Slice's Internal Backing Array and Slice Header")
	fmt.Println(s1) //[1 2 3 4]
	s3, s4 := s1[0:2], s1[1:3]
	s3[1] = 600
	fmt.Println(s1) //[1 600 3 4]
	fmt.Println(s4) //[600 3]
	// s3 modified both s1 and s4

	cars := []string{"Ford", "Dodge", "Audi", "Honda"}
	newCars := []string{}

	newCars = append(newCars, cars...)
	cars[0] = "Nissan"
	//only cars will be modified
	fmt.Println(cars, newCars) //[Nissan Dodge Audi Honda] [Ford Dodge Audi Honda]

}
