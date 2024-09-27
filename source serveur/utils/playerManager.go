package hangmanServer

import (
	"net"
)

// Ajoute un joueur à la liste des joueurs, retourne vrai si le joueur a été ajouté
func NewPlayer(connection net.Conn) bool {
	joueur := Player{}
	joueur.Connection = connection
	joueur.Name = ""
	joueur.IsListened = true

	// Met joueur dans ConnectionList
	go ListenToClient(&joueur)
	go AutoResponse(&joueur)

	PrintDebug("Je veut le nom du joueur")
	SendMessageToPlayer(joueur, "GETNAME")

	// Tant qu'on a pas le nom, on attend 10s, au dela on ferme la connexion
	// Tant que le client n'a pas envoyé son nom, il n'est pas considéré comme un joueur
	// Il n'a donc pas d'identifiant
	WaitSecondsWithCondition(30, func() bool { return joueur.Name != "" })

	if joueur.Name == "" {
		PrintDebug("Le joueur n'a pas envoyé son nom a temps, ciao")
		CloseSocket(joueur.Connection)
		return false
	} else {
		PrintDebug("Le joueur a pour nom : ", joueur.Name)
		joueur.Id = *LastPlayerIdPtr
		*LastPlayerIdPtr++
		*PlayerListPtr = append(*PlayerListPtr, joueur)
		return true
	}
}

// Vérifie si un joueur de la liste est déconnecté
func CheckIfDisconnectedPlayers() {
	for _, player := range *PlayerListPtr {
		if player.IsDisconnected {
			// On retire le joueur de la liste
			*PlayerListPtr = RemovePlayerFromList(player.Id)
		}
	}
}

// Retire un joueur de la liste
func RemovePlayerFromList(id int) []Player {
	var newList []Player
	for _, p := range *PlayerListPtr {
		if p.Id != id {
			newList = append(newList, p)
		}
	}
	return newList
}
