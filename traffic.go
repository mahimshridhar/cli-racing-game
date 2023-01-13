package main

import (
	"math/rand"
)

func NewTraffic() car {
	carIndex := rand.Intn(13)

	if carIndex == 5 || carIndex == 3 || carIndex == 10 {
		return car{
			position: carIndex,
			shape:    " ",
		}
	}

	return car{
		position: carIndex,
		shape:    "#",
	}
}
