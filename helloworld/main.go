/*
CLI executions:

go build main.go : compiles and creates an executable main which can be run separately
go run main.go: compiles and runs the code
go fmt main.go: formats the code
*/


/*package == workspace == project

Types of packages ->
- executable: for doing something (generates an executable file that can be executed)
- reusable: for helping something (put the helper code here)
main specifically means its an executable type
*/
package main

/*
Give my package main access to the code in the package fmt
fmt is a shorthand for format, a a standard library package with functions analogous to C's printf/scanf.
*/
import "fmt"

//executable package code should always have a main method
func main() {
	fmt.Println("Hello there!")

	/*
		~ Tidbit about variables ~
		var = creation of a variable
		variable name
		variable type (here it is string)
		value for the variable

		Like Java and unlike Python, Go is statically typed.

		We can omit using the variable type as it can be interpreted from the right-hand side.
		In that case we'd write, card := "Ace of Spades"
		Note that := is used at initialization that happens one time.
		Also, Variables can be initialized outside of a function, but cannot be assigned a variable.
	*/
	var card string = "Ace of Spades"
	fmt.Println(card)
}
