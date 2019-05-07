package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//add by myself
	sli1 := make([]int, 5)
	for idx := range sli1 {
		sli1[idx] = idx + 1
	}
	fmt.Println(sli1)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
