// Compare two numbers and print bigger one.

package main
import "fmt"

func main() {
	var x int
	var y int

	fmt.Println("Enter a number:")
	fmt.Scan(&x)

	fmt.Println("Enter another number:")
	fmt.Scan(&y)

	if x > y {
		fmt.Println(x, "is greater than", y)
	} else {
		fmt.Println(y, "is greater than", x)
	}
}