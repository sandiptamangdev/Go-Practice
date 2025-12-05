// Loop through numbers 1-20 => print only even ones.

package main
import "fmt"

func main() {
	// var x int
	for i := 0; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}