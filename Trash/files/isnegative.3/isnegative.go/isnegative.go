package main

import "github.com/01-edu/z01"

func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')
}
func main() {
	IsNegative(-5) // Output: T
	IsNegative(0)  // Output: F
	IsNegative(42) // Output: F
}
