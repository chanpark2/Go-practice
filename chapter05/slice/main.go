package main

import "fmt"

func addNum(slice *[]int) {
	*slice = append(*slice, 4)
}

func main() {
	var slice1 = []int{1, 2, 3}        // 대괄호 안에 길이를 지정하지 않음
	var slice2 = []int{1, 5: 2, 10: 3} // 5번째 요소와 10번째 요소만 초기화 나머지는 0으로 초기화

	addNum(&slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := append(slice1, 4, 5) // slice1에 4, 5를 추가한 새로운 슬라이스를 생성
	fmt.Println(slice3)

	slice2 = append(slice2, slice3...) // slice2에 slice3의 모든 요소를 추가
}
