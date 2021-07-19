package Domain

import "math/rand"

type Dice struct {
	FacesQuantity int
}

func (dice Dice) RollDice() int {
	return rand.Intn(dice.FacesQuantity) + 1
}
