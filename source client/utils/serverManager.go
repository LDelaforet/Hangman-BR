package hangmanClient

import (
	"net"
)

func CheckServer(connection net.Conn) bool {
	serveur := Server{}
	serveur.Connection = connection
	serveur.LastRequest = ""
	serveur.RequestResponded = false

	go ListenServer(&serveur)
	go AutoResponse(&serveur)

	PrintDebug("Véficication de l'intégrité du serveur")

	PrintDebug("J'envoie un ping")
	SendMessageToServer(serveur, "PING")

	if WaitSecondsWithCondition(5, func() bool { return serveur.RequestResponded }) && serveur.LastRequest == "PONG" {
		PrintDebug("Serveur valide")
		// On stoppe l'écoute temporaire du serveur
		serveur.IsListened = false
		*CurrentServerPtr = serveur
		return true
	} else {
		PrintDebug("Serveur invalide")
		return false
	}
}

func PermanantListen() {
	*&CurrentServerPtr.IsListened = true
	go ListenServer(CurrentServerPtr)
	go AutoResponse(CurrentServerPtr)
}
