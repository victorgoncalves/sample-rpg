package EntityFactory

import "github.com/victorgoncalves/sample-rpg/Domain"

func GetDefaultHuman(name string) Domain.Character {
	return Domain.Character{
		Name:     name,
		Type:     Domain.Human,
		Life:     12,
		Strength: 1,
		Agility:  2,
		Weapon: Domain.Weapon{
			Type:    Domain.LongSword,
			Attack:  2,
			Defense: 1,
			Dice: Domain.Dice{
				FacesQuantity: 6,
			},
		},
	}
}

func GetDefaultOrc(name string) Domain.Character {
	return Domain.Character{
		Name:     name,
		Type:     Domain.Orc,
		Life:     20,
		Strength: 2,
		Agility:  0,
		Weapon: Domain.Weapon{
			Type:    Domain.WoodenClub,
			Attack:  1,
			Defense: 0,
			Dice: Domain.Dice{
				FacesQuantity: 8,
			},
		},
	}

}
