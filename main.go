package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/victorgoncalves/sample-rpg/Business/EntityFactory"
	"github.com/victorgoncalves/sample-rpg/Business/FirstTurnAgilityCalculator"
	"github.com/victorgoncalves/sample-rpg/Business/TurnAtackCalculator"
	"github.com/victorgoncalves/sample-rpg/Domain"
)

func main() {
	sessionId := uuid.New()
	sessionDateTimeInit := time.Now().Format("2006-01-02 15:04:05")
	lastAction := Domain.TurnAtackResult{
		LastAttacker: "<<not started>>",
	}
	fmt.Println("SessionId " + sessionId.String() + " started at " + sessionDateTimeInit + "\n\n")
	fmt.Print("Enter the character name:")
	var playerName string
	fmt.Scanln(&playerName)

	playerHuman := EntityFactory.GetDefaultHuman(playerName)
	npcOrc := EntityFactory.GetDefaultOrc("Grunck")

	round := 1
	if isHumanFirstVerify(playerHuman, npcOrc) {
		fmt.Println("\n\n your turn, push enter to roll dice and try hit the orc! ")
		fmt.Scanln()
		lastAction = TurnAtackCalculator.GetLifeDamage(&playerHuman, &npcOrc)
		printCurrentStatus(playerHuman, npcOrc, lastAction, round)
		round++
	}

	printCurrentStatus(playerHuman, npcOrc, lastAction, round)

	for playerHuman.Life > 0 && npcOrc.Life > 0 {
		round++
		fmt.Println("Now is the Orc Turn!")
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("The orc will try attack you now!!")
		time.Sleep(2000 * time.Millisecond)
		lastAction = TurnAtackCalculator.GetLifeDamage(&npcOrc, &playerHuman)
		printCurrentStatus(playerHuman, npcOrc, lastAction, round)

		if playerHuman.Life > 0 {
			fmt.Println("\n\n your turn, push enter to roll dice and try hit the orc! ")
			fmt.Scanln()
			round++
			lastAction = TurnAtackCalculator.GetLifeDamage(&playerHuman, &npcOrc)
			printCurrentStatus(playerHuman, npcOrc, lastAction, round)
		}
	}

	printWinnerMessage(playerHuman, npcOrc)

	fmt.Println("\n\n - - - - - END | press the Enter key to exit - - - - - ")
	fmt.Scanln()
}

func isHumanFirstVerify(playerHuman Domain.Character, npcOrc Domain.Character) bool {
	humanAgility := 0
	orcAgility := 0

	for humanAgility == orcAgility {
		fmt.Println("\n\n Push enter to run based on your agility and try get the right of first attack !!")
		fmt.Scanln()
		humanAgility = FirstTurnAgilityCalculator.GetAgilityCalculation(playerHuman)
		orcAgility = FirstTurnAgilityCalculator.GetAgilityCalculation(npcOrc)

		fmt.Println("The Human calculated agility:", humanAgility, "The Orc calculated agility:", orcAgility)

		if humanAgility == orcAgility {
			fmt.Println("You tied, try again!")
		}
	}

	isHumanFirst := humanAgility > orcAgility

	if isHumanFirst {
		fmt.Println("You was more fast and will play first !")
	} else {
		fmt.Println("The Orc was more fast and will play first !")
	}

	return isHumanFirst
}

func printWinnerMessage(human Domain.Character, orc Domain.Character) {
	fmt.Println("\n\n-----------------------FINISHED-----------------------")
	if human.Life <= 0 {
		fmt.Println("The Orc WINS !!! Try Again !")
	}
	if orc.Life <= 0 {
		fmt.Println("Congratulations you're the WINNER !!!")
	}
}

func printCurrentStatus(human Domain.Character, orc Domain.Character, lastAction Domain.TurnAtackResult, round int) {
	roundText := ""
	if round < 10 {
		roundText = fmt.Sprintf("0%d", round)
	} else {
		roundText = fmt.Sprintf("%d", round)
	}
	fmt.Println("\n\n-----------------------Round " + roundText + "-----------------------")
	fmt.Println("\t | Human |")
	fmt.Println("Name:", human.Name)
	fmt.Println("Life:", human.Life)
	fmt.Println("------------------------------------------------------")
	fmt.Println("\t | Orc   |")
	fmt.Println("Name:", orc.Name)
	fmt.Println("Life:", orc.Life)
	fmt.Println("------------------------------------------------------")
	fmt.Println(lastAction.Message)
	fmt.Println("LastDice: ", lastAction.DamageDice, " Damage:", lastAction.CalculatedDamage, "Last Attacker:", lastAction.LastAttacker)
	fmt.Println("------------------------------------------------------")
}
