package main

import "fmt"

func main() {
	a := 10

	switch a {
	case 10:
		fmt.Println("a == 10")
	case 20:
		fmt.Println("a == 20")
		fallthrough
	default:
		fmt.Println("a != 10 && a != 20")
	}
}
