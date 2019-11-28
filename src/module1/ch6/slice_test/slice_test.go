package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var slice0 []int
	t.Log(len(slice0), cap(slice0))
	slice0 = append(slice0, 1)
	t.Log(len(slice0), cap(slice0))

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Log(len(slice1), cap(slice1))

	slice2 := make([]int, 5, 7)
	t.Log(len(slice2), cap(slice2))
	t.Log(slice2[0], slice2[1], slice2[2], slice2[3], slice2[4])
	slice2 = append(slice2, 1)
	t.Log(slice2[0], slice2[1], slice2[2], slice2[3], slice2[4], slice2[5])
}

func TestSliceGrowing(t *testing.T) {
	slice := []int{}
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		t.Log(len(slice), cap(slice))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	o2 := year[3:6]
	t.Log(o2, len(o2), cap(o2))

	summer := year[5:9]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "fufeng"
	t.Log(o2)
	t.Log(year)
}
