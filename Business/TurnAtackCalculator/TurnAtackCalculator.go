package TurnAtackCalculator

import (
	"github.com/victorgoncalves/sample-rpg/Domain"
)

// It return the number of damage to subtract defender life,
// if It returns 0 the  defender blocked the atack
// if It return -1 the defender dodged the atack
func GetLifeDamage(atacker *Domain.Character, defender *Domain.Character) Domain.TurnAtackResult {
	dice := Domain.Dice{
		FacesQuantity: 20,
	}
	turnResult := Domain.TurnAtackResult{}

	turnResult.HitAttemptValue = dice.RollDice() + atacker.Agility + atacker.Weapon.Attack
	turnResult.DodgeAttemptValue = dice.RollDice() + defender.Agility + defender.Weapon.Defense

	if turnResult.DodgeAttemptValue < turnResult.HitAttemptValue {
		weaponDiceResult := atacker.Weapon.Dice.RollDice()
		weaponDamage := weaponDiceResult + atacker.Strength
		turnResult.AttackerWeaponDiceResult = weaponDiceResult
		turnResult.CalculatedDamage = weaponDamage
		defender.Life -= weaponDamage
	} else {
		turnResult.Message = "defender.Name" + " dodged the " + atacker.Name + "'s attack! "
	}

	turnResult.LastAttacker = atacker.Name

	return turnResult
}
