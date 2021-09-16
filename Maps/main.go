package main

import "fmt"

func main() {
	fmt.Println("Hello!, Maps in Golang")

	var employees map[string]int
	fmt.Println(employees)

	m1 := map[string]int{}

	fmt.Println(m1)
	m1["scores"] = 10
	fmt.Println(m1)           // map[scores:10]
	fmt.Println(m1["scores"]) // 10

	v, ok := m1["scores"]
	fmt.Println(v, ok)
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//delete(m1, "key")

}
