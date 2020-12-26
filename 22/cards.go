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

func recursiveCombat(player1, player2 []int) (winner int, winnerDeck []int) {
	fmt.Println("=== Game X ===")
	round := 1
	for len(player1) > 0 && len(player2) > 0 {
		fmt.Printf("-- Round %d (Game X) --\n", round)
		fmt.Println("Player 1's deck:", player1)
		fmt.Println("Player 2's deck:", player2)
		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]
		fmt.Println("Player 1 plays:", card1)
		fmt.Println("Player 2 plays:", card2)
		if len(player1) < card1 || len(player2) < card2 {
			fmt.Println("Not enough cards")
			if card1 > card2 {
				winner = 1
			} else {
				winner = 2
			}
		} else {
			winner, _ = recursiveCombat(player1, player2)
		}
		fmt.Printf("Player %d wins the round!\n\n", winner)
		if winner == 1 {
			player1 = append(player1, card1)
			player1 = append(player1, card2)
		} else {
			player2 = append(player2, card2)
			player2 = append(player2, card1)
		}
		round++
	}
	if len(player1) > 0 {
		return 1, player1
	} else {
		return 2, player2
	}
}

func main() {
	// Demo data
	player1 := []int{9, 2, 6, 3, 1}
	player2 := []int{5, 8, 4, 7, 10}

	// player1 := []int{44, 31, 29, 48, 40, 50, 33, 14, 10, 30, 5, 15, 41, 45, 12, 4, 3, 17, 36, 1, 23, 34, 38, 16, 18}
	// player2 := []int{24, 20, 11, 32, 43, 9, 6, 27, 35, 2, 46, 21, 7, 49, 26, 39, 8, 19, 42, 22, 47, 28, 25, 13, 37}

	winnerDeck := combat(player1, player2)
	fmt.Println("Winner score of simple combat:", score(winnerDeck))
	fmt.Println()

	winner, winnerDeck := recursiveCombat(player1, player2)
	fmt.Println("Winner of recursive combat:", winner)
	fmt.Println("Winner's deck:", winnerDeck)
	fmt.Println("Winner's score:", score(winnerDeck))

}
