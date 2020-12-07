package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var colorContents = map[string][][]string{}
var colorAncestors = map[string][]string{}

func countBagsInside(color string) int {
	count := 1
	for _, amountOfColor := range colorContents[color] {
		n, _ := strconv.Atoi(amountOfColor[0])
		color := amountOfColor[1]
		count += n * countBagsInside(color)
	}
	return count
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	re := regexp.MustCompile(`(\d+) ([^\.,]+) bags?`)
	for _, rule := range strings.Split(text, "\n") {
		s := strings.Split(rule, " bags contain ")
		color, rest := s[0], s[1]
		if rest == "no other bags." {
			continue
		}
		bagAmounts := re.FindAllStringSubmatch(rest, -1) // returns triplets of [matched_str, group_1, group_2]
		for _, bagAmount := range bagAmounts {
			colorContents[color] = append(colorContents[color], []string{bagAmount[1], bagAmount[2]})
		}
		for _, match := range bagAmounts {
			name := string(match[2])
			colorAncestors[name] = append(colorAncestors[name], color)
		}
	}

	// in how many different bags we can put a shiny gold bag
	stack := []string{"shiny gold"}
	ancestors := make(map[string]bool)
	for len(stack) > 0 {
		color := stack[0]
		stack = stack[1:]
		for _, ancestor := range colorAncestors[color] {
			ancestors[ancestor] = true
			stack = append(stack, ancestor)
		}
	}
	fmt.Println(len(ancestors))

	// how many other bags fit in a shiny gold bag
	fmt.Println(countBagsInside("shiny gold") - 1)
}
