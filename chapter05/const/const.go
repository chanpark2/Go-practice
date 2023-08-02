package main

import "fmt"

const (
	Red int = iota
	Blue
	Green
)

const (
	BitFlag1 uint = 1 << iota
	BitFlag2
	BitFlag3
	BitFlag4
)

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	const pi1 float64 = 3.14159265358979323846264338327950288419716939937510582097494459
	var pi2 = 3.14159265358979323846264338327950288419716939937510582097494459

	fmt.Println(pi1, pi2)

}
