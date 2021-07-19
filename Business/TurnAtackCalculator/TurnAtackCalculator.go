package TurnAtackCalculator

import "github.com/victorgoncalves/sample-rpg/Domain"

// It return the number of damage to subtract defender life,
// if It returns 0 the  defender blocked the atack
// if It return -1 the defender dodged the atack
func GetLifeDamage(atacker Domain.Character, defender Domain.Character, diceValue int) int {
	dice := Domain.Dice{
		FacesQuantity: 20,
	}

	hitValue := dice.RollDice() + atacker.Agility + atacker.Weapon.Attack
	dodgeValue := dice.RollDice() + defender.Agility + defender.Weapon.Defense

	if dodgeValue >= hitValue {
		return -1
	} else {
		return getAtackDamage(atacker)
	}
}

func getAtackDamage(atacker Domain.Character) int {
	return atacker.Weapon.Dice.RollDice() + atacker.Strength
}
