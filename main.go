package main

import (
	"fmt"
	"math/rand"
	// "sort"
	"strings"
	"time"
)

type Character struct {
	Name            string
	TotalForces     int
	CorporealForces int
	EtherealForces  int
	CelestialForces int
	Strength        int
	Agility         int
	Intellect       int
	Precision       int
	Will            int
	Perception      int
	Skills          map[string]int
	Songs           map[string]int
	Attunements     []string
	Distinctions    []string
	Artifacts       []string
	Vessels         map[string]int
	Roles           map[string][]int
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

func namegen() string {
	vowels := make([]string, 6)
	copy(vowels, strings.Split("aeiouy", ""))
	consonants := make([]string, 20)
	copy(consonants, strings.Split("bcdfghjklmnpqrstvwxz", ""))

	ln := rand.Intn(5) + 5
	vorc := rand.Intn(2)
	i := 0
	nm := make([]string, ln)
	for i < ln {
		sel, c := 0, ""
		if i%2 == vorc {
			sel = rand.Intn(6)
			c = vowels[sel]
		} else {
			sel = rand.Intn(20)
			c = consonants[sel]
		}
		if i == 0 {
			c = strings.ToUpper(c)
		}
		nm = append(nm, c)
		i++
	}
	return strings.Join(nm, "")
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Keep this first in main()

	ply := Character{}
	ply.Name = namegen()
	ply.TotalForces = 9
	ply.CorporealForces = 1
	ply.EtherealForces = 1
	ply.CelestialForces = 1
	ply.Strength = 1
	ply.Agility = 1
	ply.Intellect = 1
	ply.Precision = 1
	ply.Will = 1
	ply.Perception = 1

	i := ply.TotalForces - (ply.CorporealForces + ply.EtherealForces + ply.CelestialForces)
	for i > 0 {
		r := rand.Intn(3)
		if r == 0 {
			ply.CorporealForces = ply.CorporealForces + 1
		} else if r == 1 {
			ply.EtherealForces = ply.EtherealForces + 1
		} else {
			ply.CelestialForces = ply.CelestialForces + 1
		}
		i--
	}

	// Corporeal
	coratt := (ply.CorporealForces * 4) - (ply.Strength + ply.Agility)
	for coratt > 0 {
		r := rand.Intn(2)
		if r == 0 {
			if ply.Strength == 12 || (ply.Strength >= (ply.Agility*2) && rand.Intn(3) == 0) {
				continue
			}
			ply.Strength = ply.Strength + 1
		} else {
			if ply.Agility == 12 || (ply.Agility >= (ply.Strength*2) && rand.Intn(3) == 0) {
				continue
			}
			ply.Agility = ply.Agility + 1
		}
		coratt--
	}

	// Ethereal
	ethatt := (ply.EtherealForces * 4) - (ply.Intellect + ply.Precision)
	for ethatt > 0 {
		r := rand.Intn(2)
		if r == 0 {
			if ply.Intellect == 12 || (ply.Intellect >= (ply.Precision*2) && rand.Intn(3) == 0) {
				continue
			}
			ply.Intellect = ply.Intellect + 1
		} else {
			if ply.Precision == 12 || (ply.Precision >= (ply.Intellect*2) && rand.Intn(3) == 0) {
				continue
			}
			ply.Precision = ply.Precision + 1
		}
		ethatt--
	}

	// Corporeal
	celatt := (ply.CelestialForces * 4) - (ply.Will + ply.Perception)
	for celatt > 0 {
		r := rand.Intn(2)
		if r == 0 {
			if ply.Will == 12 || (ply.Will >= (ply.Perception*2) && rand.Intn(3) == 0) {
				continue
			}
			ply.Will = ply.Will + 1
		} else {
			if ply.Perception == 12 || (ply.Perception >= (ply.Will*2) && rand.Intn(3) == 0) {
				continue
			}
			ply.Perception = ply.Perception + 1
		}
		celatt--
	}

	fmt.Println("=== Generated In Nomine character ===")
	fmt.Println(ply.Name)
	fmt.Printf("Cor %d\tEth %d\tCel %d\n",
		ply.CorporealForces,
		ply.EtherealForces,
		ply.CelestialForces)
	fmt.Printf("Str %d\tInt %d\tCel %d\n",
		ply.Strength,
		ply.Intellect,
		ply.Will)
	fmt.Printf("Agi %d\tPre %d\tPer %d\n",
		ply.Agility,
		ply.Precision,
		ply.Perception)
}
