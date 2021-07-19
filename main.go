package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/victorgoncalves/sample-rpg/Business/EntityFactory"
)

func main() {
	sessionId := uuid.New()
	sessionDateTimeInit := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("SessionId " + sessionId.String() + " started at " + sessionDateTimeInit + "\n\n")

	playerHuman := EntityFactory.GetDefaultHuman("Nome do Player")
	npcOrc := EntityFactory.GetDefaultOrc("Groob")

	fmt.Printf("%+v \n\n", playerHuman)
	fmt.Printf("%+v \n", npcOrc)

	fmt.Println("\n\n - - END - -")
}
