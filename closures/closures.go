package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//add by myself
func test(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := initSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := initSeq()
	fmt.Println(newInts())

	mytest := test(5)
	fmt.Println(mytest())
	fmt.Println(mytest())
	fmt.Println(mytest())

	mytestNew := test(5)
	fmt.Println(mytestNew())
}
