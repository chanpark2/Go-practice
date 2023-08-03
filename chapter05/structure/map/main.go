package main

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}

	// 키가 16인 요소를 삭제한다.
	delete(m, 16)

	// 키가 46인 요소가 있는지 확인한다.
	if item, ok := m[46]; ok {
		println(item.Name, item.Price)
	}

	for k, v := range m {
		println(k, v.Name, v.Price)
	}
}
