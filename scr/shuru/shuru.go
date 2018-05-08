package shuru

import (
	"fmt"
)

// Get 获取输入
func Get() []int {
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
	return slice
}
