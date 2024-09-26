package hangman

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Alors oui c'etait pas du tout demandé mais c'etait faisable donc je l'ai fait

func readLeaderBoard() []LeaderboardEntry {
	scores := []LeaderboardEntry{}

	if FileExists(LeaderboardFileName + ".bak") {
		if IsFileEmpty(LeaderboardFileName) {
			if FileExists(LeaderboardFileName + ".bak") {
				CopyFile(LeaderboardFileName+".bak", LeaderboardFileName)
			} else {
				WriteFile(LeaderboardFileName, "a: 0")
			}
		}
	} else {
		WriteFile(LeaderboardFileName, "a: 0")
	}

	fileContent, err := ReadFile(LeaderboardFileName)
	if err != nil {
		panic(err)
	}
	board := SplitAndFormatLines(fileContent)
	for _, line := range board {
		if line == "a: 0" {
			continue
		}
		// Split la ligne en deux
		split := strings.Split(line, ": ")
		score, err := strconv.Atoi(split[1])
		// Si on est en mode debug on va panic pour régler le pb sinon osef on skip tant pis
		if err != nil {
			if DebugMode {
				fmt.Print("Split: ")
				fmt.Println(split)
				panic(err)
			} else {
				continue
			}
		}
		scores = append(scores, LeaderboardEntry{name: split[0], score: score})
	}
	return scores
}

func AddToLeaderboard(name string, score int) {
	// Fais un backup de leaderboard.txt
	CopyFile(LeaderboardFileName, LeaderboardFileName+".bak")

	scores := readLeaderBoard()
	scores = append(scores, LeaderboardEntry{name: name, score: score})
	scoreFileContent := ""
	for _, entry := range scores {
		scoreFileContent += entry.name + ": " + strconv.Itoa(entry.score) + "\n"
	}
	// Enleve les deux derniers caractères pour éviter le \n de fin
	scoreFileContent = scoreFileContent[:len(scoreFileContent)-1]
	WriteFile(LeaderboardFileName, scoreFileContent)
}

// Fonction pour trier les scores
func SortLeaderboard(scores []LeaderboardEntry) []LeaderboardEntry {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})
	return scores
}

func ScoreCalc(totalFound int, totalLength int, totalLives int, totalTimer int) int {
	viesMaxTotal := totalFound * 9 // Nombre total de vies
	bonus := 0

	if totalFound == 0 {
		bonus = 50
	}

	score := (float64(totalFound) * 20) + // Points par mot trouvé
		(float64(totalLength) * 4) + // Points par longueur totale
		(float64(totalLives) / float64(viesMaxTotal) * 150) - // Points pour les vies restantes
		(float64(totalTimer) * 1) + // Impact du temps
		float64(bonus)

	return int(score) + 10 // S'assurer que le score est toujours positif
}
