package main

import (
	"fmt"

	"cmp"
)

// 1. Generic Types
type Pair[T any, U any] struct {
	First  T
	Second U
}

type Number interface {
	float64 | int
}

func Add[T Number](a, b T) T {
	return a + b
}

// 2. Generic Functions
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// 3. Generic Variables
// Go does not support standalone generic variables. Use generics with types or functions.

// 4. Generic Returns
func FirstElement[T any](s []T) (T, error) {

	if len(s) == 0 {
		var zero T
		return zero, fmt.Errorf("slice is empty")
	}

	return s[0], nil
}

// 5. Constraints
type Adder[T any] interface {
	Add(a, b T) T
}

func Sum[T Adder[T]](a, b T) T {
	return a.Add(a, b)
}

// Using built-in constraints for ordered types
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 6. Instantiating Generics
func main() {
	p := Pair[int, string]{First: 1, Second: "hello"}
	fmt.Println(p)

	k := Add(1, 2)
	_ = k

	PrintSlice([]int{1, 2, 3})

	nums := []int{10, 20, 30}

	first, err := FirstElement(nums)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("First element:", first)
	}

	fmt.Println("Max of 5 and 10:", Max(5, 10))
}
