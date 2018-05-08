package main

import (
	"fmt"
)

func main() {
	var count int
	var shit string
	slice := []int{}
	fmt.Scanf("%d", &count)
	fmt.Scanf("%s", &shit)
	for i := 0; i < count; i++ {
		var num int
		fmt.Scanf("%d", &num)
		slice = append(slice, num)
	}
	sum := 0
	mostsum := 0
	for _, value := range slice {
		sum += value
		if sum > mostsum {
			mostsum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	fmt.Printf("%d", mostsum)
}
