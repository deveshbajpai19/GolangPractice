package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main()  {

	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
	}

	alex.updateFirstName1("Alexis")
	fmt.Println(alex) //update is not reflected, as the struct is passed as value to the method

	(&alex).updateFirstName2("Alexis")
	fmt.Println(alex) //update is reflected, as the struct is passed as reference to the method


}

// classic pass by value usage
func (p person) updateFirstName1(newFirstName string) {

	p.firstName = newFirstName

}

// pass by reference using the pointer to the struct
func(personPtr *person) updateFirstName2(newFirstName string) {

	(* personPtr).firstName = newFirstName

}
