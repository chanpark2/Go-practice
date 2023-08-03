package main

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {
	println(sum(1, 2, 3))
	println(sum(4, 5, 6, 7, 8, 9))
}
