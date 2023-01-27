package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	mp1 := map[int]string{1: "1", 2: "2", 3: "3"}
	keys1 := Keys(mp1)

	assert.Equal(t, len(mp1), len(keys1))
	for _, k := range []int{1, 2, 3} {
		assert.Contains(t, keys1, k)
	}

	mp2 := map[struct{ name string }]int{
		{"foo"}: 10,
		{"bar"}: 20,
	}
	keys2 := Keys(mp2)
	assert.Equal(t, len(mp2), len(keys2))
	for _, k := range []struct{ name string }{{"foo"}, {"bar"}} {
		assert.Contains(t, keys2, k)
	}
}

func TestValues(t *testing.T) {
	mp1 := map[int]string{1: "1", 2: "2", 3: "3"}
	values1 := Values(mp1)

	assert.Equal(t, len(mp1), len(values1))
	for _, v := range []string{"1", "2", "3"} {
		assert.Contains(t, values1, v)
	}

	mp2 := map[struct{ name string }]int{
		{"foo"}: 10,
		{"bar"}: 20,
	}
	values2 := Values(mp2)

	assert.Equal(t, len(mp2), len(values2))
	for _, v := range []int{10, 20} {
		assert.Contains(t, values2, v)
	}
}

func TestEntries(t *testing.T) {
	mp := map[int]string{1: "1", 2: "2", 3: "3"}
	entries := Entries(mp)

	assert.Equal(t, len(mp), len(entries))
	for _, e := range []MapEntry[int, string]{{1, "1"}, {2, "2"}, {3, "3"}} {
		assert.Contains(t, entries, e)
	}
}
