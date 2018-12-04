package model

func GetPerson(name string, age int) *person {
	return &person{name: name, age: age}
}
