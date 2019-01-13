package utils

import (
	"fmt"
	"math/rand"
)

var count = 0

var adjectives = []string {
	"Mean",
	"Scheming",
	"Rude",
	"Agressive",
	"Unpleasant",
	"Dirty",
	"Mean",
	"Ruthless",
	"Stinky",
	"Bitter",
	"Spiteful",
	"Sloppy",
	"Angry",
	"Jealous",
	"Itchy",
	"Obstreperous",
	"Obnoxious",
	"Scurvy",
	"Nasty",
	"Toothy",
}

var names = []string{
	"Bob",
	"Dave",
	"Stan",
	"Harry",
	"Jake",
	"Sanchez",
	"Rasputan",
	"Caligula",
	"Jim",
	"Jack",
	"Steve",
	"Pete",
	"Dan",
	"Joe",
	"Nero",
	"Dante",
	"Gilbert",
	"Zak",
	"Bill",
	"Fred",
}

func GenerateName() string {
	count++
	a := adjectives[rand.Intn(len(adjectives))]
	n := names[rand.Intn(len(names))]

	return fmt.Sprintf("%s-%s-%d", a, n, count)
}