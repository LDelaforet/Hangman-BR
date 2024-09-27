package hangmanServer

import (
	"net"
)

type Player struct {
	Id             int      // Identifiant unique
	Name           string   // Nom du joueur
	Score          int      // Score actuel
	Word           []rune   // Mot actuel a trouver
	TriedLetters   []rune   // Lettres déjà essayées
	RemainingLives int      // Nombre de vies restantes
	IsListened     bool     // Défini si le joueur est sur écoute (WHEN YOU WALK TO THE GARDEN, YOU BETTER WATCH YOUR BACK)
	ReadyToStart   bool     // S'active si le joueur a voté prêt
	ReceivedBuffer []byte   // Buffer de réception
	IsDisconnected bool     // Défini si le joueur est déconnecté
	Connection     net.Conn // Connexion du joueur
}

// Détermine si le mode debug est activé
var DebugMode bool = true
var DebugModePtr = &DebugMode

// Nom du serveur
var ServerName string = "TestServ"

// Dernier identifiant attribué a un joueur
var LastPlayerId int
var LastPlayerIdPtr *int = &LastPlayerId

// Liste les joueurs
var PlayerList []Player
var PlayerListPtr *[]Player = &PlayerList

// Données recues
var ReceivedData []byte
var ReceivedDataPtr *[]byte = &ReceivedData

// Listener actuel
var CurrentListener net.Listener
var CurrentListenerPtr *net.Listener = &CurrentListener

// Défini si le listener est actif
var IsListening bool
var IsListeningPtr *bool = &IsListening
