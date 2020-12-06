package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countUnique(strAnswers string) int {
	chars := make(map[rune]bool, 30)
	for _, ans := range strings.Split(strAnswers, "\n") {
		for _, char := range ans {
			chars[char] = true
		}
	}
	count := 0
	for range chars {
		count++
	}
	return count
}

func countIntersecting(strAnswers string) int {
	answers := strings.Split(strAnswers, "\n")
	intersectingChars := make(map[rune]bool, 30)
	for _, a := range answers[0] { // seed with the first word
		intersectingChars[a] = true
	}
	for _, ans := range answers {
		curChars := make(map[rune]bool, 30)
		for _, char := range ans {
			curChars[char] = true
		}
		for k := range intersectingChars {
			if !curChars[k] {
				delete(intersectingChars, k)
			}
		}
	}
	count := 0
	for range intersectingChars {
		count++
	}
	return count
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	text := string(content)
	groupAnswers := strings.Split(text, "\n\n")

	sumUniques := 0
	for _, groupAnswer := range groupAnswers {
		sumUniques += countUnique(groupAnswer)
	}
	fmt.Println(sumUniques)

	sumIntersecting := 0
	for _, groupAnswer := range groupAnswers {
		sumIntersecting += countIntersecting(groupAnswer)
	}
	fmt.Println(sumIntersecting)
}
