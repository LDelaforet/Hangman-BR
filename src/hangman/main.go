package hangman

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var timer time.Time
var scoreFileContent string
var currentPlayer int // Joueur actuel
var player1Lost bool  // Joueur 1 a perdu
var player2Lost bool  // Joueur 2 a perdu

// Valeurs pour le score
var totalFound int  // Nombre de mots trouvés
var totalLenght int // Longueur totale des mots trouvés
var totalLives int  // Nombre total de vies perdues
var totalTimer int  // Temps total passé à jouer

// Initialise toutes les valeurs de base
func MainProgram(filename string) {
	*FileNamePtr = filename
	VarInit()
	initASCII()
	MainLoop()
}

func playerSwitch(p int) int {
	if p == 1 {
		return 2
	}
	return 1
}

func MainLoop() {
	soloPlay := true // On joue seul si il est sur true

	for { // Boucle principale
		for { // Boucle du menu principal
			ClearScreen()
			DisplayMainMenu()
			fmt.Print(ToCenter("Choix : " + string(rune(0))))
			chx := GetInput()
			ClearScreen()
			// check si chx est dans les options
			// Equivalent a if chx in []string{"1", "2", "3", "99"} { :
			if checkInArray(chx, []string{"1", "2", "3", "99"}) {
				if chx == "99" {
					return
				}
				if chx == "3" {
					DisplayLeaderBoard()
					fmt.Print("Appuyez sur entrée pour continuer.")
					GetInput()
				}
				if chx == "1" {
					soloPlay = true
					break
				}
				if chx == "2" {
					soloPlay = false
					break
				}
			}
		}

		// Reset les valeurs utilisées pour le score
		totalFound = 0
		totalLenght = 0
		totalLives = 0
		totalTimer = 0
		player1Lost = false
		player2Lost = false
		currentPlayer = 1

		for {
			// if soloPlay {currentPlayer = playerSwitch(currentPlayer)}
			// Un mot aléatoire est choisi
			ChooseWord(soloPlay, playerSwitch(currentPlayer))
			*RemainingLivesPtr = 9
			*TriedLettersPtr = make([]rune, 0)
			*TriesPtr = 1

			RevealLetter(2) // Révèle le nombre donné en argument de lettres, si l'argument est 99: révèle la première et la dernière lettre
			if soloPlay {
				timer = StartTimer()
			} else {
				ClearScreen()
				fmt.Println(ToCenter("Le joueur " + strconv.Itoa(currentPlayer) + " doit trouver le mot"))
				fmt.Print(ToCenter("Appuyez sur entrée pour continuer"))
				GetInput()
			}
			// Boucle de guess
			for {
				// Clear l'écran et affiche les infos du jeu
				ClearScreen()
				// fmt.Println(ToCenter(string(*CurrentWordPtr)))
				DisplayWord()
				DisplayHangman()
				DisplayTried()
				if soloPlay {
					fmt.Println(ToCenter("Temps écoulé: " + strconv.Itoa(int(time.Since(timer).Seconds()))))
				}
				fmt.Print(ToCenter("Choix : " + string(rune(0))))
				input := []rune(GetInput())

				// Check si le mot est: vide, une lettre ou un mot
				if len(input) == 0 {
					continue
				} else if len(input) == 1 {
					if !checkLetter(input[0]) {
						if !checkIfTried(input[0]) {
							*RemainingLivesPtr -= 1
						}
					}
					*TriedLettersPtr = append(*TriedLettersPtr, input[0])
				} else {
					if !checkWholeWord(input) {
						*RemainingLivesPtr -= 2
					}
				}

				// Incrémente le nombre d'essais
				*TriesPtr++

				// Si le mot formé par les lettres de foundLetters est le mot mystère
				if checkWord(FoundLetters) {
					ClearScreen()
					fmt.Println(ToCenter("Vous avez gagné !"))
					break
				}

				// Si le joueur n'a plus de vies
				if RemainingLives <= 0 {
					ClearScreen()
					fmt.Println(ToCenter("Vous avez perdu !"))
					break
				}
			}
			if soloPlay {
				if RemainingLives <= 0 {
					fmt.Println(ToCenter("Le mot était: " + string(CurrentWord)))
					fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
					GetInput()
					break
				} else {
					totalFound++
					totalLenght += len(CurrentWord)
					totalLives += RemainingLives
					totalTimer += StopTimer(timer)
					// Affiche les stats
					ClearScreen()
					fmt.Println(ToCenter("Vous avez trouvé le mot !"))
					fmt.Println(ToCenter("Le mot était: " + string(CurrentWord)))
					fmt.Println(ToCenter("Il vous restait " + strconv.Itoa(RemainingLives) + " vies."))
					fmt.Println(ToCenter("Vous avez trouvé le mot en " + strconv.Itoa(*TriesPtr) + " essais."))
					fmt.Println(ToCenter("Temps écoulé: " + strconv.Itoa(int(time.Since(timer).Seconds()))))
					fmt.Println(ToCenter("Score actuel: " + strconv.Itoa((ScoreCalc(totalFound, totalLenght, totalLives, totalTimer)))))
					fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
					GetInput()
				}
			} else {
				if RemainingLives <= 0 {
					fmt.Println(ToCenter("Le mot était: " + string(CurrentWord) + "\n"))
					if currentPlayer == 1 {
						player1Lost = true
					} else {
						player2Lost = true
					}
					if player1Lost && player2Lost {
						fmt.Println(ToCenter("Egalité !"))
						fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
						GetInput()
						ClearScreen()
						break
					} else {
						fmt.Println(ToCenter("Le joueur " + strconv.Itoa(playerSwitch(currentPlayer)) + " prends la relève."))
						fmt.Println(ToCenter("Si il trouve le mot, il a gagné."))
						fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
						GetInput()
						ClearScreen()
					}
					currentPlayer = playerSwitch(currentPlayer)
				} else {
					won := false
					if currentPlayer == 1 {
						player1Lost = false
						if player2Lost {
							won = true
						}
					} else {
						player2Lost = false
						if player1Lost {
							won = true
						}
					}

					if won {
						fmt.Println(ToCenter("Félicitations !"))
						fmt.Println(ToCenter("Le joueur " + strconv.Itoa(currentPlayer) + " à gagné."))
						fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
						GetInput()
						ClearScreen()
						break
					} else {
						fmt.Println(ToCenter("Au tour du joueur" + strconv.Itoa(playerSwitch(currentPlayer)) + "."))
						fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
						GetInput()
						currentPlayer = playerSwitch(currentPlayer)
						ClearScreen()
					}
				}
			}
		}
		if soloPlay {
			ClearScreen()
			fmt.Println(ToCenter("Vous avez trouvé " + strconv.Itoa(totalFound) + " mots."))
			fmt.Println(ToCenter("Vous avez trouvé " + strconv.Itoa(totalLenght) + " lettres."))
			fmt.Println(ToCenter("Vous avez perdu " + strconv.Itoa(totalLives) + " vies."))
			fmt.Println(ToCenter("Vous avez mis " + strconv.Itoa(totalTimer) + " secondes.\n"))

			if totalFound > 0 {
				fmt.Println(ToCenter("Votre score final est de: " + strconv.Itoa((ScoreCalc(totalFound, totalLenght, totalLives, totalTimer)))))

				// Ajout au leaderboard
				fmt.Print(ToCenter("Entrez votre nom pour le leaderboard: " + strings.Repeat(string(rune(0)), 5)))
				name := GetInput()
				AddToLeaderboard(name, ScoreCalc(totalFound, totalLenght, totalLives, totalTimer))

				DisplayLeaderBoard()
				fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
				GetInput()
			} else {
				fmt.Println(ToCenter("Vous n'avez pas trouvé de mots, vous ne pouvez donc pas entrer dans le leaderboard."))
				fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
				GetInput()
			}
		} else {
			fmt.Println(ToCenter("Jouez en solo pour pouvoir profiter du leaderboard."))
			fmt.Print(ToCenter("Appuyez sur entrée pour continuer."))
			GetInput()
		}
	}
}
