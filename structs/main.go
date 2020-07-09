package main

import "fmt"

// define the struct

// inner struct
type contactInfo struct {
	email string
	zipcode int
}

// master struct with embedding
type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main()  {

	fmt.Println("----------------------------------")

	// first way of defining a struct (uses the order of properties to define the values), not a good design
	david := person{"David", "Chucking", contactInfo{"david@gmail.com", 95111}}
	fmt.Println(david)

	// second way of defining a struct
	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
		contact: contactInfo {
			email: "alex@apple.com",
			zipcode: 96111,
		}, //every line of the struct should end with ,
	}

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	// third way of defining a struct
	var brian person
	fmt.Println(brian) // empty properties of the struct person
	brian.firstName = "Brian"
	brian.lastName = "Brandon"
	var brianContact contactInfo
	brianContact.email = "brian@rediff.com"
	brianContact.zipcode = 97111
	brian.contact = brianContact
	fmt.Println(brian)

}
