package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Fonction pour obtenir un mot aléatoire à partir d'un fichier
func GetRandomWord(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'ouverture du fichier : %v", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du fichier : %v", err)
	}

	if len(words) == 0 {
		return "", fmt.Errorf("le fichier est vide")
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex], nil
}

func playHangman() {
	word, err := GetRandomWord("words.txt")
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du mot aléatoire : %v", err)
	}

	clear()
	fmt.Printf("Le mot à chercher possède %d lettres.\n", len(word))

	guessedLetters := make(map[rune]bool)
	wrongLetters := []rune{} // Liste pour stocker les lettres fausses
	tries := 6

	// Crée un tableau de runes pour représenter le mot deviné
	guessedWord := make([]rune, len(word))

	// Affiche la première lettre et masque le reste
	guessedWord[0] = rune(word[0])

	if len(word) >= 10 {
		// Sélectionne une deuxième lettre aléatoire dans le mot, mais pas la première
		var secondIndex int
		for {
			secondIndex = rand.Intn(len(word))
			if secondIndex != 0 { // S'assure que ce n'est pas le premier index
				break
			}
		}
		guessedWord[secondIndex] = rune(word[secondIndex]) // Affiche la deuxième lettre
	}

	for i := 1; i < len(guessedWord); i++ {
		if guessedWord[i] == 0 { // Si cette position n'a pas été définie, la masquer
			guessedWord[i] = '_'
		}
	}

	for tries > 0 {
		clear()
		StickMan(6 - tries)
		fmt.Println("Mot :", string(guessedWord))

		// Affiche les lettres fausses
		fmt.Print("Lettres fausses : [ ")
		for _, letter := range wrongLetters {
			fmt.Printf("%c ", letter)
		}
		fmt.Println("]")

		fmt.Printf("Il te reste %d essais. Devine une lettre : ", tries)
		var guess string
		fmt.Scanln(&guess)

		if len(guess) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		letter := rune(guess[0])
		if guessedLetters[letter] {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		guessedLetters[letter] = true

		found := false
		for i, c := range word {
			if c == letter {
				guessedWord[i] = letter
				found = true
			}
		}

		if found {
			fmt.Println("Bien joué !")
		} else {
			tries--
			wrongLetters = append(wrongLetters, letter) // Ajoute la lettre incorrecte
			fmt.Println("Mauvaise lettre.")
		}

		if string(guessedWord) == word {
			fmt.Printf("Félicitations, vous avez deviné le mot : %s\n", word)
			break
		}
	}

	if tries == 0 {
		clear()
		StickMan(6)
		fmt.Printf("Dommage, vous avez perdu. Le mot était : %s\n", word)
	}
}
