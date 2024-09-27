package hangmanClient

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"

	tsize "github.com/kopoli/go-terminal-size"
	"golang.org/x/term"
)

func InitASCII() {
	AsciiArtsInit()
}

func DisplayMainMenu() {
	//fmt.Println(ASCIIToCenter(ASCIIArts["title"]))
	fmt.Println(ASCIIToCenter(ASCIIArts["title"]))
	fmt.Println(ToCenter("Bienvenue dans:"))
	fmt.Println(ToCenter("Hangman Battle Royale\n"))
	fmt.Println(ToCenter(" Veuillez choisir une option:"))
	fmt.Println(ToCenter("1. Rejoindre un serveur"))
	fmt.Println(ToCenter("2. Changer son nom"))
	fmt.Println(ToCenter("9. Quitter"))
	fmt.Print(ToCenter("Choix : " + string(rune(0))))
}

func DisplayNameMenu() {
	fmt.Println(ASCIIToCenter(ASCIIArts["title"]))
	fmt.Println(ToCenter("Nom actuel: " + *PlayerNamePtr))
	fmt.Print(ASCIIToCenter(ASCIIArts["boxName"])[0 : len(ASCIIToCenter(ASCIIArts["boxName"]))-1])
	// I35 -> H12: haut:1, gauche:23
	MoveCursorRelative(-1, -23)
}

func DisplayServerMenu() {
	fmt.Println(ASCIIToCenter(ASCIIArts["title"]))
	fmt.Println(ToCenter("Liste des serveurs disponibles:"))
}

// Déplace le curseur de la console relativement à sa position actuelle
func MoveCursorRelative(rowOffset int, colOffset int) {
	if rowOffset > 0 {
		fmt.Printf("\033[%dB", rowOffset) // Déplacer vers le bas
	} else if rowOffset < 0 {
		fmt.Printf("\033[%dA", -rowOffset) // Déplacer vers le haut
	}
	if colOffset > 0 {
		fmt.Printf("\033[%dC", colOffset) // Déplacer vers la droite
	} else if colOffset < 0 {
		fmt.Printf("\033[%dD", -colOffset) // Déplacer vers la gauche
	}
}

// Renvoi la position du curseur
func GetCursorPosition() (row int, col int, err error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, 0, err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print("\033[6n")

	var buf [16]byte
	n, err := os.Stdin.Read(buf[:])
	if err != nil {
		return 0, 0, err
	}
	response := string(buf[:n])
	if !strings.HasPrefix(response, "\033[") || !strings.HasSuffix(response, "R") {
		return 0, 0, fmt.Errorf("réponse inattendue du terminal: %s", response)
	}
	fmt.Sscanf(response, "\033[%d;%dR", &row, &col)
	return row, col, nil
}

func ToCenter(s string) string {
	width, _ := GetSize()
	// Oui psq len sur une string compte le nombre d'octets et pas le nombre de runes, donc tt ce qui est pas ascii est compté en double (au moins)
	strLen := utf8.RuneCountInString(s)
	return strings.Repeat(" ", (width-strLen)/2) + s
}

func ASCIIToCenter(s string) string {
	asc := ""
	for _, line := range strings.Split(s, "\n") {
		asc += ToCenter(line) + "\n"
	}
	return asc
}

func GetSize() (Width int, Height int) {
	var s tsize.Size

	s, _ = tsize.GetSize()
	Width, Height = s.Width, s.Height
	return
}

// Affiche les lettres déjà essayées
func DisplayTried() {
	fmt.Println(ToCenter("Lettres déjà essayées: "))
	tried := ""
	for index, letter := range *TriedLettersPtr {
		if index == len(*TriedLettersPtr)-1 {
			tried += string(letter)
		} else {
			tried += string(letter) + ", "
		}
	}
	fmt.Println(ToCenter(tried))
}

// Affiche le mot à deviner
func DisplayWord() {
	firstCol := ""
	secondCol := ""
	for _, letter := range *FoundLettersPtr {
		firstCol += "+-"
		if letter == 0 {
			secondCol += "|."
		} else {
			secondCol += "|" + string(letter)
		}
	}
	firstCol += "+"
	secondCol += "|"
	fmt.Println(ToCenter(firstCol))
	fmt.Println(ToCenter(secondCol))
	fmt.Println(ToCenter(firstCol))
}

// Affiche le pendu
func DisplayHangman() {
	fmt.Println(ASCIIToCenter(ASCIIArts["lifeCounter_"+strconv.Itoa(9-RemainingLives)]))
}

// Vide l'écran
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Affiche un message de debug si DebugMode est activé.
func PrintDebug(v ...interface{}) {
	if DebugMode {
		funcName := getFunctionName()
		fmt.Printf("[%s] ", funcName)
		fmt.Println(v...)
	}
}
