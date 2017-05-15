package main

import (
	"errors"
	"fmt"
)

type binaryOperaton func(operand1 int, operand2 int) (result int, err error)

func operate(op1 int, op2 int, bop binaryOperaton) (result int, err error) {

	if bop == nil {
		err = errors.New("Invalid binary operation function")
		return
	}

	return bop(op1, op2)
}

type myInt int

func (i myInt) add(anthor int) int {

	i = i + myInt(anthor)

	return int(i)
}

func (i *myInt) add2(anthor int) myInt {

	// *myInt表示指针类型
	// i 表示一个指针
	// *i 表示指针 i 指向的值

	*i = *i + myInt(anthor)

	return *i
}

func main() {

	i1 := myInt(1)

	i2 := i1.add(2)

	fmt.Println(i1, i2)

	fmt.Println("===================")

	i3 := myInt(1)

	i4 := i3.add2(2)

	fmt.Println(i3, i4)
}
