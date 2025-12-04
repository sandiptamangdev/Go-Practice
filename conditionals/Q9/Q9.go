// Check if a number is positive or negative.

package main
import "fmt"

func main() {
	var x int
	fmt.Println("Enter a positive or negative number:")
	fmt.Scan(&x) 

	if x < 0 {
		fmt.Println("Negative Number")
	} else if x == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive Number")
	}
}