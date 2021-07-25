package Domain

type TurnAtackResult struct {
	HitAttemptValue          int
	DodgeAttemptValue        int
	DamageDice               int
	CalculatedDamage         int
	AttackerWeaponDiceResult int
	LastAttacker             string
	Message                  string
}
