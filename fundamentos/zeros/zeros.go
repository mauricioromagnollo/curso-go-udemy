package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	// Valores default
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e) // 0 0 false "" <nil>
}
