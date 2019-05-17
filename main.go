package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	"util"
)

type Character struct {
	Name            string
	Kind            string
	Word            string
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

func main() {
	rand.Seed(time.Now().UnixNano()) // Keep this first in main()

	corSkills := []string{"Acrobatics", "Climbing", "Dodge", "Escape", "Fighting", "Large Weapon", "Move Silently", "Running", "Swimming", "Throwing"}
	ethSkills := []string{"Knowledge", "Knowledge", "Knowledge", "Area Knowledge", "Area Knowledge", "Area Knowledge", "Chemistry", "Computer Operation", "Driving", "Electronics", "Engineering", "Language", "Lockpicking", "Lying", "Medicine", "Ranged Weapon", "Savoir-Faire", "Small Weapon", "Tactics"}
	celSkills := []string{"Artistry", "Detect Lies", "Emote", "Fast-Talk", "Seduction", "Singing", "Survival", "Tracking"}
	akList := []string{"Heaven", "Hell", "Marches", "Caribbean", "New York", "New England", "Florida", "Atlanta", "Texas", "California", "American Southwest", "Pacific Northwest", "Portland", "Toronto", "Vancouver", "Mexico", "Central America", "Brazil", "Argentina", "England", "London", "France", "Paris", "Norway", "Scandinavia", "Greece", "Egypt", "North Africa", "Sub-Saharan Africa", "Saudi Arabia", "Middle East", "Russia", "Moscow", "China", "Shanghai", "Hong Kong", "Japan", "Hokkaido", "Tokyo", "Australia", "Sydney", "Melbourne", "Perth", "Fiji", "Antarctica"}
	knowList := []string{"Astronomy", "Biology", "Literature", "Aircraft", "American Football", "Football", "Baseball", "Sumo", "Giant Robot Anime", "German Cuisine", "Catholicism", "Islam", "Buddhism", "Shinto", "Architecture", "Eschatology", "Numinology", "Role-Playing Games", "Spelunking", "Parliamentary Procedure", "Olympic History", "18th-Century Botanical Manuals", "Photography", "Marine Biology", "Entomology", "Archaeology"}
	langList := []string{"Mandarin", "Spanish", "English", "Hindi", "Arabic", "Portuguese", "Bengali", "Russian", "Japanese", "Punjabi", "German", "Javanese", "Wu", "Malay", "Telugu", "Vietnamese", "Korean", "French", "Marathi", "Tamil", "Urdu", "Turkish", "Italian", "Yue (Cantonese)", "Thai", "Latin", "Greek", "Ancient Egyptian", "Apache", "Ainu", "Aleut", "Inuit", "Mayan"}
	songList := []string{"Attraction", "Charm", "Dreams", "Entropy", "Form", "Harmony", "Healing", "Motion", "Numinous Corpus", "Possession", "Projection", "Shields", "Thunder", "Tongues"}

	choirs := []string{"Seraph", "Cherub", "Ofanite", "Elohite", "Malakite", "Kyriotate", "Mercurian"}
	bands := []string{"Balseraph", "Djinn", "Calabite", "Habbalite", "Lilim", "Shedite", "Impudite"}

	aWords := []string{"Dreams", "Children", "Stone", "Judgment", "Creation", "Fire", "the Wind", "Lightning", "Animals", "Faith", "the Sword", "Revelation", "Trade", "War", "Flowers", "Destiny", "Protection"}
	iWords := []string{"Nightmares", "Lust", "the Game", "the War", "Fire", "Drugs", "Hardcore", "Secrets (shhh!)", "Gluttony", "Dark Humor", "Fate", "Freedom", "Factions", "the Media", "Death", "Theft", "Technology"}

	ply := Character{}
	ply.Name = util.Namegen()
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
	ply.Skills = make(map[string]int)
	ply.Songs = make(map[string]int)

	ply.Skills["Language (Local)"] = 3
	if rand.Intn(2) == 0 {
		ply.Kind = util.ChoiceStr(choirs)
		ply.Word = util.ChoiceStr(aWords)
	} else {
		ply.Kind = util.ChoiceStr(bands)
		ply.Word = util.ChoiceStr(iWords)
	}

	ply.Attunements = append(ply.Attunements, ply.Kind+" of "+ply.Word)

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

	cp, corBrk, ethBrk := ply.TotalForces*4, ply.CorporealForces, ply.CorporealForces+ply.EtherealForces
	sklBrk := 12

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

	// Celestial
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

	b := 7 + rand.Intn(6)
	for i := cp; i > b; i-- {
		r := rand.Intn(15)
		if r < sklBrk { // it's a skill
			if rand.Intn(3) == 0 && len(ply.Skills) != 0 {
				key, val := util.ChoiceMap(ply.Skills)
				if ply.Skills[key] == 6 {
					i++
				} else {
					ply.Skills[key] = val + 1
				}
			} else {
				q := rand.Intn(ply.TotalForces)
				skl := ""
				if q < corBrk { // it's a corporeal skill
					skl = util.ChoiceStr(corSkills)
					// fmt.Println("Selected Corporeal skill: ", skl)
				} else if q < ethBrk { // it's an ethereal skill
					skl = util.ChoiceStr(ethSkills)
					if skl == "Knowledge" {
						subskl := util.ChoiceStr(knowList)
						skl = "Knowledge (" + subskl + ")"
					} else if skl == "Area Knowledge" {
						subskl := util.ChoiceStr(akList)
						skl = "Area Knowledge (" + subskl + ")"
					} else if skl == "Language" {
						subskl := util.ChoiceStr(langList)
						skl = "Language (" + subskl + ")"
					}
					// fmt.Println("Selected Ethereal skill: ", skl)
				} else { // it's a celestial skill
					skl = util.ChoiceStr(celSkills)
					// fmt.Println("Selected Celestial skill: ", skl)
				}
				ply.Skills[skl] = ply.Skills[skl] + 1
			}
		} else { // Songs
			if rand.Intn(3) == 0 && len(ply.Songs) != 0 {
				key, val := util.ChoiceMap(ply.Songs)
				if ply.Songs[key] == 6 {
					i++
				} else {
					ply.Songs[key] = val + 1
				}
			} else {
				skl := util.ChoiceStr(songList)
				//			fmt.Println("Selected song: ", skl)
				ply.Songs[skl] = ply.Songs[skl] + 1
			}
		}
	}

	skills := make([]string, 0, len(ply.Skills))
	songs := make([]string, 0, len(ply.Songs))
	// fmt.Println(len(ply.Skills), len(ply.Songs), ply.Skills, ply.Songs)
	for k, v := range ply.Skills {
		skills = append(skills, k+"/"+strconv.Itoa(v))
	}
	for k, v := range ply.Songs {
		songs = append(songs, k+"/"+strconv.Itoa(v))
	}
	sort.Strings(skills)
	sort.Strings(songs)
	sklout := strings.Join(skills[:], ", ")
	sngout := strings.Join(songs[:], ", ")

	fmt.Println("=== Generated In Nomine character ===")
	fmt.Println(ply.Name)
	fmt.Printf("%s of %s\n", ply.Kind, ply.Word)
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
	fmt.Println("Skills:", sklout)
	fmt.Println("Songs:", sngout)
	fmt.Println("Attunements:", strings.Join(ply.Attunements, ", "))
	fmt.Println("CP Remaining:", b)
}
