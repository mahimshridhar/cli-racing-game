package main

import (
	"math/rand"
)

func NewTraffic() car {
	carIndex := rand.Intn(13)
	if carIndex%2 == 0 && carIndex-1 > 0 {
		carIndex = carIndex - 1
	}

	return car{
		position: carIndex,
	}
}
