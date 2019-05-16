package util

import (
	"fmt"
	"math/rand"
	"time"
)

// This doesn't work
// func Choice(ary ...interface{}) interface{} {
// 	rand.Seed(time.Now().UnixNano())
// 	l := len(ary)
// 	return ary[rand.Intn(l)]
// }

func Debug(msg ...interface{}) {
	fmt.Println(msg)
}

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

func ChoiceMap(ary map[string]int) (string, int) {
	rand.Seed(time.Now().UnixNano())
	l := len(ary)
	// fmt.Println(l)
	key, val, rot := "", 0, 0
	for key == "" {
		Debug(rot, key, val)
		for k, v := range ary {
			Debug(k, v)
			// Debug(rot)
			r := 0
			if rot < 3 {
				r = rand.Intn(l)
			}
			if r == 0 {
				key, val = k, v
				break
			}
		}
		rot++
		// fmt.Println(key, val)
	}
	return key, val
}
