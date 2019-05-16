package util

import (
	"math/rand"
)

func Roll(x, y int) []int {
	rolls := make([]int, x)
	for i := 0; i < x; i++ {
		rolls[i] = (rand.Intn(y) + 1)
	}
	return rolls
}

func Sum(ary []int) int {
	s := 0
	for _, v := range ary {
		s += v
	}
	return s
}
