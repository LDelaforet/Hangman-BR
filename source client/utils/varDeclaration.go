package hangmanClient

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	Name             string   // Nom du serveur
	Description      string   // Description du serveur
	Wordlists        []string // Listes de mots disponibles
	PlayerCount      int      // Nombre de joueurs
	IsListened       bool     // Si le serveur est sur écoute
	ReceivedBuffer   []byte   // Buffer de données reçues
	RequestResponded bool     // Si la dernière requête a été répondue
	LastRequest      string   // Dernière requête envoyée
	LateSetBypass    bool     // Termine le LateSet tout de suite
	Connection       net.Conn // Connexion au serveur
}

// Contient la liste des ASCII arts
var ASCIIArts map[string]string
var ASCIIArtsPtr = &ASCIIArts

// Détermine si le mode debug est activé
var DebugMode bool = true
var DebugModePtr = &DebugMode

// Connexion actuelle
var CurrentServer Server
var CurrentServerPtr *Server = &CurrentServer

// Données recues
var ReceivedData []byte
var ReceivedDataPtr *[]byte = &ReceivedData

// Nom du joueur
var PlayerName string
var PlayerNamePtr *string = &PlayerName

// Dernière ip utilisée
var LastIP string
var LastIPPtr *string = &LastIP

// Dernier port utilisé
var LastPort string
var LastPortPtr *string = &LastPort

// A RECUPERER AU SERVEUR

// Lettres déjà essayées
var TriedLetters []rune
var TriedLettersPtr = &TriedLetters

// Lettres déjà trouvées dans le mot
var FoundLetters []rune
var FoundLettersPtr = &FoundLetters

// Nombre de vies restantes au joueur
var RemainingLives int
var RemainingLivesPtr = &RemainingLives

func AsciiArtsInit() {
	// Initialise ASCIIArts
	*ASCIIArtsPtr = make(map[string]string)

	// Lis la liste des fichiers du dossier ASCII_arts
	files, err := os.ReadDir("./ASCII_arts")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier ASCII_arts")
	}

	// Associe chaque fichier à son contenu dans la map ASCIIArts
	for _, file := range files {
		asciiArt, err := ReadFile("./ASCII_arts/" + file.Name())
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier " + file.Name())
		}
		(*ASCIIArtsPtr)[file.Name()] = asciiArt
	}
}
