package main

const Y int = 3

func main() {
	var nums [5]int // [0, 0, 0, 0, 0]
	_ = nums

	days := [7]string{"월", "화", "수", "목", "금", "토", "일"}
	week := [...]string{"월", "화", "수", "목", "금", "토", "일"}
	_ = days
	_ = week

	var temps [5]float64 = [5]float64{23.2, 24.0, 25.5, 27.1, 26.3}
	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	_ = temps
	_ = matrix

	var s = [5]int{1: 10, 3: 30}
	_ = s

	b := [Y]int{1, 2, 3}
	_ = b

	for i, v := range temps {
		println(i, v)
	}
}
