package main

import (
	"time"
	"fmt"
)

type student struct {
	name string
	age  int
}

func (s student) AddAge(age int) int {
	return s.age + age
}

func main() {

	fmt.Println(time.Now())

	s := student{name: "lijie", age: 11 }

	fmt.Println(s)

	age := s.AddAge(11)

	fmt.Println("new age: ", age)

	fmt.Println(s.age)
}
