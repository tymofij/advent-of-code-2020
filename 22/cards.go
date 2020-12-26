package main

import "fmt"

func score(winner []int) int {
	s := 0
	for i := len(winner); i >= 1; i-- {
		s += winner[len(winner)-i] * i
	}
	return s

}

func combat(player1, player2 []int) (winnerDeck []int) {
	for len(player1) > 0 && len(player2) > 0 {
		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]
		if card1 > card2 {
			player1 = append(player1, card1)
			player1 = append(player1, card2)
		} else {
			player2 = append(player2, card2)
			player2 = append(player2, card1)
		}
	}
	if len(player1) > 0 {
		winnerDeck = player1
	} else {
		winnerDeck = player2
	}
	return winnerDeck
}

var debug bool

func log(args ...interface{}) {
	if debug {
		fmt.Println(args...)
	}
}

type game struct {
	winner int
}

var seenGames = map[string]game{}

func recursiveCombat(player1, player2 []int) (winner int, winnerDeck []int) {
	gameKey := fmt.Sprint(player1, player2)
	gameResult, ok := seenGames[gameKey]
	if ok {
		return gameResult.winner, []int{}
	}

	seenStates := map[string]bool{}
	log("\n=== Game ===")
	round := 1
	for len(player1) > 0 && len(player2) > 0 {
		log(fmt.Sprintf("-- Round %d --", round))
		log("Player 1's deck:", player1)
		log("Player 2's deck:", player2)
		card1 := player1[0]
		card2 := player2[0]
		roundKey := fmt.Sprint(player1, player2)
		seen := seenStates[roundKey]
		if seen {
			log("Reason: Loop")
			seenGames[gameKey] = game{1}
			return 1, player1
		}
		seenStates[roundKey] = true
		player1 = player1[1:]
		player2 = player2[1:]
		log("Player 1 plays:", card1)
		log("Player 2 plays:", card2)
		if len(player1) < card1 || len(player2) < card2 {
			log("Reason: Not enough cards")
			if card1 > card2 {
				winner = 1
			} else {
				winner = 2
			}
		} else {
			subdeck1 := make([]int, card1) // Daammmn, should have read the rules more carefully!
			subdeck2 := make([]int, card2) // it is _not_ the rest of the deck, but only some of it
			copy(subdeck1, player1)
			copy(subdeck2, player2)
			winner, _ = recursiveCombat(subdeck1, subdeck2)
		}
		log(fmt.Sprintf("Player %d wins the round!\n", winner))
		if winner == 1 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}
		round++
	}
	if len(player1) > 0 {
		seenGames[gameKey] = game{1}
		return 1, player1
	} else {
		seenGames[gameKey] = game{2}
		return 2, player2
	}
}

func main() {
	// Demo data
	// player1 := []int{9, 2, 6, 3, 1}
	// player2 := []int{5, 8, 4, 7, 10}

	// Loop data
	// player1 := []int{43, 19}
	// player2 := []int{2, 29, 14}

	player1 := []int{44, 31, 29, 48, 40, 50, 33, 14, 10, 30, 5, 15, 41, 45, 12, 4, 3, 17, 36, 1, 23, 34, 38, 16, 18}
	player2 := []int{24, 20, 11, 32, 43, 9, 6, 27, 35, 2, 46, 21, 7, 49, 26, 39, 8, 19, 42, 22, 47, 28, 25, 13, 37}

	winnerDeck := combat(player1, player2)
	fmt.Println("Winner score of simple combat:", score(winnerDeck))
	fmt.Println()

	winner, winnerDeck := recursiveCombat(player1, player2)
	fmt.Println("Winner of recursive combat:", winner)
	fmt.Println("Winner's deck:", winnerDeck)
	fmt.Println("Winner's score:", score(winnerDeck))

}
