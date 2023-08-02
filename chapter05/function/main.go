package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func DivideWithName(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c := Add(1, 2)
	fmt.Println(c)

	d, ok := Divide(9, 3)
	fmt.Println(d, ok)
	d, ok = Divide(9, 0)
	fmt.Println(d, ok)

	f := 2
	fmt.Println(f)
	i, ok := DivideWithName(9, 3)
	fmt.Println(i, ok)
	i, ok = DivideWithName(9, 0)
	fmt.Println(i, ok)
}
