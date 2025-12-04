// Ask age and print "adult" if >= 18 else "minor".

package main
import "fmt"

func main() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}