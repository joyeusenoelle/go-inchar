package util

import (
	"math/rand"
	"strings"
)

func Namegen() string {
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
