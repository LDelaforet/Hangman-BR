package hangmanServer

import "strings"

// Partie qui gere les slices de PlayerList

// Fonction très spécifique qui vérifie si l'id donné est dans la slice de Player donnée
func IsIdInPlayerSlice(slice []Player, id int) bool {
	for _, player := range slice {
		if player.Id == id {
			return true
		}
	}
	return false
}

// Partie qui gère les requetes

// Renvoi les arguments de la requete contenu dans received
func GetArgsFromRequest(received, request string) []string {
	received = strings.TrimSuffix(received, "$")

	if parts := strings.Split(received, request); len(parts) > 1 {
		argsString := strings.TrimSpace(parts[1])
		args := strings.Fields(argsString)
		return args
	}
	return nil
}

// Partie sur les types de slices

// Fonction qui vérifie si deux slices de bytes sont égales
func AreByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Fonction qui renvoi une nouvelle slice de bytes sans les éléments de toRemoveList
func RemoveFromByteSlice(slice []byte, toRemoveList []byte) []byte {
	newSlice := []byte{}
	for _, sliceElem := range slice {
		skip := false
		for _, toRemove := range toRemoveList {
			if sliceElem == toRemove {
				skip = true
				break
			}
		}
		if !skip {
			newSlice = append(newSlice, sliceElem)
		}
	}
	return newSlice
}
