package slice

// Function takes an element of T type and returns the output of V type
// Example: 1 => Function[int, string] => "1" takes int converts it to string; here T int and V string
// Code Example: func(i int) string { return strconv.Itoa(i) }
type Function[T, V any] func(item T) V

type Predicate[T comparable] func(item T) bool

type Consumer[T any] func(item T)

type Accumulator[T, V any] func(acc V, item T) V

// Map takes the slice of T type and a mapper Function[T, V] and returns new slice of type V
// Function[T, V]; T is the source slice element type and V determines the type of output slice type.
// Output slice is completely different.
// If the underneath element of source slice is reference value, it may or may not be same (pointing to the same reference/address) depending on the mapper function.
//
// Example:
// 		src := []int{1, 2, 3}
// 		mapper := func(item int) string { return strconv.Itoa(item) }
//  	out := Map(src, mapper)
//  	-----------------------
//  	Output:
// 			[]string{"1", "2", "3"}
// For more examples, see  	-Test_Map_type_conversion
//							-Test_Map_extract_from_struct
func Map[T, V any](arr []T, mapper Function[T, V]) []V {
	var narr []V
	for _, item := range arr {
		narr = append(narr, mapper(item))
	}
	return narr
}

// Filter takes a slice of T type and a Predicate function.
// Filters the element based on Predicate function.
// Output slice's type is as same as the source slice but the output slice is completely different.
// Elements of the source slice and output slice may or may not be the same (pointing to the same reference/address) depending on the elements of source slice.
// Example:
//		src := []int{1, 2, 3, 4}
//		predicate := func(item int) bool { return item % 2 == 0 } // is even?
//		out := Filter(src, predicate) // out contains only the even values
//		----------------------
//		Output:
//			[]int[2, 4]
// For more examples, see	-TestFilter
func Filter[T comparable](arr []T, pred Predicate[T]) []T {
	var narr []T
	for _, item := range arr {
		if pred(item) {
			narr = append(narr, item)
		}
	}
	return narr
}

func ForEach[T any](arr []T, cons Consumer[T]) {
	for _, item := range arr {
		cons(item)
	}
}

// Reduce takes the slice, an initial value and an Accumulator function and returns a combined single value of type V.
// Combines the elements of the source slice into a single value of different type using the Accumulator function where base or initial value is v of type V.
// Example:
//		Calculate total sum
// 		src := []int{1, 2, 3}
//		acc := func(acc int, i int) int { // total sum accumulator
//			s := acc + i
//			return s
//		}
//		sum := Reduce(src, 0, acc)
// 		-------------------------
//		Output:
//			6
func Reduce[T, V any](arr []T, v V, acc Accumulator[T, V]) V {
	for _, item := range arr {
		v = acc(v, item)
	}
	return v
}

// Flat creates a new slice by flattening first nested child slice and returns the new slice.
// Example:
// [][]int{{1, 2}, {4, 5, 6, 7}, {-1}, {-2, -3}} -> Flat() -> []int{1, 2, 4, 5, 6, 7, -1, -2, -3}
// [][][]int{{{1, 2}, {4, 5, 6, 7}}, {{-1}, {-2, -3}}} -> Flat() -> [][]int{{1, 2}, {4, 5, 6, 7}, {-1}, {-2, -3}}
func Flat[T any](arr [][]T) []T {
	acc := func(acc []T, item []T) []T {
		return append(acc, item...)
	}
	return Reduce(arr, []T{}, acc)
}

// FlatMap flats the slice and maps the data using mapper Function.
// It's a chain call of Reduce (flat) and Map.
func FlatMap[T, V any](arr [][]T, mapper Function[T, V]) []V {
	acc := func(acc []T, i []T) []T {
		return append(acc, i...)
	}
	return Map(Reduce(arr, []T{}, acc), mapper)
}

// Find finds the first element based on predicate condition and returns the reference of the element, if not found, returns nil
func Find[T comparable](arr *[]T, pred Predicate[T]) *T {
	for i := 0; i < len(*arr); i++ {
		if pred((*arr)[i]) {
			return &(*arr)[i]
		}
	}
	return nil
}

// FindIndex finds the first element's index based on predicate condition and returns the reference of the element, if not found, returns -1
func FindIndex[T comparable](arr *[]T, pred Predicate[T]) int {
	for i, item := range *arr {
		if pred(item) {
			return i
		}
	}
	return -1
}

// Some takes a reference of slice and a Predicate function
// Some returns true if any element of the source slice satisfies the condition of the Predicate, else returns false
func Some[T comparable](arr *[]T, pred Predicate[T]) bool {
	for _, item := range *arr {
		if pred(item) {
			return true
		}
	}
	return false
}

// Every takes a reference of slice and a Predicate function.
// it returns true if all elements of the source slice satisfy the condition of the Predicate, if any element fails to satisfy, it returns false.
func Every[T comparable](arr *[]T, pred Predicate[T]) bool {
	for _, item := range *arr {
		if !pred(item) {
			return false
		}
	}
	return true
}
