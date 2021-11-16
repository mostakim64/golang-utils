package methods

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

func MapToStruct(input map[string]interface{}, output interface{}) error {
	if b, err := json.Marshal(input); err == nil {
		return json.Unmarshal(b, &output)
	} else {
		return err
	}
}

func InArray(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func GenerateRandomStringOfLength(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	if length == 0 {
		length = 8
	}

	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func StringToIntArray(stringArray []string) []int {
	var res []int

	for _, v := range stringArray {
		if i, err := strconv.Atoi(v); err == nil {
			res = append(res, i)
		}
	}

	return res
}

func RecoverPanic() {
	if r := recover(); r != nil {
		log.Error(r)
	}
}

func IsEmpty(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func MaxOf(vars ...int64) int64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsFloat64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Round - will return a float with 2 decimal point
func Round(x float64) float64 {
	return math.Round(x*100) / 100
}

func PrettyPrint(msg string, data interface{}) {
	if r, err := json.MarshalIndent(&data, "", "  "); err == nil {
		fmt.Printf("[INFO] %v %v: \n %v\n", time.Now(), msg, string(r))
	}
}

func Contains(s []uint, item uint) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func ConvertToUintSlice(value string) []uint {
	var ids []uint
	if value == "" {
		return ids
	}
	splitted := strings.Split(value, ",")
	for _, v := range splitted {
		id, _ := strconv.Atoi(v)
		ids = append(ids, uint(id))
	}

	return ids
}

func ConvertToIntSlice(value string) []int {
	var ids []int
	if value == "" {
		return ids
	}
	splitted := strings.Split(value, ",")
	for _, v := range splitted {
		id, _ := strconv.Atoi(v)
		ids = append(ids, id)
	}

	return ids
}

func StructToStruct(input interface{}, output interface{}) error {
	if b, err := json.Marshal(input); err == nil {
		return json.Unmarshal(b, &output)
	} else {
		return err
	}
}

// remove suffix from the string if string ends with the suffix
func TrimSuffix(s, suffix string) string {
	if ok := strings.HasSuffix(s, suffix); ok {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// Difference returns the elements of `a` that are not in `b`.
func Difference(a, b []int) []int {
	m := make(map[int]bool, len(b))

	for _, v := range b {
		m[v] = true
	}

	var diff []int

	for _, v := range a {
		if _, found := m[v]; !found {
			diff = append(diff, v)
		}
	}

	return diff
}

func Unique(arr []int) []int {
	valueMap := make(map[int]struct{})
	unique := make([]int, 0)

	for _, v := range arr {
		if _, ok := valueMap[v]; !ok {
			valueMap[v] = struct{}{}
			unique = append(unique, v)
		}
	}

	return unique
}
