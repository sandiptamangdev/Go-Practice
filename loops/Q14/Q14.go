// Print "hello" 5 times using a `for` loop (Go has no `while`).

package main
import "fmt"

func main() {
	var x string = "hello"

	for i := 1; i <= 5; i++ {
		fmt.Println(x);
	}
}
