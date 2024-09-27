package hangmanClient

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unicode"

	"golang.org/x/term"
)

func StringToHex(s string) string {
	hex := ""
	for _, c := range s {
		hex += fmt.Sprintf("%02x", c)
	}
	return hex
}

// Lis l'input utilisateur selon des filtres et une longueur maximale et le renvoi en string
func GetInput(maxLength int, inputType string, ipPos int) string {
	var input strings.Builder
	// Sauvegarde de l'état actuel du terminal
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return ""
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// Gérer les signaux pour restaurer l'état du terminal à la sortie
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()
	pos := 0
	modeView := false
	for {
		var char [1]byte                 // Déclare un tableau de bytes de taille 1
		_, err := os.Stdin.Read(char[:]) // Lire un caractère
		if err != nil {
			return ""
		}

		if modeView {
			fmt.Println("Char: ", char)
			continue
		}

		// Gestion des caractères spéciaux (flèches, etc.)
		if char[0] == 27 {
			var seq [2]byte
			_, err := os.Stdin.Read(seq[:])
			if err != nil {
				return ""
			}
			if seq[0] == '[' && seq[1] == 'D' {
				if pos > 0 {
					MoveCursorRelative(0, -1)
					pos--
				}
				continue // Passer à l'itération suivante
			}
			if seq[0] == '[' && seq[1] == 'C' {
				if pos < input.Len() {
					MoveCursorRelative(0, 1)
					pos++
				}
				continue // Passer à l'itération suivante
			}
			continue
		}

		if char[0] == '\n' || char[0] == '\r' { // Fin de la saisie
			break
		}

		if char[0] == 127 { // Touche Retour Arrière (Backspace, ASCII 127)
			if input.Len() > 0 && pos > 0 { // Vérifiez s'il y a des caractères à supprimer et si la position du curseur est > 0
				inputStr := input.String()

				// Supprimer le caractère avant la position du curseur
				input.Reset()
				input.WriteString(inputStr[:pos-1]) // Conserver tout avant le curseur
				if pos < len(inputStr) {
					input.WriteString(inputStr[pos:]) // Conserver tout après le curseur
				}

				// Déplacer le curseur d'une position vers la gauche
				MoveCursorRelative(0, -1)

				// Effacer le caractère supprimé et réafficher le reste
				fmt.Print(inputStr[pos:])
				fmt.Print(" ")                                // Effacer le dernier caractère visuellement
				MoveCursorRelative(0, -len(inputStr[pos:])-1) // Replacer le curseur

				// Mettre à jour la position du curseur
				pos--
			}
			continue // Passer à l'itération suivante
		}

		if input.Len() == maxLength {
			continue
		}

		filter := false
		switch inputType {
		case "lettersNoSpaces":
			filter = unicode.IsLetter(rune(char[0])) && char[0] < 128
		case "digitsOnly":
			filter = unicode.IsDigit(rune(char[0]))
		case "lettersAndDigits":
			filter = unicode.IsLetter(rune(char[0])) || unicode.IsDigit(rune(char[0]))
		default:
			filter = true // Accepter tout autre caractère
		}

		if filter {
			// Récupère la chaîne actuelle de l'input
			inputStr := input.String()
			input.Reset()

			// Conserver tout avant le curseur
			if pos > 0 {
				input.WriteString(inputStr[:pos]) // Conserver tout avant la position du curseur
			}

			// Ajouter le caractère à la position actuelle
			input.WriteByte(char[0])
			pos++ // Avancer la position du curseur après l'ajout

			// Conserver tout après le curseur
			if pos <= len(inputStr) {
				input.WriteString(inputStr[pos-1:]) // Conserver tout après le curseur
			}

			// Afficher uniquement la partie de la chaîne modifiée et le reste
			fmt.Print(string(char[0])) // Affiche le caractère ajouté

			if pos <= len(inputStr) {
				// Affiche la partie après le curseur qui reste inchangée
				fmt.Print(inputStr[pos-1:])
				// Replacer le curseur après l'ajout
				MoveCursorRelative(0, -len(inputStr[pos-1:]))
			}
		}

	}

	return input.String()
}
