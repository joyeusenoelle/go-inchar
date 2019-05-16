package util

import (
	// "fmt"
	"math/rand"
	"time"
)

// This doesn't work
// func Choice(ary ...interface{}) interface{} {
// 	rand.Seed(time.Now().UnixNano())
// 	l := len(ary)
// 	return ary[rand.Intn(l)]
// }

func ChoiceInt(ary []int) int {
	rand.Seed(time.Now().UnixNano())
	l := len(ary)
	return ary[rand.Intn(l)]
}

func ChoiceStr(ary []string) string {
	rand.Seed(time.Now().UnixNano())
	l := len(ary)
	r := rand.Intn(l)
	// fmt.Printf("Selected %s (element %d)\n", ary[r], r)
	return ary[r]
}
