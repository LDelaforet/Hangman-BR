package hangman

import (
	"os"
	"strings"
)

// Lis le fichier donné en paramètre et retourne son contenu
func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	return string(content), err
}

// Écrit le contenu donné dans le fichier donné en paramètre
func WriteFile(fileName string, content string) error {
	return os.WriteFile(fileName, []byte(content), 0644)
}

// Copie un fichier dans un autre
func CopyFile(src string, dest string) error {
	content, err := ReadFile(src)
	if err != nil {
		return err
	}
	return WriteFile(dest, content)
}

// Supprime un fichier
func DeleteFile(fileName string) error {
	return os.Remove(fileName)
}

// Verifie si un fichier existe
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

// Verifie si un fichier est vide
func IsFileEmpty(fileName string) bool {
	content, _ := ReadFile(fileName)
	return content == ""
}

// Prend une chaines de caractères et la coupe a chaque \n\r
func SplitAndFormatLines(content string) []string {
	cont := strings.Split(content, "\n")
	for i, line := range cont {
		cont[i] = SliceToLetterOnly(string(line))
	}
	return cont
}
