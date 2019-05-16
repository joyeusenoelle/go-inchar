package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"util"
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

func stats() []int {
	rolls := make([]int, 6)
	for i := 0; i < 6; i++ {
		thisRoll := util.Roll(4, 6)
		for k, v := range thisRoll {
			// fmt.Println(thisRoll)
			for v == 1 {
				// fmt.Println("Replacing 1")
				newv := util.Roll(1, 6)
				thisRoll[k] = newv[0]
				v = newv[0]
			}
		}
		sort.Ints(thisRoll)
		top3 := make([]int, 3)
		copy(top3, thisRoll[1:])
		rolls[i] = util.Sum(top3)
	}
	return rolls
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Keep this the first line in main()

	barb := Class{"Barbarian", 12, []string{"Int", "Wis", "Cha", "Dex", "Con", "Str"}}
	fmt.Println("=== Generated D&D character ===")
	rolls := stats()
	fmt.Println("Rolls: ", rolls)
	sort.Ints(rolls)
	ply := Player{util.Namegen(), barb, make(map[string]int)}
	for k, v := range barb.StatPriority {
		ply.Stats[v] = rolls[k]
	}
	fmt.Printf("%s - %s\n", ply.Name, ply.Class.Name)
	fmt.Printf("Str: %d\tInt: %d\nDex: %d\tWis: %d\nCon: %d\tCha: %d\n",
		ply.Stats["Str"],
		ply.Stats["Int"],
		ply.Stats["Dex"],
		ply.Stats["Wis"],
		ply.Stats["Con"],
		ply.Stats["Cha"])
}
