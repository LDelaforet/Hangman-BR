package hangmanClient

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	tsize "github.com/kopoli/go-terminal-size"
	"golang.org/x/term"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procKeybdEvent       = user32.NewProc("keybd_event")
	VK_F11          byte = 0x7A
	KEYEVENTF_KEYUP      = 0x0002
)

func InitASCII() {
	AsciiArtsInit()
}

func SetConsoleTitle(title string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "title", title)
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("printf", "\033]0;%s\007", title)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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
	fmt.Print(ASCIIToCenter(ASCIIArts["boxIP"])[0 : len(ASCIIToCenter(ASCIIArts["boxIP"]))-1])
	MoveCursorRelative(-7, -13)
}

// Ajoute du surlignage au texte passé en paramètre
func ToHighlight(text string) string {
	return ("\033[7m" + text + "\033[0m")
}

func PressF11() {
	procKeybdEvent.Call(uintptr(VK_F11), 0, 0, 0)
	time.Sleep(100 * time.Millisecond)
	procKeybdEvent.Call(uintptr(VK_F11), 0, uintptr(KEYEVENTF_KEYUP), 0)
}

// Retourne une chaine de caractère qui ajoute ou retire du poids pour la fonction ToCenter
func Weight(weight int) string {
	if weight < 0 {
		return strings.Repeat("`", 0-weight)
	}
	return strings.Repeat("\x00", weight)
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

// Déplace le curseur vers des coordonées absolues
func MoveCursorAbsolute(row int, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
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
	// Parcours la liste s
	sf := ""
	count := 0
	for _, char := range s {
		if char == '`' {
			count -= 1
			continue
		}
		count += 1
		sf += string(char)
	}
	width, _ := GetSize()
	return strings.Repeat(" ", (width-count)/2) + sf
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

func TitleDebug(v ...interface{}) {
	if DebugMode {
		toTitle := ""
		funcName := getFunctionName()
		toTitle = fmt.Sprintf("[%s] ", funcName)
		for _, val := range v {
			toTitle += fmt.Sprintf("%v", val)
		}
		SetConsoleTitle(toTitle)
	}
}
