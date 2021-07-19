package Domain

import "fmt"

type WeaponType int

const (
	LongSword WeaponType = iota
	WoodenClub
)

func (weaponType WeaponType) String() string {
	switch weaponType {
	case LongSword:
		return "Long Sword"
	case WoodenClub:
		return "Wooden Club"
	default:
		return fmt.Sprintf("%d", weaponType)
	}
}
