package largest_subcolumn_sum

func Get(l []int) int {
	sum := 0
	mostsum := 0
	for _, value := range l {
		sum += value
		if sum > mostsum {
			mostsum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return mostsum
}
