package hangmanClient

import (
	"fmt"
	"strings"
)

func MainMenu() {
	//PressF11()
	settingsRead()
	for {
		ClearScreen()
		DisplayMainMenu()
		input := GetInput(2, "digitsOnly", 0)
		switch input {
		case "1":
			changeIPandPort()
			ClearScreen()
			ConnectToServer()
		case "2":
			changeName()
		case "9":
			//PressF11()
			return
		}
	}
}

func ConnectToServer() {
	// Boucle infinie pour choisir un serveur
	for {
		// Bon pour l'instant l'ip et le port sont hardcodés
		conn := ConnectSocket(*LastIPPtr, *LastPortPtr)

		// Check le serveur
		if CheckServer(conn) {
			serveur := *CurrentServerPtr
			fmt.Println("Serveur choisi, nom du serveur: ")
			SendMessageToServer(serveur, "GETNAME")
			if !WaitSecondsWithCondition(5, func() bool { return serveur.RequestResponded }) || serveur.LastRequest == "GETNAME" {
				fmt.Println("Serveur invalide")
				continue
			}
			fmt.Println(CurrentServer.Name)
			PermanantListen()
			// Si on est là, c'est que le serveur est valide
			for {
			}
		}
	}
}

func changeName() {
	ClearScreen()
	DisplayNameMenu()
	input := GetInput(20, "lettersAndDigits", 0)
	*PlayerNamePtr = StringToLetterOnly(input, false)
	settingsWrite()
}

var port string
var ip []string

func changeIPandPort() {
	ClearScreen()
	DisplayServerMenu()

	ipCoords := [2]int{0, 0}
	ipCoords[0], ipCoords[1], _ = GetCursorPosition()

	IpPos := 1
	inputSize := 3

	if *LastPortPtr != "" {
		port = *LastPortPtr
	} else {
		port = "     "
		MoveCursorRelative(0, 2)
	}

	if *LastIPPtr != "" {
		ip = strings.Split(*LastIPPtr, ".")
		for _, val := range ip {
			if len(val) == 1 {
				val = " " + val + " "
			} else if len(val) == 2 {
				val = val + " "
			}
			fmt.Print(val)
			MoveCursorRelative(0, 1)
		}
		MoveCursorRelative(6, -12)
		fmt.Print(port)
		MoveCursorRelative(1, 50)
		fmt.Println("\n" + ToCenter("Voulez vous vous connecter a ce serveur ?"))
		keepServer := false
		yesNoCoord := [2]int{0, 0}
		yesNoCoord[0], yesNoCoord[1], _ = GetCursorPosition()
		for {
			if keepServer {
				fmt.Println(ToCenter("+-----+" + "  " + "+-----+"))
				fmt.Println(ToCenter("| " + ToHighlight("OUI") + Weight(-8) + " |" + "  | NON |"))
				fmt.Println(ToCenter("+-----+" + "  " + "+-----+"))
			} else {
				fmt.Println(ToCenter("+-----+" + "  " + "+-----+"))
				fmt.Println(ToCenter("| OUI |" + "  | " + ToHighlight("NON") + Weight(-8) + " |"))
				fmt.Println(ToCenter("+-----+" + "  " + "+-----+"))
			}
			inpt := GetInput(1, "arrowOnly", 0)
			if inpt == "VALIDATE" {
				break
			} else {
				keepServer = !keepServer
			}
			MoveCursorAbsolute(yesNoCoord[0], yesNoCoord[1])
		}
		if keepServer {
			return
		} else {
			MoveCursorRelative(-4, 0)
			fmt.Println(ToCenter("                                               "))
			fmt.Println(ToCenter("                                               "))
			fmt.Println(ToCenter("                                               "))
			fmt.Println(ToCenter("                                               "))
			MoveCursorRelative(-5, 0)
			fmt.Print(ToCenter("+-------+ "))
			MoveCursorRelative(-7, -13)
		}
	} else {
		ip = []string{"0", "0", "0", "0"}
	}

	for {
		if IpPos < 5 {
			inputSize = 3
		} else {
			inputSize = 5
		}

		input := GetInput(inputSize, "digitsOnly", IpPos)
		SetConsoleTitle(input)

		// Si l'input commence par TORIGHT, on déplace le curseur à droite
		if strings.HasPrefix(input, "TORIGHT") {
			ip[IpPos-1] = input[7:]
			lastVal := input[7:]
			if len(lastVal) == 1 {
				MoveCursorRelative(0, -1)
				fmt.Print(" " + lastVal + " ")
				lastVal = "000"
			}

			if IpPos == 4 {
				MoveCursorRelative(6, ((3-len(lastVal))+1)-11)
				fmt.Print("     ")
				MoveCursorRelative(0, -5)
			} else {
				MoveCursorRelative(0, (3-len(lastVal))+1)
				fmt.Print("   ")
				MoveCursorRelative(0, -3)
			}
			IpPos++
			// Si l'input commence par TOLEFT, on déplace le curseur à gauche
		} else if strings.HasPrefix(input, "TOLEFT") {
			if IpPos == 5 {
				port = input[6:]
			} else {
				ip[IpPos-1] = input[6:]
			}
			if IpPos == 5 {
				MoveCursorRelative(-6, 7)
				fmt.Print("   ")
				MoveCursorRelative(0, -3)
			} else {
				MoveCursorRelative(0, -4)
				fmt.Print("   ")
				MoveCursorRelative(0, -3)
			}
			IpPos--
			// Si l'input commence par TOPORT, on décalle vers la case du port
		} else if strings.HasPrefix(input, "FINISHED") {
			if IpPos == 5 {
				port = input[8:]
			} else {
				ip[IpPos-1] = input[8:]
			}

			*LastIPPtr = strings.Join(ip, ".")
			*LastPortPtr = port
			ClearScreen()
			settingsWrite()
			return
		} else {
			if IpPos == 5 {
				port = input[8:]
			} else {
				ip[IpPos-1] = input[8:]
			}
		}
	}
}

func settingsRead() {
	if !FileExists("settings.txt") {
		return
	}
	settings, _ := ReadFile("settings.txt")
	lines := strings.Split(settings, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Name: ") {
			*PlayerNamePtr = line[6:]
		} else if strings.HasPrefix(line, "LastIP: ") {
			*LastIPPtr = line[8:]
		} else if strings.HasPrefix(line, "LastPort:") {
			*LastPortPtr = line[9:]
		}
	}
}

func settingsWrite() {
	settings := "Name: " + *PlayerNamePtr + "\n" + "LastIP: " + *LastIPPtr + "\n" + "LastPort: " + *LastPortPtr
	PrintDebug("\n" + settings)
	WriteFile("settings.txt", settings)
}
