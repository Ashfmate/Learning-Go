package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	sum = 0
	for i, num := range nums {
		if i != 3 {
			sum += num
		}
	}
	fmt.Println(sum)
}
