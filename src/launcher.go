package main

import (
	"fmt"
	hangman "hangman"
	"os"
)

func main() {
	fichier := "words.txt"
	args := os.Args[1:]
	for i, arg := range args {
		if arg == "--help" {
			fmt.Println("Utilisation: hangman [paramètres]")
			fmt.Println("paramètres:")
			fmt.Println("  --help: Affiche ce message d'aide")
			fmt.Println("  --file: Specifie le fichier de mots à utiliser")
			return
		}
		if arg == "--file" {
			fichier = args[i+1]
		}
	}
	if !hangman.FileExists(fichier) {
		fmt.Println("ERREUR: Le fichier " + fichier + " n'existe pas.")
		return
	}
	hangman.MainProgram(fichier)
}
