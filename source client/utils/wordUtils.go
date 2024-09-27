package hangmanClient

func stringToRunes(s string) []rune {
	return []rune(s)
}
func toMaj(letter rune) rune {
	if letter >= 97 && letter <= 122 {
		letter -= 32
	}
	return letter
}

func StringToLetterOnly(s string, keepSpace bool) string {
	res := ""
	for _, c := range []rune(s) {
		// Verifie si c est entre 65 et 90 (majuscules) ou 97 et 122 (minuscules) ou si il est 32 (espace)  si il est 58 (:) ou si il est entre 48 et 57 (chiffres)
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || c == 32 || c == 58 || (c >= 48 && c <= 57) {
			if c == 32 && !keepSpace {
				continue
			}
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

func CheckInArray(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
