package FirstTurnAgilityCalculator

import "github.com/victorgoncalves/sample-rpg/Domain"

func GetAgilityCalculation(character Domain.Character) int {
	dice := Domain.Dice{
		FacesQuantity: 20,
	}

	return character.Agility + dice.RollDice()
}
