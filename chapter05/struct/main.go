package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

type User struct {
	Name string
	ID   string
	Age  int
}

type Person struct {
	A int8
	B int8
	C int
	D int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	var house House
	house.Address = "서울시 강남구"
	house.Size = 28
	house.Price = 9.8
	house.Category = "아파트"

	fmt.Println(house)

	var house2 House = House{"서울시 강남구", 28, 9.8, "아파트"}
	fmt.Println(house2)

	house3 := House{
		Address:  "서울시 강남구",
		Size:     28,
		Price:    9.8,
		Category: "아파트", // 마지막 요소에도 콤마를 붙여야 합니다.
	}
	fmt.Println(house3)

	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}
	fmt.Println(user)
	fmt.Println(vip)
}
