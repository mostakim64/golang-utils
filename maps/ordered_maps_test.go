package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap_Set(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(1, "one")

	value, exists := om.Get(1)
	assert.True(t, exists)
	assert.Equal(t, "one", value)
}

func TestOrderedMap_Get(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	value, exists := om.Get(1)
	assert.False(t, exists)
	assert.Empty(t, value)

	om.Set(1, "one")

	value, exists = om.Get(1)
	assert.True(t, exists)
	assert.Equal(t, "one", value)
}

func TestOrderedMap_Delete(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(1, "one")

	om.Delete(1)

	_, exists := om.Get(1)
	assert.False(t, exists)
}

func TestOrderedMap_ReorderKeys_Uint(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(3, "three")
	om.Set(1, "one")
	om.Set(4, "four")
	om.Set(2, "two")

	om.ReorderKeys(true)
	keys := om.Keys()
	expectedKeys := []uint{1, 2, 3, 4}
	assert.Equal(t, expectedKeys, keys)

	om.ReorderKeys(false)
	keys = om.Keys()
	expectedKeys = []uint{4, 3, 2, 1}
	assert.Equal(t, expectedKeys, keys)
}

func TestOrderedMap_ReorderKeys_String(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("apple", 5)
	om.Set("cherry", 3)
	om.Set("banana", 1)

	om.ReorderKeys(true)
	keys := om.Keys()
	expectedKeys := []string{"apple", "banana", "cherry"}
	assert.Equal(t, expectedKeys, keys)

	om.ReorderKeys(false)
	keys = om.Keys()
	expectedKeys = []string{"cherry", "banana", "apple"}
	assert.Equal(t, expectedKeys, keys)
}

func TestOrderedMap_BulkSet(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	data := map[uint]string{
		1: "one",
		2: "two",
		3: "three",
	}

	om.BulkSet(data)

	value, exists := om.Get(2)
	assert.True(t, exists)
	assert.Equal(t, "two", value)
}

func TestOrderedMap_Print(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	om.Print()
	// Output:
	// Key: 1, Value: one
	// Key: 2, Value: two
	// Key: 3, Value: three
}

func TestOrderedMap_ToJSON(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	jsonData, err := om.ToJSON()
	assert.NoError(t, err)

	expectedJSON := `{"1":"one","2":"two","3":"three"}`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}

func TestOrderedMap_Iterate(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	// Iterate and collect the Key-Value pairs
	var keyValues []KeyValue[uint, string]
	for kv := range om.Iterate() {
		keyValues = append(keyValues, kv)
	}

	expectedKeyValues := []KeyValue[uint, string]{
		{Key: 1, Value: "one"},
		{Key: 2, Value: "two"},
		{Key: 3, Value: "three"},
	}

	assert.ElementsMatch(t, expectedKeyValues, keyValues)
}

func TestOrderedMap_Length(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	assert.Equal(t, 0, om.Length())

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	assert.Equal(t, 3, om.Length())

	om.Delete(2)

	assert.Equal(t, 2, om.Length())
}

func TestOrderedMap_Keys(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	keys := om.Keys()
	assert.Empty(t, keys)

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	keys = om.Keys()
	expectedKeys := []uint{1, 2, 3}
	assert.Equal(t, expectedKeys, keys)
}

func TestOrderedMap_Values(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	values := om.Values()
	assert.Empty(t, values)

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	values = om.Values()
	expectedValues := []string{"one", "two", "three"}
	assert.Equal(t, expectedValues, values)
}

func TestOrderedMap_Clear(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	om.Set(1, "one")
	om.Set(2, "two")
	om.Set(3, "three")

	om.Clear()

	assert.Equal(t, 0, om.Length())
}

func TestOrderedMap_Empty(t *testing.T) {
	om := NewOrderedMap[uint, string]()

	assert.True(t, om.Empty())

	om.Set(1, "one")

	assert.False(t, om.Empty())
}
