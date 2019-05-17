package util

import (
	"fmt"
	"math/rand"
	"time"
)

func Debug(msg ...interface{}) {
	fmt.Println(msg)
}

func ChoiceInt(ary []int) int {
	l := len(ary)
	return ary[rand.Intn(l)]
}

func ChoiceStr(ary []string) string {
	l := len(ary)
	r := rand.Intn(l)
	// fmt.Printf("Selected %s (element %d)\n", ary[r], r)
	return ary[r]
}

func ChoiceMap(ary map[string]int) (string, int) {
	l := len(ary)
	keys := make([]string, 0, l)
	for k := range ary {
		keys = append(keys, k)
	}
	sel := ChoiceStr(keys)
	return sel, ary[sel]
}
