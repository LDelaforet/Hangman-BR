package hangmanServer

import "time"

// Partie chronomètre

// Lance un chronomètre
func StartTimer() time.Time {
	return time.Now()
}

// Remet le chronomètre passé en argument à zéro
func ResetTimer(timer *time.Time) {
	*timer = time.Now()
}

// Retourne le temps écoulé depuis le début du chronomètre
func CheckTimer(start time.Time) int {
	return int(time.Since(start).Seconds())
}

// Partie autres

// Attend un certain nombre de secondes
func WaitSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// Attends un certain nombre de millisecondes
func WaitMilliseconds(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// Attend un certain nombre de secondes mais vérifie une condition passée en argument toutes les secondes
// Une des fonctions les + stylées que j'ai écrites, je suis fan
func WaitSecondsWithCondition(seconds int, condition func() bool) bool {
	for i := 0; i < seconds; i++ {
		if condition() {
			return true
		}
		WaitSeconds(1)
	}
	return false
}

// Met la variable passée en argument à une nouvelle valeur après un certain nombre de secondes si stopVar est toujours à false
func LateSet[T any](varToSet *T, newValue T, ms int, bypassVar *bool) {
	for i := 0; i < ms; i += 50 {
		// Si bypassVar est à true, on arrête le compteur et on set directement
		if *bypassVar {
			*bypassVar = false
			break
		}
		WaitMilliseconds(50)
	}
	*varToSet = newValue
}
