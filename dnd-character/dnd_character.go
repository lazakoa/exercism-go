package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10.0) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := []int{}
	for i := 0; i < 4; i++ {
		rolls = append(rolls, rand.Intn(5)+1)
	}

	minRoll := rolls[0]

	for i := 1; i < 4; i++ {
		if rolls[i] < minRoll {
			minRoll = rolls[i]
		}
	}

	total := 0
	for i := 0; i < 4; i++ {
		total += rolls[i]
	}

	return total - minRoll
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	consitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: consitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(consitution),
	}
}
