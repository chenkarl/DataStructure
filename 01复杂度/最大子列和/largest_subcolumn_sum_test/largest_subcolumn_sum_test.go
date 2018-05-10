package largest_subcolumn_sum

import (
	"testing"

	"../largest_subcolumn_sum"
)

func TestLargest_subcolumn_sum_test1(t *testing.T) {
	t.Log("Start Test 1")
	testdata1 := []int{-2, 11, -4, 13, -5, -2}
	mostsum := largest_subcolumn_sum.Get(testdata1)
	if mostsum == 20 {
		t.Log("Test 1 : Success")
	} else {
		t.Log("Test 1 : Failed")
	}
}
func TestLargest_subcolumn_sum_test2(t *testing.T) {
	t.Log("Start Test 2")
	testdata1 := []int{-2, 11, -4, 13, -5, -2}
	mostsum := largest_subcolumn_sum.Get(testdata1)
	if mostsum == 20 {
		t.Log("Test 2 : Success")
	} else {
		t.Log("Test 2 : Failed")
	}
}
