package main

import (
	"fmt"
)

type Data struct {
	value int
	data  [200]int
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{name, age}
}

func ChangeData(arg *Data) {
	(*arg).value = 999
	(*arg).data[100] = 999
}
func main() {
	var a int = 10
	var pa *int = &a

	fmt.Println(pa)
	fmt.Println(*pa)

	*pa = 20
	fmt.Println(a)

	var p1 *int = &a
	var p2 *int = &a
	fmt.Println(p1 == p2)

	var data Data
	ChangeData(&data)
	fmt.Println("value:", data.value)
	fmt.Println("data[100]:", data.data[100])

	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
