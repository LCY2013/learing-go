package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	m1[4] = func(op int) int { return op * op * op * op }
	t.Log(m1[1](2), m1[2](2), m1[3](2), m1[4](2))
}

func TestMapForSet(t *testing.T) {
	mapSet := map[int]bool{}
	mapSet[1] = true
	n := 1
	if mapSet[n] {
		t.Logf("%d is exiting", n)
	} else {
		t.Logf("%d is not exiting", n)
	}
	t.Log("mapSet len is ", len(mapSet))

	delete(mapSet, 1)
	t.Log(mapSet)
}
