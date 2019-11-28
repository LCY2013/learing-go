package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr, arr1)
	t.Log(arr[1], arr1[1])
	arr2 := [...]int{1, 3, 4, 2, 5, 6, 7}
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for index, value := range arr {
		t.Log(index, value)
	}

	for _, value := range arr {
		t.Log(value)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	arr_sec := arr[3:]
	arr_sec1 := arr[:3]
	arr_sec2 := arr[1:3]
	arr_sec3 := arr[:]
	t.Log(arr_sec, arr_sec1, arr_sec2, arr_sec3)
}
