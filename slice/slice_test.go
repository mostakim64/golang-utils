package slice

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
)

func Test_Map_type_conversion(t *testing.T) {
	src := []int{1, 2, 3}
	t.Run("Convert int to string of each element", func(t *testing.T) {
		exp := []string{"1", "2", "3"}
		f := func(item int) string { return strconv.Itoa(item) }
		act := Map(src, f)

		assert.Equal(t, exp, act)
		assert.NotSame(t, exp, act)
	})

	t.Run("Convert int to float64 of each element", func(t *testing.T) {
		exp := []float64{float64(1), float64(2), float64(3)}
		f := func(item int) float64 { return float64(item) }
		act := Map(src, f)

		assert.Equal(t, exp, act)
		assert.NotSame(t, exp, act)
	})
}

type (
	person struct {
		name string
		age  int
	}
)

func getName(p person) string {
	return p.name
}

func getAge(p person) int {
	return p.age
}

func nameStarWith(p *person, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(p.name), strings.ToLower(prefix))
}

var (
	persons = []person{
		{name: "foo", age: 20},
		{name: "bar", age: 25},
		{name: "john", age: 29},
		{name: "mr. foo", age: 15},
		{name: "mrs. bar"},
		{name: ""},
		{name: "mr.   fo foooo", age: 21},
		{name: "Mr. ba bbbbaaaar", age: 33},
		{name: "Dr. drake"},
		{name: "Dr. ---GODoctor"},
	}

	expectedNames           = []string{"foo", "bar", "john", "mr. foo", "mrs. bar", "", "mr.   fo foooo", "Mr. ba bbbbaaaar", "Dr. drake", "Dr. ---GODoctor"}
	expectedAges            = []int{20, 25, 29, 15, 0, 0, 21, 33, 0, 0}
	personsNameStartsWithMr = []person{
		{name: "mr. foo", age: 15},
		{name: "mr.   fo foooo", age: 21},
		{name: "Mr. ba bbbbaaaar", age: 33},
	}
)

func Test_Map_extract_from_struct(t *testing.T) {
	t.Run("extract field value", func(t *testing.T) {
		names := Map(persons, func(p person) string { return p.name })
		assert.Equal(t, expectedNames, names)

		names1 := Map(persons, getName)
		assert.Equal(t, expectedNames, names1)

		ages := Map(persons, getAge)
		assert.Equal(t, expectedAges, ages)
	})

	t.Run("pointer", func(t *testing.T) {
		type student struct {
			name string
		}
		type dept struct {
			students *[]student
			count    int
		}

		sd1 := []student{
			{"s1"}, {"s2"}, {"s3"},
		}
		sd2 := []student{
			{"s4"}, {"s5"}, {"s6"}, {"s7"},
		}

		d := []dept{
			{students: &sd1},
			{students: &sd2},
		}

		count := func(d dept) int { return len(*d.students) }
		counts := Map(d, count)
		assert.Equal(t, []int{3, 4}, counts)

		f2 := func(d dept) dept {
			d.count = len(*d.students)
			return d
		}
		exp1 := []dept{
			{students: &sd1, count: 3},
			{students: &sd2, count: 4},
		}
		act := Map(d, f2)
		assert.Equal(t, exp1, act)
		assert.Equal(t, d[0].students, act[0].students)
		// element remains the same
		assert.Same(t, d[0].students, act[0].students)
		assert.Exactly(t, d[0].students, act[0].students)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter numerics", func(t *testing.T) {
		src := []int{1, 2, 3, 4}

		isEven := func(item int) bool { return item%2 == 0 }

		evens := Filter(src, isEven)
		assert.Equal(t, []int{2, 4}, evens)
	})

	t.Run("filter name starts with Mr.", func(t *testing.T) {
		// searches starts with prefix
		search := func(prefix string) Predicate[person] {
			return func(p person) bool {
				return nameStarWith(&p, prefix)
			}
		}
		namesStartsWithMr := Filter(persons, search("Mr. "))
		assert.Equal(t, personsNameStartsWithMr, namesStartsWithMr)
	})
}

func TestReduce(t *testing.T) {
	t.Run("cumulative sum", func(t *testing.T) {
		src := []int{1, 2, 3}
		acc := func(acc *int, i *int) int {
			s := *acc + *i
			return s
		}
		sum := Reduce(&src, 0, acc)
		assert.Equal(t, 6, *sum)
	})

	t.Run("find max", func(t *testing.T) {
		src := []int{1, 2, 3}
		acc := func(acc *int, i *int) int {
			if *i > *acc {
				return *i
			}
			return *acc
		}
		max := Reduce(&src, math.MinInt, acc)
		assert.Equal(t, 3, *max)
	})

	t.Run("flat array", func(t *testing.T) {
		src := [][]int{{1, 2}, {4, 5, 6, 7}, {-1}, {-2, -3}}
		ac1 := func(acc *[]int, i *[]int) []int {
			return append(*acc, *i...)
		}
		exp := []int{1, 2, 4, 5, 6, 7, -1, -2, -3}
		assert.Equal(t, &exp, Reduce(&src, []int{}, ac1))
	})

	t.Run("extract field -> cumulative sum", func(t *testing.T) {
		extractAge := func(acc *[]int, p *person) []int {
			return append(*acc, p.age)
		}
		cumSum := func(acc *int, a *int) int {
			s := *acc + *a
			return s
		}
		cumulativeAge := Reduce(Reduce(&persons, []int{}, extractAge), 0, cumSum)
		assert.Equal(t, 143, *cumulativeAge)
	})
}

func TestChaining(t *testing.T) {
	t.Run("filter -> map -> reduce", func(t *testing.T) {
		// searches starts with prefix
		search := func(prefix string) Predicate[person] {
			return func(p person) bool {
				return nameStarWith(&p, prefix)
			}
		}
		findMax := func(acc *int, i *int) int {
			if *i > *acc {
				return *i
			}
			return *acc
		}
		mp := Map(Filter(persons, search("Mr.")), getAge)
		// todo refactor
		maxAge := Reduce(&mp, math.MinInt, findMax)
		assert.Equal(t, 33, *maxAge)
	})
}

func TestFlat(t *testing.T) {
	t.Run("Flat", func(t *testing.T) {
		src1 := [][]int{{1, 2}, {4, 5, 6, 7}, {-1}, {-2, -3}}
		exp1 := []int{1, 2, 4, 5, 6, 7, -1, -2, -3}

		flatted1 := Flat(&src1)
		assert.Equal(t, &exp1, flatted1)

		src2 := [][][]int{{{1, 2}, {4, 5, 6, 7}}, {{-1}, {-2, -3}}}
		exp2 := [][]int{{1, 2}, {4, 5, 6, 7}, {-1}, {-2, -3}}

		flatted2 := Flat(&src2)
		assert.Equal(t, exp2, *flatted2)
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("FlatMap", func(t *testing.T) {
		src := [][]int{{1, 2}, {4, 5, 6, 7}, {-1}, {-2, -3}}
		exp := []int{1, 4, 16, 25, 36, 49, 1, 4, 9}

		square := func(i int) int { return i * i }

		//  todo refactor
		fm := FlatMap(&src, square)
		assert.Equal(t, exp, *fm)
	})
}

func TestFind(t *testing.T) {
	t.Run("Find from numerics", func(t *testing.T) {
		src := []int{1, 4, 16, 25, 36, 49, 1, 4, 9}
		is49 := func(i int) bool { return i == 49 }

		act := Find(&src, is49)
		assert.Equal(t, 49, *act)
	})

	t.Run("Find from structs", func(t *testing.T) {
		greaterThan30 := func(p person) bool { return p.age > 30 }
		act := Find(&persons, greaterThan30)

		assert.Equal(t, persons[7], *act)
		assert.Exactly(t, persons[7], *act)
	})

	t.Run("Find from structs, nil if not found", func(t *testing.T) {
		greaterThan50 := func(p person) bool { return p.age > 50 }
		act := Find(&persons, greaterThan50)
		assert.Nil(t, act)
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("Find index from numerics", func(t *testing.T) {
		src := []int{1, 4, 16, 25, 36, 49, 1, 4, 9}
		is49 := func(i int) bool { return i == 49 }

		assert.Equal(t, 5, FindIndex(&src, is49))
	})

	t.Run("Find index from structs", func(t *testing.T) {
		greaterThan30 := func(p person) bool { return p.age > 30 }
		assert.Equal(t, 7, FindIndex(&persons, greaterThan30))
	})

	t.Run("Find index from structs, -1 if not found", func(t *testing.T) {
		greaterThan50 := func(p person) bool { return p.age > 50 }
		assert.Equal(t, -1, FindIndex(&persons, greaterThan50))
	})
}

func TestSome(t *testing.T) {
	t.Run("find any match", func(t *testing.T) {
		search := func(prefix string) Predicate[person] {
			return func(p person) bool {
				return nameStarWith(&p, prefix)
			}
		}
		assert.True(t, Some(&persons, search("dr. ")))
		assert.False(t, Some(&persons, search("drssss. ")))
	})
}

func TestEvery(t *testing.T) {
	t.Run("find every match", func(t *testing.T) {
		greaterThanAge := func(age int) Predicate[person] {
			return func(p person) bool {
				return p.age > age
			}
		}
		assert.True(t, Every(&persons, greaterThanAge(-1)))
		assert.False(t, Every(&persons, greaterThanAge(100)))

	})

}
