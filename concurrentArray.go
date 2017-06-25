package main

import (
	"sync/atomic"
	"fmt"
	"errors"
)

type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}

type concurrentArray struct {
	length uint32
	val    atomic.Value
}

func (arr *concurrentArray) Set(index uint32, elem int) (err error) {

	if err = arr.checkIndex(index); err != nil {
		return err
	}

	if err = arr.checkValue(); err != nil {
		return err
	}

	newArray := make([]int, arr.length)

	copy(newArray, arr.val.Load().([]int))

	newArray[index] = elem

	arr.val.Store(newArray)

	return nil
}

func (arr *concurrentArray) Get(index uint32) (elem int, err error) {

	if err = arr.checkIndex(index); err != nil {
		return
	}

	if err = arr.checkValue(); err != nil {
		return
	}

	elem = arr.val.Load().([]int)[index]

	return
}

func (arr *concurrentArray) Len() uint32 {
	return arr.length
}

func (arr *concurrentArray) checkIndex(index uint32) error {

	if index >= arr.length {
		return fmt.Errorf("Index out of range [0, %d]", arr.length)
	}

	return nil
}

func (arr *concurrentArray) checkValue() error {

	v := arr.val.Load()

	if v == nil {
		return errors.New("Invalid int array")
	}

	return nil
}

func NewConcurrentArray(length uint32) ConcurrentArray {

	arr := concurrentArray{}

	arr.length = length

	arr.val.Store(make([]int, arr.length))

	return &arr
}

func main() {

	arr := NewConcurrentArray(100)

	fmt.Println(arr)

	arr.Set( uint32(1), 100 )
	arr.Set( uint32(0), 120 )

	fmt.Println( arr.Get( 1 ) )
	fmt.Println( arr.Get( 0 ) )

	//arr.Set(1, 100)
	//
	//fmt.Println(arr)
}
