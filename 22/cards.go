package main

import "fmt"

func score(winner []int) int {
	s := 0
	for i := len(winner); i >= 1; i-- {
		s += winner[len(winner)-i] * i
	}
	return s

}

func combat(player1, player2 []int) (winner []int) {
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
		winner = player1
	} else {
		winner = player2
	}
	return winner
}

func main() {
	// Demo data
	// player1 := []int{9, 2, 6, 3, 1}
	// player2 := []int{5, 8, 4, 7, 10}

	player1 := []int{44, 31, 29, 48, 40, 50, 33, 14, 10, 30, 5, 15, 41, 45, 12, 4, 3, 17, 36, 1, 23, 34, 38, 16, 18}
	player2 := []int{24, 20, 11, 32, 43, 9, 6, 27, 35, 2, 46, 21, 7, 49, 26, 39, 8, 19, 42, 22, 47, 28, 25, 13, 37}

	winner := combat(player1, player2)
	fmt.Println("Winner score of simple combat:", score(winner))

}
