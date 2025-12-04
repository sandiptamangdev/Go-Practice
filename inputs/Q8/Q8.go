// Ask for two numbers and add them.

package main
import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Enter first number:")
	fmt.Scan(&x);
	fmt.Println("Enter second number:")
	fmt.Scan(&y)

	var z int
	z = x + y;
	fmt.Println(z)
}