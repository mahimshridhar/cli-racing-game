package main

import (
	"math/rand"
	"strings"
)

func NewTraffic() []string {
	carIndex := rand.Intn(13)
	if carIndex%2 == 0 && carIndex-1 > 0 {
		carIndex = carIndex - 1
	}
	traffic := strings.Split("|"+strings.Repeat(" |", 6), "")
	traffic[carIndex] = "#"

	return traffic
}
