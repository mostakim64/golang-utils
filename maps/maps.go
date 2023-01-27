package maps

// Keys returns the slice[T comparable] of keys of the corresponding map
func Keys[T comparable, V any](mp map[T]V) []T {
	keys := make([]T, 0, len(mp))
	for key := range mp {
		keys = append(keys, key)
	}
	return keys
}

// Values returns the slice[V any] of values of the corresponding map
func Values[T comparable, V any](mp map[T]V) []V {
	values := make([]V, 0, len(mp))
	for key := range mp {
		values = append(values, mp[key])
	}
	return values
}

// MapEntry defines the map's key value definition
type MapEntry[T comparable, V any] struct {
	Key   T
	Value V
}

// Entries returns the slice of MapEntry[T, V] of the corresponding map
func Entries[T comparable, V any](mp map[T]V) []MapEntry[T, V] {
	entries := make([]MapEntry[T, V], 0, len(mp))
	for key := range mp {
		entries = append(entries, MapEntry[T, V]{Key: key, Value: mp[key]})
	}
	return entries
}
