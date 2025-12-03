// package main
// import "fmt"

// func main() {
// 	var x int = 10
// 	fmt.Println(x)
// }



// Variables

// Create a variable `x` with any number.

// package main
// import "fmt"

// func main() {
// 	var x int= 10
// 	fmt.Println(x)
// }


// Change `x` to a new number

// package main
// import "fmt"

// func main(){
// 	var x int = 5	
// 	fmt.Println(x)
// 	x = 20
// 	fmt.Println(x)
// }

// Create two variables and add them.

// package main
// import "fmt"

// func main() {
// 	var x int = 5
// 	var y int = 10

// 	var z int = x + y
// 	fmt.Println(z)
// }



// Create a string variables and print it.

// package main
// import "fmt"

// func main() {
// 	// var student1 string = "John"
// 	student1 := "John"
// 	// var student2 string = "Hero"
// 	student2 := "Hero"
// 	// var x int = 5
// 	x := 5

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }


// Create two strings and join them.

// package main
// import "fmt"

// func main() {
// 	name := "John"
// 	greeting := "Welcome "

// 	hello := greeting + name

// 	fmt.Println(hello)
// }



// Input & Type Conversion 

// Ask the user for their name and print it

// package main
// import "fmt"

// func main() {
// 	var x string
// 	fmt.Scan(&x)
// 	fmt.Println(x)
// }

// Ask for a number, convert to int, multiply by 5

package main
import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x = x * 5
	fmt.Println(x)
}

