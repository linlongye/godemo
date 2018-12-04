package main

import (
	"OOP/model"
	"fmt"
)

func main() {
	/*	s := Student{
			Name:  "hl",
			Age:   20,
			Score: 100,
		}
		fmt.Println(s)
		var p = new(Student)
		p.Score = 100
		fmt.Println(*p)

		var p2 = &Student{Name: "ss", Age: 20, Score: 300}

		fmt.Println(p2.GetName())
		p2.Calc()
		fmt.Println(*p2)*/

	s := model.Student{"aa", 20}
	fmt.Println(s)
	fmt.Println(s.GetName())
	st := model.GetPerson("lll", 30)
	fmt.Println(st.GetPersonName())
}

/*
type Student struct {
	Name  string
	Age   int
	Score float64
}

func (stu Student) GetName() string {
	//fmt.Println(stu.Name)
	return stu.Name


}

func (stu Student) Calc() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(res)
}
*/
