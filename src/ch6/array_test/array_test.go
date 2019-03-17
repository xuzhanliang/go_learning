package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{2, 4, 6, 8}
	t.Log(arr[0], arr[1])
	t.Log(arr1[1])
	t.Log(arr2[3], len(arr2))
}
