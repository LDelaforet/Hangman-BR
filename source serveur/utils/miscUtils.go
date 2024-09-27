package hangmanServer

import (
	"fmt"
	"runtime"
	"strings"
)

// Vérifie si une valeur est dans une slice
func ContainsInSlice[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Partie debug

// Affiche un message de debug si DebugMode est activé.
func PrintDebug(v ...interface{}) {
	if DebugMode {
		funcName := getFunctionName()
		fmt.Printf("[%s] ", funcName)
		fmt.Println(v...)
	}
}

// Retourne le nom de la fonction courante.
func getFunctionName() string {
	pc, _, _, ok := runtime.Caller(2) // 2 pour obtenir l'appelant de printDebug
	if !ok {
		return "Unknown"
	}
	funcName := runtime.FuncForPC(pc).Name()
	return funcName[strings.LastIndex(funcName, ".")+1:]
}
