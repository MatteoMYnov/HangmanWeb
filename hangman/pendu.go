package hangman

import "fmt"

// ==========Y=
// ||/       |
// ||        O
// ||       /|\
// ||       / \
///||
//==============
func StickMan(errs int) {
	if errs == 0 {
		fmt.Println(" ==========Y=\n ||/\n ||\n ||\n ||\n/||\n==============")
	} else if errs == 1 {
		fmt.Println(" ==========Y=\n ||/       |\n ||        O\n ||\n ||\n/||\n==============")
	} else if errs == 2 {
		fmt.Println(" ==========Y=\n ||/       |\n ||        O\n ||        |\n ||\n/||\n==============")
	} else if errs == 3 {
		fmt.Println(" ==========Y=\n ||/       |\n ||        O\n ||       /|\n ||\n/||\n==============")
	} else if errs == 4 {
		fmt.Println(" ==========Y=\n ||/       |\n ||        O\n ||       /|\\\n ||\n/||\n==============")
	} else if errs == 5 {
		fmt.Println(" ==========Y=\n ||/       |\n ||        O\n ||       /|\\\n ||       /\n/||\n==============")
	} else if errs == 6 {
		fmt.Println(" ==========Y=\n ||/       |\n ||        O\n ||       /|\\\n ||       / \\\n/||\n==============")
	}
}

func clear() {
	fmt.Printf("\033[H\033[2J")
}
