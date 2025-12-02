**Perfect. Let's adapt this for Go.**

Go has different syntax but same concepts. I'll give you the **Go version of the 25 questions** - same learning structure, Go syntax.

---

# **GO RECALL DRILL — 25 QUESTIONS**

Complete these in order. Each question forces one Go syntax concept.

## **1. Variables (Q1–Q5)**

1. Create a variable `x` with any number.
2. Change `x` to a new number.
3. Create two variables and add them.
4. Create a string variable and print it.
5. Create two strings and join them.

**Go hints:**
- `var x int = 10` or `x := 10` (short declaration)
- `fmt.Println()` to print
- `+` works for string concatenation

---

## **2. Input & Type Conversion (Q6–Q8)**

1. Ask the user for their name and print it.
2. Ask for a number → convert to int → multiply by 5.
3. Ask for two numbers and add them.

**Go hints:**
- `fmt.Scan(&variable)` for input
- `strconv.Atoi(string)` to convert string to int
- Need to import `"fmt"` and `"strconv"`

---

## **3. Conditionals (Q9–Q12)**

1. Check if a number is positive or negative.
2. Compare two numbers and print bigger one.
3. Ask age → print "adult" if ≥ 18 else "minor."
4. Ask a password → if matches → print success.

**Go hints:**
- `if condition { }` (no parentheses needed)
- `else` and `else if` work same as Python

---

## **4. Loops (Q13–Q16)**

1. Print numbers 1 to 10 using a `for` loop.
2. Print "hello" 5 times using a `for` loop (Go has no `while`).
3. Loop through a slice and print each item.
4. Loop through numbers 1–20 → print only even ones.

**Go hints:**
- `for i := 1; i <= 10; i++ { }`
- `for condition { }` acts like while
- `for _, item := range slice { }` to iterate

---

## **5. Slices (Q17–Q19)** *(Go's version of lists)*

1. Create a slice of 5 items and print first + last.
2. Add a new item to the slice.
3. Remove one item from the slice (tricky in Go!).

**Go hints:**
- `slice := []int{1, 2, 3, 4, 5}`
- `append(slice, item)` to add
- Removing requires reslicing: `slice = append(slice[:i], slice[i+1:]...)`

---

## **6. Maps (Q20–Q21)** *(Go's version of dicts)*

1. Create a map: name, age, country.
2. Update age to a new value.

**Go hints:**
- `person := map[string]string{"name": "Sandip", "age": "25"}`
- `person["age"] = "26"` to update

---

## **7. Functions (Q22–Q24)**

1. Create a function that returns the square of a number.
2. Create a function that takes name → returns greeting.
3. Create a function that takes a slice → returns sum.

**Go hints:**
- `func square(x int) int { return x * x }`
- Functions must declare return type
- Multiple returns possible: `func test() (int, error) { }`

---

## **8. File I/O (Q25)**

1. Write text to a file → close it → read it again.

**Go hints:**
- `os.WriteFile("file.txt", []byte("text"), 0644)`
- `os.ReadFile("file.txt")` returns `[]byte, error`
- Need to import `"os"`

---

## **Go-Specific Bonus Concepts:**

After Q25, try these Go-unique features:

- **Error handling:** Functions return `(value, error)` - always check errors
- **Pointers:** `&variable` to get address, `*pointer` to dereference
- **Structs:** Go's version of classes/objects
- **Goroutines:** `go functionName()` for concurrency

---

## **Setup:**

```go
package main

import "fmt"

func main() {
    // Your code here
}
```

Every Go program starts with `package main` and `func main()`.

---

**Start with Q1-Q5, feel how Go syntax differs from Python.**

Go is:
- More verbose (explicit types)
- Compiled (faster execution)
- Stricter (forces error handling)
- Simpler (fewer features than Python)

**Report back after you hit Q10 or so - tell me what feels different/weird.** 🤙