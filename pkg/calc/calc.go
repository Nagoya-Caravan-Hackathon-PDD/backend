package calc

import (
	"math/rand"
)

func Attack(attack, defense int) int {
	return int((float64(attack))*(0.9+rand.Float64()*(1.1-0.9)) - (float64(defense) * 0.75))
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
