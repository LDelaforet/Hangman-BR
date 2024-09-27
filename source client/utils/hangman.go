package hangmanClient

import "strings"

func MainMenu() {
	settingsRead()
	for {
		ClearScreen()
		DisplayMainMenu()
		input := GetInput(2, "digitsOnly", 0)
		switch input {
		case "1":
			changeName()
		case "2":
			changeName()
		case "9":
			return
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

func changeIPandPort() {
	ClearScreen()
	DisplayServerMenu()
	IpPos := 1
	inputSize := 3
	ip := []string{"0", "0", "0", "0"}

	port = *LastPortPtr
	for {
		if IpPos < 5 {
			inputSize = 3
		} else {
			inputSize = 5
		}
		input := GetInput(inputSize, "digitsOnly", IpPos)
		// Si l'input commence par TORIGHT, on déplace le curseur à droite
		if strings.HasPrefix(input, "TORIGHT") {
			ip[IpPos-1] = input[7:]
			IpPos++
			// Si l'input commence par TOLEFT, on déplace le curseur à gauche
		} else if strings.HasPrefix(input, "TOLEFT") {
			ip[IpPos-1] = input[6:]
			IpPos--
			// Si l'input commence par TOPORT, on décalle vers la case du port
		} else if strings.HasPrefix(input, "TOPORT") {
			ip[IpPos-1] = input[6:]
			IpPos = 5
			// Si l'input commence par TOIP, on décalle vers la case de l'ip, la dernière case
		} else if strings.HasPrefix(input, "TOIP") {
			port = input[4:]
			IpPos = 4
		} else if strings.HasPrefix(input, "FINISHED") {
			if IpPos == 5 {
				port = input[8:]
			} else {
				ip[IpPos-1] = input[8:]
			}

			*LastIPPtr = strings.Join(ip, ".")
			*LastPortPtr = port
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
	PrintDebug(settings)
	WriteFile("settings.txt", settings)
}
