package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func roll(x, y int) []int {
	rolls := make([]int, x)
	for i := 0; i < x; i++ {
		rolls[i] = (rand.Intn(y) + 1)
	}
	return rolls
}

func sum(ary []int) int {
	s := 0
	for _, v := range ary {
		s += v
	}
	return s
}

func stats() []int {
	rolls := make([]int, 6)
	for i := 0; i < 6; i++ {
		thisRoll := roll(4, 6)
		sort.Ints(thisRoll)
		top3 := make([]int, 3)
		copy(top3, thisRoll[1:])
		rolls[i] = sum(top3)
	}
	return rolls
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("=== Generated D&D character ===")
	fmt.Println("Rolls: ", stats())
}
