// Ask a password and if it matches then print success.

package main
import "fmt"

func main() {
	var x string = "password123"
	var y string

	fmt.Println("Enter your password:")
	fmt.Scan(&y)

	if x == y {
		fmt.Println("Success")
	} else {
		fmt.Println("Retry")
	}
}