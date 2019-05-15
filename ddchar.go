package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Class struct {
	Name         string
	HitDie       int
	StatPriority []string
}

type Player struct {
	Name  string
	Class Class
	Stats map[string]int
}

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
		for k, v := range thisRoll {
			// fmt.Println(thisRoll)
			for v == 1 {
				// fmt.Println("Replacing 1")
				newv := roll(1, 6)
				thisRoll[k] = newv[0]
				v = newv[0]
			}
		}
		sort.Ints(thisRoll)
		top3 := make([]int, 3)
		copy(top3, thisRoll[1:])
		rolls[i] = sum(top3)
	}
	return rolls
}

func main() {
	barb := Class{"Barbarian", 12, []string{"Int", "Wis", "Cha", "Dex", "Con", "Str"}}
	rand.Seed(time.Now().UnixNano())
	fmt.Println("=== Generated D&D character ===")
	rolls := stats()
	fmt.Println("Rolls: ", rolls)
	sort.Ints(rolls)
	ply := Player{"Bob", barb, make(map[string]int)}
	for k, v := range barb.StatPriority {
		ply.Stats[v] = rolls[k]
	}
	fmt.Println(ply)
}
