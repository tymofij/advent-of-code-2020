package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// each line in input is minA-maxA char: password
// count number of lines which have number of `char`s between minA and maxA
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
		min, _ := strconv.Atoi(parsed[0])
		max, _ := strconv.Atoi(parsed[1])
		char := parsed[2]
		pwd := parsed[3]
		if strings.Count(pwd, char) >= min && strings.Count(pwd, char) <= max {
			validLines++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(validLines)

}
