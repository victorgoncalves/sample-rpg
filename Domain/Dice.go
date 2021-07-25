package Domain

import (
	"math/rand"
	"time"
)

type Dice struct {
	FacesQuantity int
}

func (dice Dice) RollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(dice.FacesQuantity) + 1
}
