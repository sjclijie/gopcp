package main

func main() {

	//new(T)分配了零值填充的T类型的内存空间，并返回一个指针
	//make(T,args)返回一个有初始值的T类型引用，而不是*T

	//引用在使用前必须被初始化, 例如slice，是一个包含指向数据的指针，长度和容量的三项描述符,在这些项目被初始化前，slice为nil

	b := make([]int, 5, 10) // len = 5, cap = 10
	b = b[:cap(b)]          // len = 5, cap = 5
	b = b[1:]               //len = 4 cap = 4

}
