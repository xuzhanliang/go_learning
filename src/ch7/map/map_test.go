package _map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1, len(m1))
	m2 := map[string]int{}
	m2["four"] = 4
	t.Log(m2, len(m2))
	m3 := make(map[string]int, 10)
	t.Log(len(m3))
}

/**
go中判断key的值是否存在的方法
*/
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 333
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestRangeMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
