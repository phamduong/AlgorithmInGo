package fixture

import (
	"math/rand"
	"sort"
	"time"
)

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandomOneInt() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return RandInt(50, 100)
}

func GenerateArray(n int) []int {
	var result []int
	result = make([]int, n, n)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < n; i++ {
		result[i] = RandInt(1, 9999)
	}
	return result
}

func GetExpectedResult(a []int) []int {
	var length int = len(a)
	var b []int
	b = make([]int, length, length)
	copy(b[:], a)
	sort.Ints(b)
	return b
}
