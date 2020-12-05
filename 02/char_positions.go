package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// each line in input is posA-posB char: password
// count number of lines which have `char` in exactly one of positions A or B
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	validLines := 0
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := re.FindStringSubmatch(line)[1:]
		posA, _ := strconv.Atoi(parsed[0])
		posB, _ := strconv.Atoi(parsed[1])
		char := []rune(parsed[2])[0]
		pwd := []rune(parsed[3])
		occurrences := 0
		if pwd[posA-1] == char {
			occurrences++
		}
		if pwd[posB-1] == char {
			occurrences++
		}
		if occurrences == 1 {
			validLines++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(validLines)

}
