package concurrent_patterns

import "testing"

func TestInitMap(t *testing.T) {
	map1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(map1[3])
	t.Logf("map1 len is %d", len(map1))

	map2 := map[int]int{}
	map2[2] = 22
	t.Logf("map2 len is %d", len(map2))

	map3 := make(map[int]int, 10)
	t.Logf("map3 len is %d", len(map3))
}

func TestAccessNotExistingKey(t *testing.T) {
	map1 := map[int]int{}
	t.Log(map1[1])

	map1[1] = 0
	t.Log(map1[1])

	if value, ok := map1[13]; ok {
		t.Logf("key is exiting, value is %d", value)
	} else {
		t.Log("key is not exiting")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 3: 3, 4: 4}
	for key, value := range m {
		t.Log(key, value)
	}
}
