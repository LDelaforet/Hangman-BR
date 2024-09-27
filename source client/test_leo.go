package main

import (
	"fmt"
	client "hangmanClient/utils"
)

func main() {
	client.InitASCII()
	client.MainMenu()
}

func testClient() {
	fmt.Print("Nom du joueur: ")
	*client.PlayerNamePtr = client.GetInput(0, "all", 0)

	// Boucle infinie pour choisir un serveur
	for {
		// Bon pour l'instant l'ip et le port sont hardcodés
		conn := client.ConnectSocket("127.0.0.1", "1597")

		// Check le serveur
		if client.CheckServer(conn) {
			serveur := *client.CurrentServerPtr
			fmt.Println("Serveur choisi, nom du serveur: ")
			client.SendMessageToServer(serveur, "GETNAME")
			if !client.WaitSecondsWithCondition(5, func() bool { return serveur.RequestResponded }) || serveur.LastRequest == "GETNAME" {
				fmt.Println("Serveur invalide")
				continue
			}
			fmt.Println(client.CurrentServer.Name)
			client.PermanantListen()
			// Si on est là, c'est que le serveur est valide
			for {
			}
		}
	}
}
