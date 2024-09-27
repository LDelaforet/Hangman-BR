package main

import (
	"fmt"
	server "hangmanServer/utils"
)

func main() {
	serverMain()
}

func serverMain() {
	fmt.Println("Lancement du serveur...")
	// Lance le serveur
	*server.CurrentListenerPtr = server.CreateServer("1597")
	*server.IsListeningPtr = true
	go server.AcceptConnection(*server.CurrentListenerPtr)
	for {
		// Si pas assez de joueur, on attend
		if len(*server.PlayerListPtr) <= 2 {
			continue
		} else {
			fmt.Println("Liste des joueurs:")
			for _, joueur := range *server.PlayerListPtr {
				fmt.Println(joueur.Name)
			}
			break
		}
	}
}
