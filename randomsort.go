package orindex

import (
	"math/rand"
)

var num int

func checkDuplication(index []int, val int) bool {
	for _, d := range index {
		if d == val {
			return true
		}
	}
	return false
}

func makeIndex(seed int64) []int {
	rand.Seed(seed)
	var index []int
	tmp := rand.Int() % num
	for i := 0; i < num; tmp = rand.Int() % num {
		if !checkDuplication(index, tmp) {
			index = append(index, tmp)
			i++
		}
	}
	return index
}

func s2seed(s string) (sum int64) {
	for i, c := range s {
		sum += int64(i * int(c))
	}
	return sum
}

func Randindex(name string, next_idx int, max int) int {
	num = max
	index := makeIndex(s2seed(name))
	return index[next_idx]
}
