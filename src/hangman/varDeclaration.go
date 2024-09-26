package hangman

import (
	"fmt"
	"os"
	"strings"
)

// Détermine si le mode debug est activé
var DebugMode bool = false
var DebugModePtr = &DebugMode

// Contient la liste des mots possibles
var WordList []string
var WordListPtr = &WordList

// Contient la liste des ASCII arts
var ASCIIArts map[string]string
var ASCIIArtsPtr = &ASCIIArts

// Nombre de vies restantes au joueur
var RemainingLives int
var RemainingLivesPtr = &RemainingLives

// Mot actuel en slices de runes
var CurrentWord []rune
var CurrentWordPtr = &CurrentWord

// Lettres déjà trouvées dans le mot
var FoundLetters []rune
var FoundLettersPtr = &FoundLetters

// Lettres déjà essayées
var TriedLetters []rune
var TriedLettersPtr = &TriedLetters

// Nombre d'essais
var Tries int
var TriesPtr = &Tries

// Score actuel
var Score int
var ScorePtr = &Score

// Nom du fichier contenant les mots
var FileName string
var FileNamePtr = &FileName

// Leaderboard filename
var LeaderboardFileName string
var LeaderBoardFileNamePtr = &LeaderboardFileName

type LeaderboardEntry struct {
	name  string
	score int
}

// Initialise toute les déclarations de variables
func VarInit() {
	wordListInit()
	AsciiArtsInit()
	*RemainingLivesPtr = 9
	*LeaderBoardFileNamePtr = "leaderboard_" + strings.Split(FileName, ".txt")[0] + ".txt"
}

// Lis le fichier containant les mots et les ajoute à la liste
func wordListInit() {
	wordListFile, err := ReadFile(*FileNamePtr)
	if err != nil {
		panic(err)
	}
	*WordListPtr = SplitAndFormatLines(wordListFile)
}

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
