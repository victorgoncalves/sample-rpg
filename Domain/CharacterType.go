package Domain

import "fmt"

type CharacterType int

const (
	Human CharacterType = iota
	Orc
)

func (characterType CharacterType) String() string {
	switch characterType {
	case Human:
		return "Human"
	case Orc:
		return "Orc"
	default:
		return fmt.Sprintf("%d", characterType)
	}
}
