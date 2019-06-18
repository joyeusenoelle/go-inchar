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

type Race struct {
	Name     string
	Str      int
	Dex      int
	Con      int
	Int      int
	Wis      int
	Cha      int
	Features []string
}

type Player struct {
	Name  string
	Class *Class
	Race  *Race
	Stats map[string]int
}

func choiceClass(ary []Class) Class {
	r := rand.Intn(len(ary))
	return ary[r]
}

func choiceRace(ary []Race) Race {
	r := rand.Intn(len(ary))
	return ary[r]
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

func pickClass() *Class {
	barbarian := Class{"Barbarian", 12, []string{"Int", "Wis", "Cha", "Dex", "Con", "Str"}}
	bard := Class{"Bard", 8, []string{"Str", "Con", "Wis", "Int", "Dex", "Cha"}}
	cleric := Class{"Cleric", 8, []string{"Dex", "Int", "Cha", "Str", "Con", "Wis"}}
	druid := Class{"Druid", 8, []string{"Cha", "Str", "Dex", "Int", "Con", "Wis"}}
	fighter_arch := Class{"Fighter", 10, []string{"Int", "Cha", "Wis", "Str", "Con", "Dex"}}
	fighter_melee := Class{"Fighter", 10, []string{"Int", "Cha", "Wis", "Dex", "Con", "Str"}}
	monk := Class{"Monk", 8, []string{"Cha", "Int", "Str", "Con", "Wis", "Dex"}}
	paladin := Class{"Paladin", 10, []string{"Dex", "Int", "Wis", "Con", "Cha", "Str"}}
	ranger := Class{"Ranger", 10, []string{"Cha", "Int", "Con", "Str", "Wis", "Dex"}}
	rogue := Class{"Rogue", 8, []string{"Wis", "Str", "Con", "Int", "Cha", "Dex"}}
	sorcerer := Class{"Sorcerer", 6, []string{"Str", "Dex", "Wis", "Int", "Con", "Cha"}}
	warlock := Class{"Warlock", 8, []string{"Str", "Dex", "Wis", "Int", "Con", "Cha"}}
	wizard := Class{"Wizard", 6, []string{"Str", "Dex", "Cha", "Con", "Wis", "Int"}}

	classlist := []Class{barbarian, bard, cleric, druid, fighter_arch, fighter_melee, monk, paladin, ranger, rogue, sorcerer, warlock, wizard}
	rclass := choiceClass(classlist)
	return &rclass
}

func (hmn *Race) humanAttr() {
	for i := 0; i < 2; i++ {
		switch r := rand.Intn(6); r {
		case 0:
			hmn.Str++
		case 1:
			hmn.Dex++
		case 2:
			hmn.Con++
		case 3:
			hmn.Int++
		case 4:
			hmn.Wis++
		case 5:
			hmn.Cha++
		}
	}
}

func pickRace() *Race {
	dragonborn := Race{"Dragonborn", 2, 0, 0, 0, 0, 1, []string{"Draconic Ancestry", "Breath Weapon", "Damage Resistance"}}
	dwarf := Race{"Dwarf", 0, 0, 2, 0, 0, 0, []string{"Darkvision", "Dwarven Resilience", "Dwarven Combat Training", "Stonecunning"}}
	elf := Race{"Elf", 0, 2, 0, 0, 0, 0, []string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"}}
	gnome := Race{"Gnome", 0, 0, 0, 2, 0, 0, []string{"Darkvision", "Gnome Cunning"}}
	halfelf := Race{"Half-Elf", 0, 0, 0, 0, 0, 2, []string{"Darkvision", "Fey Ancestry", "Skill Versatility"}}
	halfling := Race{"Halfling", 0, 2, 0, 0, 0, 0, []string{"Lucky", "Brave", "Halfling Nimbleness"}}
	halforc := Race{"Half-Orc", 2, 0, 1, 0, 0, 0, []string{"Darkvision", "Menacing", "Relentless Endurance", "Savage Attacks"}}
	human := Race{"Human", 1, 1, 1, 1, 1, 1, []string{"Extra Language"}}
	tiefling := Race{"Tiefling", 0, 0, 0, 1, 0, 2, []string{"Darkvision", "Hellish Resistance", "Infernal Legacy"}}

	racelist := []Race{dragonborn, dwarf, elf, gnome, halfelf, halfling, halforc, human, tiefling}
	rrace := choiceRace(racelist)
	if rrace.Name == "Half-Elf" {
		rrace.humanAttr()
	}
	return &rrace
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Keep this the first line in main()

	fmt.Println("=== Generated D&D character ===")
	rolls := stats()
	fmt.Println("Rolls: ", rolls)
	sort.Ints(rolls)
	ply := Player{util.Namegen(), pickClass(), pickRace(), make(map[string]int)}
	for k, v := range ply.Class.StatPriority {
		ply.Stats[v] = rolls[k]
	}
	ply.Stats["Str"] += ply.Race.Str
	ply.Stats["Dex"] += ply.Race.Dex
	ply.Stats["Con"] += ply.Race.Con
	ply.Stats["Int"] += ply.Race.Int
	ply.Stats["Wis"] += ply.Race.Wis
	ply.Stats["Cha"] += ply.Race.Cha
	fmt.Printf("%s - %s %s\n", ply.Name, ply.Race.Name, ply.Class.Name)
	fmt.Printf("Str: %d\tInt: %d\nDex: %d\tWis: %d\nCon: %d\tCha: %d\n",
		ply.Stats["Str"],
		ply.Stats["Int"],
		ply.Stats["Dex"],
		ply.Stats["Wis"],
		ply.Stats["Con"],
		ply.Stats["Cha"])
}
