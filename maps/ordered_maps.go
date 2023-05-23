package maps

import (
	"container/list"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

type (
	KeyValue[K comparable, V any] struct {
		Key   K
		Value V
	}

	OrderedMap[K comparable, V any] struct {
		keys   *list.List
		values map[K]*list.Element
	}
)

// NewOrderedMap takes the key of K type and Value of type V and returns new OrderedMap of type K, V
// K is the map key element type and V determines the type of value a key will have.
// Output Map will preserve the order of keys as they inserted/set.
//
// Example:
//
//		om := NewOrderedMap[uint, string]()
//		om.Set(2, "two")
//	 	om.Set(1, "one")
//		om.Set(3, "three")
//
//		om.Print()
//	 	-----------------------
//		Output:
//			Key: 2, Value: two
//			Key: 1, Value: one
//			Key: 3, Value: three
//
// For more examples, see  	-ordered_maps_test
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   list.New(),
		values: make(map[K]*list.Element),
	}
}

// Set will insert a key and value of type K, V to the Ordered Map
func (om *OrderedMap[K, V]) Set(key K, value V) {
	if elem, ok := om.values[key]; ok {
		elem.Value.(*KeyValue[K, V]).Value = value
		return
	}

	kv := &KeyValue[K, V]{Key: key, Value: value}
	elem := om.keys.PushBack(kv)
	om.values[key] = elem
}

// BulkSet will assign a map[K]V of type K, V to the Ordered Map
func (om *OrderedMap[K, V]) BulkSet(data map[K]V) {
	for key, value := range data {
		om.Set(key, value)
	}
}

// Get will take a key of type K and returns its value of type V and exists boolean flag from the Ordered Map
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	if elem, ok := om.values[key]; ok {
		return elem.Value.(*KeyValue[K, V]).Value, true
	}
	var defaultValue V
	return defaultValue, false
}

// Delete will take a key of type K and find it in the Ordered Map and remove its key and value from the Ordered Map
func (om *OrderedMap[K, V]) Delete(key K) {
	if elem, ok := om.values[key]; ok {
		om.keys.Remove(elem)
		delete(om.values, key)
	}
}

// ReorderKeys will take a boolean flag called ascending and order the map key value pair
// as ascending or descending order
func (om *OrderedMap[K, V]) ReorderKeys(ascending bool) {
	keys := make([]K, 0, om.keys.Len())
	for key := range om.values {
		keys = append(keys, key)
	}

	if ascending {
		sort.Slice(keys, func(i, j int) bool {
			return om.less(keys[i], keys[j])
		})
	} else {
		sort.Slice(keys, func(i, j int) bool {
			return om.less(keys[j], keys[i])
		})
	}

	newKeys := list.New()
	newValues := make(map[K]*list.Element)

	for _, key := range keys {
		value := om.values[key].Value
		newElement := newValues[key]
		if newElement == nil {
			newElement = newKeys.PushBack(value)
		} else {
			newElement.Value = value
		}
		newValues[key] = newElement
	}

	om.keys = newKeys
	om.values = newValues
}

// less will be used to reorder the map keys of type K
//
// currently supported 3 key types are uint, int & string
func (om *OrderedMap[K, V]) less(a, b K) bool {
	reflectA := fmt.Sprint(a)
	reflectB := fmt.Sprint(b)

	switch reflect.TypeOf(a).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectA < reflectB
	case reflect.String:
		return om.stringLess(reflectA, reflectB)
	default:
		panic("Unsupported key type")
	}
}

// stringLess will be used for keys
func (om *OrderedMap[K, V]) stringLess(a, b string) bool {
	return a < b
}

// Keys will return a new slice of keys of type K
func (om *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, om.keys.Len())
	for elem := om.keys.Front(); elem != nil; elem = elem.Next() {
		keyValue := elem.Value.(*KeyValue[K, V])
		key := reflect.ValueOf(keyValue.Key).Interface().(K)
		keys = append(keys, key)
	}
	return keys
}

// Values will return a new slice of values of type V
func (om *OrderedMap[K, V]) Values() []V {
	values := make([]V, 0, om.keys.Len())
	for elem := om.keys.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(*KeyValue[K, V]).Value)
	}
	return values
}

// Print will log the key value of ordered map in the console
func (om *OrderedMap[K, V]) Print() {
	for elem := om.keys.Front(); elem != nil; elem = elem.Next() {
		keyValue := elem.Value.(*KeyValue[K, V])
		fmt.Printf("Key: %v, Value: %v\n", keyValue.Key, keyValue.Value)
	}
}

// Iterate will go through the ordered map and return a channel of KeyValue pair of type K, V
// until OrderedMap if nil or empty
func (om *OrderedMap[K, V]) Iterate() <-chan KeyValue[K, V] {
	ch := make(chan KeyValue[K, V])

	go func() {
		defer close(ch)

		for elem := om.keys.Front(); elem != nil; elem = elem.Next() {
			keyValue := elem.Value.(*KeyValue[K, V])
			ch <- KeyValue[K, V]{keyValue.Key, keyValue.Value}
		}
	}()

	return ch
}

// ToJSON will transform the OrderedMap into JSON data
func (om *OrderedMap[K, V]) ToJSON() ([]byte, error) {
	// Create a map to hold the Key-Value pairs
	jsonMap := make(map[K]V)

	// Populate the map with the Key-Value pairs from the OrderedMap
	for elem := om.keys.Front(); elem != nil; elem = elem.Next() {
		keyValue := elem.Value.(*KeyValue[K, V])
		jsonMap[keyValue.Key] = keyValue.Value
	}

	// Marshal the map into JSON
	return json.Marshal(jsonMap)
}

// Length will return the length of the OrderedMap
func (om *OrderedMap[K, V]) Length() int {
	return om.keys.Len()
}

// Clear will reinitialize the OrderedMap and clear any previous key value data it has.
func (om *OrderedMap[K, V]) Clear() {
	om.keys = list.New()
	om.values = make(map[K]*list.Element)
}

// Empty will check if OrderedMap is empty or not
func (om *OrderedMap[K, V]) Empty() bool {
	return om.Length() == 0
}
