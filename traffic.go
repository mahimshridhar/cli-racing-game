package main

import (
	"math/rand"
)

func NewTraffic() car {
	carIndex := rand.Intn(WIDTH)

	if carIndex%2 == 0 {
		return car{
			position: carIndex,
			shape:    NOCAR,
		}
	}

	return car{
		position: carIndex,
		shape:    CAR,
	}
}
