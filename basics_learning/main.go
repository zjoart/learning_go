package main

import (
	"fmt"
	"strconv"
)

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go World!")
	distance := 60.8
	fmt.Println("The seconds in hour is :", secondsInHour)
	fmt.Printf("The distance in miles is %f \n", distance*0.62137)

	//declaring variales

	var age int = 20
	var name string = "lola"
	//a variable is set to 0 or empty string if not initialized
	var i int
	var j string
	var k bool

	fmt.Println("i is", i)
	fmt.Println("j is", j)
	fmt.Println("k is", k)
	fmt.Println("Age is:", age)
	fmt.Println("Name is:", name)

	name1 := name
	fmt.Println("New name is:", name1)

	dept := "Physics"
	fmt.Println("Department is:", dept)
	fmt.Println("Age is", age, "Name is:", name)
	//fmt formatters
	shape := "Circle"
	radius := 4
	pi := 3.142
	isdone := false

	fmt.Printf("Is it done: %t\n", isdone) // %t shows boolean
	// %s - String, % d - Decimal, %f - floating value(.(to decimal places eg %.3f))
	// \n to print in a new line
	fmt.Printf("The shape is %s with radius %d and pi constant %f\n", shape, radius, pi)
	fmt.Printf("Shape: %q\n", shape)                                        // %qoute strings
	fmt.Printf("shape is: %T, radius is %T, pi is %T\n", shape, radius, pi) // %T shows variable type
	fmt.Printf("radius is %b in binary\n", radius)                          // %b coverts decimal to binay or base 2

	//Constants

	const daysOfWeek int = 7 // const must be initalized
	fmt.Println(daysOfWeek)
	/* Constants rules
	1.you cannot change a constant
	2.you cannot initate a constant at runtume
	const power = math.Pow(2,3)
	3. you cannot use a variable to initialize a constant
	t:= 5
	const tc = t// ERROR

	*/

	//Untyped Constants

	const a float64 = 5.1 //type constant
	const b = 6.7         // untyped constant
	_, _ = a, b

	//iota
	const (
		c1 = iota
		c2 = iota
		c3 = iota
		c4
		c5
	)
	fmt.Println(c1, c2, c3, c4, c5) //OUTPUT : 0 1 2 3 4
	// Go datatypes
	var i1 int8 = 127 // INT TYPE
	_ = i1
	var f1 float64 = -3.2 // Float Datatype
	_ = f1
	var r1 rune = 'f' // Rune type
	_ = r1
	var b1 bool = false // Bool type
	_ = b1
	var s1 string = " Hello World" // String Typw
	_ = s1
	var numbers = [4]int{4, 5, -9, 100} //Array type
	_ = numbers
	var cities = []string{"London", "Kenya", "Abeokta"} // Slice type
	_ = cities
	fmt.Printf("cities: %v\n", cities)

	balances := map[string]float64{
		"usd": 2332.2,
		"eur": 511.11,
	} // Map type
	_ = balances

	fmt.Printf("balances: %v\n", balances["eur"])

	type Person struct {
		name string
		age  int
	} // Struct type

	var you Person = Person{"ayo", 5}

	fmt.Printf("you: %v\n", you.age)

	var x int = 2
	ptr := &x // Pointer type
	_ = ptr

	// Function type
	fmt.Printf("%T\n", f)

	//OPERATORS IN GO
	sa, sb := 4, 2

	r := (sa + sb) / (sa - sb) * 2
	fmt.Println("The solution is:", r) // 6

	// Asignment operators
	sa += sb // increment assignment
	fmt.Println(sa)

	sb -= 1 // Decrment assignment
	fmt.Println(sb)

	sb *= 10 // multiplication assignment
	fmt.Println(sb)

	sb /= 6 // Division assignment
	fmt.Println(sb)

	sa %= 4 // Modulus assignment
	fmt.Println(sa)

	sa++ // increment assignment
	fmt.Println(sa)

	sa-- // decrement assignment
	fmt.Println(sa)

	// Comparison and Logical operators
	ca, cb := 5, 10
	fmt.Println(ca == cb) //false
	fmt.Println(ca != cb) // true
	// >,<, <=,>=, &&(And), ||(or), !(Not)

	// Overflow and Underflow of Integers
	var ox uint8 = 255
	ox++            //overflow
	fmt.Println(ox) //0

	//Converting Numeric types

	var (
		nx = 3
		ny = 3.1
	)
	nx = nx * int(ny)
	fmt.Println(nx)
	// foat64()for float

	//Converting number to strings
	var ss = fmt.Sprint(99.4) //or fmt.sprintf()
	fmt.Println(ss + "g")

	//converting string to float
	s12 := "3.123"
	var f12, err = strconv.ParseFloat(s12, 64)
	_ = err
	fmt.Println(f12)
	//Named and Defined types

	type speed int

	var ss1 speed = 10
	var ss2 speed = 20
	_ = ss2
	fmt.Printf("%T\n", ss1)

	var ss3 int = 2

	ss4 := ss1 / speed(ss3)
	fmt.Println(ss4)

	//Aliases
	var aaa uint8 = 10
	var bbb byte = aaa // uint8 and bytes are of same type so alias
	fmt.Println(bbb)

}

func f() {

}
