package hangman

import (
	"math/rand"
)

func ChooseWord(solo bool, currentPlayer int) {
	choosenWord := ""
	if solo {
		randomIndex := rand.Intn(len(WordList))
		choosenWord = WordList[randomIndex]
	} else {
		DisplayWordChoice(currentPlayer)
		choosenWord = GetInput()
	}
	*CurrentWordPtr = []rune(choosenWord)
	*FoundLettersPtr = make([]rune, len(*CurrentWordPtr))
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

func checkLetter(letter rune) bool {
	res := false
	for i, c := range *CurrentWordPtr {
		if toMaj(c) == toMaj(letter) {
			(*FoundLettersPtr)[i] = c
			res = true
		}
	}
	return res
}

func toMaj(letter rune) rune {
	if letter >= 97 && letter <= 122 {
		letter -= 32
	}
	return letter
}

func checkIfTried(letter rune) bool {
	for _, c := range *TriedLettersPtr {
		if c == letter {
			return true
		}
	}
	return false
}

func checkWholeWord(word []rune) bool {
	if len(word) != len(*CurrentWordPtr) {
		return false
	}
	for i, c := range word {
		if toMaj((*CurrentWordPtr)[i]) != toMaj(c) {
			return false
		}
	}
	*FoundLettersPtr = *CurrentWordPtr
	return true
}

func checkWord(word []rune) bool {
	for i, c := range *CurrentWordPtr {
		if i >= len(word) {
			return false
		}
		if c != word[i] {
			return false
		}
	}

	*FoundLettersPtr = *CurrentWordPtr
	return true
}

func SliceToLetterOnly(s string) string {
	res := ""
	for _, c := range []rune(s) {
		// Verifie si c est entre 65 et 90 (majuscules) ou 97 et 122 (minuscules) ou si il est 32 (espace)  si il est 58 (:) ou si il est entre 48 et 57 (chiffres)
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || c == 32 || c == 58 || (c >= 48 && c <= 57) {
			res += string(c)
		} else {
			// Eradique les accents
			if c == 233 || c == 232 || c == 234 {
				res += "e"
			} else if c == 224 || c == 226 || c == 225 {
				res += "a"
			} else if c == 231 {
				res += "c"
			} else if c == 244 || c == 243 {
				res += "o"
			} else if c == 251 {
				res += "u"
			} else if c == 238 || c == 239 {
				res += "i"
			}
		}
	}
	return res
}

func RevealLetter(number int) {
	if number == 99 {
		(*FoundLettersPtr)[0] = (*CurrentWordPtr)[0]
		(*FoundLettersPtr)[len(*CurrentWordPtr)-1] = (*CurrentWordPtr)[len(*CurrentWordPtr)-1]
	} else {
		if number <= 0 || number >= len(*CurrentWordPtr) {
			return
		} else {
			for i := 0; i < number; i++ {
				randNum := rand.Intn(len(*CurrentWordPtr))
				(*FoundLettersPtr)[randNum] = (*CurrentWordPtr)[randNum]
			}
		}
	}
}

func checkInArray(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
