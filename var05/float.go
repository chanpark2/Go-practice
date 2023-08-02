package main

import "fmt"

func main() {
	var a float32 = 1234.14
	var b float32 = 2456.14
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a, b, c, d)
}
