package main

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}
	withdrawFunc(a, 30)
	a.withdrawMethod(30)

	println(a.balance)
}
