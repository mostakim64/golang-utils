package maps

import (
	"container/list"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

type OrderedMap[K comparable, V any] struct {
	keys   map[K]*list.Element
	values *list.List
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make(map[K]*list.Element),
		values: list.New(),
	}
}

func (om *OrderedMap[K, V]) Set(key K, value V) {
	if elem, ok := om.keys[key]; ok {
		elem.Value.(*KeyValue[K, V]).Value = value
		return
	}

	kv := &KeyValue[K, V]{Key: key, Value: value}
	elem := om.values.PushBack(kv)
	om.keys[key] = elem
}

func (om *OrderedMap[K, V]) BulkSet(data map[K]V) {
	for key, value := range data {
		om.Set(key, value)
	}

	// default ordering: ASC
	om.ReorderKeys(true)
}

func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	if elem, ok := om.keys[key]; ok {
		return elem.Value.(*KeyValue[K, V]).Value, true
	}
	var defaultValue V
	return defaultValue, false
}

func (om *OrderedMap[K, V]) Delete(key K) {
	if elem, ok := om.keys[key]; ok {
		om.values.Remove(elem)
		delete(om.keys, key)
	}
}

func (om *OrderedMap[K, V]) ReorderKeys(ascending bool) {
	keys := make([]K, 0, len(om.keys))
	for key := range om.keys {
		keys = append(keys, key)
	}

	if ascending {
		sort.Slice(keys, func(i, j int) bool {
			return om.Less(keys[i], keys[j])
		})
	} else {
		sort.Slice(keys, func(i, j int) bool {
			return om.Less(keys[j], keys[i])
		})
	}

	newKeys := make(map[K]*list.Element)
	newValues := list.New()

	for _, key := range keys {
		value := om.keys[key].Value
		newElement := newValues.PushBack(value)
		newKeys[key] = newElement
	}

	om.keys = newKeys
	om.values = newValues
}

func (om *OrderedMap[K, V]) Less(a, b K) bool {
	reflectA := reflect.ValueOf(a)
	reflectB := reflect.ValueOf(b)

	switch reflectA.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflectA.Uint() < reflectB.Uint()
	case reflect.String:
		return reflectA.String() < reflectB.String()
	default:
		panic("Unsupported key type")
	}
}

func (om *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(om.keys))
	for key := range om.keys {
		keys = append(keys, key)
	}
	return keys
}

func (om *OrderedMap[K, V]) Values() []V {
	values := make([]V, 0, om.values.Len())
	for elem := om.values.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(*KeyValue[K, V]).Value)
	}
	return values
}

func (om *OrderedMap[K, V]) Print() {
	for elem := om.values.Front(); elem != nil; elem = elem.Next() {
		keyValue := elem.Value.(*KeyValue[K, V])
		fmt.Printf("Key: %v, Value: %v\n", keyValue.Key, keyValue.Value)
	}
}

func (om *OrderedMap[K, V]) Iterate() <-chan KeyValue[K, V] {
	ch := make(chan KeyValue[K, V])

	go func() {
		defer close(ch)

		for elem := om.values.Front(); elem != nil; elem = elem.Next() {
			keyValue := elem.Value.(*KeyValue[K, V])
			ch <- KeyValue[K, V]{keyValue.Key, keyValue.Value}
		}
	}()

	return ch
}

func (om *OrderedMap[K, V]) ToJSON() ([]byte, error) {
	// Create a map to hold the Key-Value pairs
	jsonMap := make(map[K]V)

	// Populate the map with the Key-Value pairs from the OrderedMap
	for elem := om.values.Front(); elem != nil; elem = elem.Next() {
		keyValue := elem.Value.(*KeyValue[K, V])
		jsonMap[keyValue.Key] = keyValue.Value
	}

	// Marshal the map into JSON
	return json.Marshal(jsonMap)
}

func (om *OrderedMap[K, V]) Length() int {
	return len(om.keys)
}

func (om *OrderedMap[K, V]) Clear() {
	om.keys = make(map[K]*list.Element)
	om.values = list.New()
}

func (om *OrderedMap[K, V]) Empty() bool {
	return om.Length() == 0
}
