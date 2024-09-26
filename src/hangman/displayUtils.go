package hangman

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	tsize "github.com/kopoli/go-terminal-size"
)

func initASCII() {
	VarInit()
}

func DisplayMainMenu() {
	fmt.Println(ASCIIToCenter(ASCIIArts["title"]))
	fmt.Println(ToCenter("Bienvenue dans le jeu du pendu !"))
	fmt.Println(ToCenter(" Veuillez choisir une option:\n"))
	fmt.Println(ToCenter("1. Jouer seul"))
	fmt.Println(ToCenter("2. Jouer à deux"))
	fmt.Println(ToCenter("3. Voir le leaderboard"))
	fmt.Println(ToCenter("99. Quitter"))
}

func ToCenter(s string) string {
	width, _ := GetSize()
	return strings.Repeat(" ", (width-len(s))/2) + s
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

func DisplayWordChoice(currentPlayer int) {
	fmt.Println(ToCenter("Joueur " + strconv.Itoa(currentPlayer) + ": Choisissez un mot à faire deviner à votre adversaire."))
	fmt.Print(ToCenter(":"))
}

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

func DisplayWord() {
	fmt.Println(ToCenter("Essai n°" + strconv.Itoa(*TriesPtr) + "."))
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

func DisplayHangman() {
	fmt.Println(ASCIIToCenter(ASCIIArts["lifeCounter_"+strconv.Itoa(9-RemainingLives)]))
}

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

func DisplayLeaderBoard() {
	scores := SortLeaderboard(readLeaderBoard())
	fmt.Println(ToCenter("Leaderboard:"))
	fmt.Println(ToCenter("+" + strings.Repeat("-", 32) + "+" + strings.Repeat("-", 9) + "+"))
	fmt.Println(ToCenter("| " + "Nom" + strings.Repeat(" ", 27) + " | " + "Score" + strings.Repeat(" ", 2) + " |"))
	fmt.Println(ToCenter("+" + strings.Repeat("-", 32) + "+" + strings.Repeat("-", 9) + "+"))

	for _, entry := range scores {
		score := entry.score
		name := entry.name

		strscore := strconv.Itoa(score)
		fmt.Println(ToCenter("| " + name + strings.Repeat(" ", 30-len(name)) + " | " + strscore + strings.Repeat(" ", 7-len(strscore)) + " |"))
	}
	fmt.Println(ToCenter("+" + strings.Repeat("-", 32) + "+" + strings.Repeat("-", 9) + "+"))
}
