package hangman

import "time"

// Lance un chronomètre
func StartTimer() time.Time {
	return time.Now()
}

func StopTimer(start time.Time) int {
	return int(time.Since(start).Seconds())
}
