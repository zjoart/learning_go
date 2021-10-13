package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in Golang")

	type Car struct {
		year  int
		maker string
		model string
		speed int
	}

	newCar := Car{2000, "Toyota", "Camry", 230}
	snewCar := Car{2000, "Toyota", "Camry", 230}

	fmt.Println(newCar)

	type SecondCar struct {
		year  int
		maker string
		model string
		speed int
	} //best format to create struct
	secondCar := SecondCar{year: 000, maker: "Toyota", model: "Camry", speed: 230}

	fmt.Println(secondCar)

	//Retrieving and Updating Struct Field

	fmt.Println(secondCar.year) //retrieving
	secondCar.year = 2000       //updating
	fmt.Println(secondCar.year) //retrieving back

	//comparing structs
	// using == and !=

	fmt.Println(newCar == snewCar)

	//	Anonymous Structs and Anonymous Struct Fields

	ayo := struct {
		firtname string
		age      int
	}{
		firtname: "ayo",
		age:      12,
	}

	fmt.Println(ayo)
	//fields
	type Book struct {
		string
		float64
		bool
	}
	book := Book{"Name", 200, false}
	fmt.Println(book)

	//Embedded Structs

	type Contact struct {
		email, address string
		phone          int
	}

	type Social struct {
		contact    Contact
		socialType string
	}

	josh := Social{
		socialType: "Fb",
		contact: Contact{
			email:   "Admin@gmail.com",
			address: "Abeokuta",
			phone:   00022,
		},
	}
	fmt.Printf("%+v \n", josh)
	//{contact:{email:Admin@gmail.com address:Abeokuta phone:18} socialType:Fb}
	fmt.Println(josh.contact.address)
}
