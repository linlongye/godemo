package model

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() string {
	return s.Name
}

type person struct {
	name string
	age  int
}

func (s person) GetPersonName() string {
	return s.name
}

type Account struct {
	AccountNo string
	Balance   float64
	Pwd       string
}

func (account *Account) Despite(pwd string, money float64) bool {
	if pwd != account.Pwd {
		return false
	}
	if money <= 0 {
		return false
	}
	account.Balance += money
	return true
}

func (account *Account) WithDraw(pwd string, money float64) bool {
	if pwd != account.Pwd {
		return false
	}
	if money <= 0 {
		return false
	}
	account.Balance -= money
	return true
}

func (account *Account) Query(pwd string) float64 {
	if pwd != account.Pwd {
		return 0
	}
	return account.Balance
}
